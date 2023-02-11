package management

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"os"
	"strconv"
	"strings"
	"time"
)

// GetRecord 超管处理提现订单
func GetRecord(c *gin.Context) {

	who, _ := c.Get("who")
	whoMap := who.(model.Admin)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.Record, 0)
		db := mysql.DB.Where("kinds=1").Where("runner_id=0")
		var total int
		if status, IsE := c.GetPostForm("merchant_num"); IsE == true {
			db = db.Where("mer_chant_num=?", status)
		}
		//填写代理名字
		if aUsername, isE := c.GetPostForm("agency_runner_name"); isE == true {
			runner := model.AgencyRunner{Username: aUsername}
			db = db.Where("agency_runner_id=?", runner.GetId(mysql.DB))
		}
		//条件 状态 1 审核中 2审核通过 3审核失败
		if status, IsE := c.GetPostForm("status"); IsE == true {
			db = db.Where("status=?", status)
		}
		//时间条件
		if start, IsE := c.GetPostForm("start"); IsE == true {
			if end, IsE := c.GetPostForm("end"); IsE == true {
				db = db.Where("created  <=  ? and created >=  ?", end, start)
			}
		}
		//平台订单号(提现)
		if status, IsE := c.GetPostForm("order_num"); IsE == true {
			db = db.Where("order_num=?", status)
		}

		db.Model(&model.Record{}).Count(&total)
		db = db.Model(&model.Record{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}

	//审核
	if action == "audit" {
		id := c.PostForm("id")
		result, _ := redis.Rdb.Get("CashOrder_" + id).Result()
		if result != "" {
			tools.ReturnErr101Code(c, "Do not work on the same order for a short period of time")
			return
		}

		//判断这个订单是否存在
		re := model.Record{}
		err := mysql.DB.Where("id=?", id).First(&re).Error
		if err != nil {
			tools.ReturnErr101Code(c, "the order  is not exist")
			return
		}
		if re.Status != 1 {
			tools.ReturnErr101Code(c, "Audit failed, the status of the order is incorrect")
			return
		}
		//对提交的状态限制
		status, _ := strconv.Atoi(c.PostForm("status"))
		if status != 2 && status != 3 {
			tools.ReturnErr101Code(c, "status  is not  illegality")
			return
		}

		var path string
		var remark string
		var exchangeRate float64
		var ActualAmount float64
		if status == 2 {
			//实际转账金额  //实际转账汇率
			exchangeRate, _ = strconv.ParseFloat(c.PostForm("exchange_rate"), 64)
			if exchangeRate == 0 {
				tools.ReturnErr101Code(c, "exchange_rate is fail")
				return
			}

			ActualAmount, _ = strconv.ParseFloat(c.PostForm("actual_amount"), 64)
			if ActualAmount <= 0 {
				tools.ReturnErr101Code(c, "actual_amount is fail")
				return
			}
			//成功 需要上传  转账 凭证
			certificate, err := c.FormFile("certificate")
			if err != nil {
				tools.ReturnErr101Code(c, "Payment voucher should not be empty")
				return
			}
			nameArray := strings.Split(certificate.Filename, ".")
			//判断文件夹是否存在
			path = "./static/upload/" + time.Now().Format("20060102") + "/"
			if noExist, _ := tools.IsFileNotExist(path); noExist {
				if err := os.MkdirAll(path, 0777); err != nil {
					tools.ReturnErr101Code(c, err.Error())
					return
				}
			}
			path = path + time.Now().Format("20060102150405") + nameArray[0] + "." + nameArray[1]
			err = c.SaveUploadedFile(certificate, path)
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
		}
		if status == 3 {
			remark = c.PostForm("remark")
			if remark == "" {
				tools.ReturnErr101Code(c, "remark must have value")
				return
			}
		}
		db := mysql.DB.Begin()
		defer db.Commit()
		//商户订单
		if re.MerchantNum != "" && re.AgencyRunnerId == 0 {
			//查询商户
			mer := model.Merchant{}
			err = db.Where("merchant_num=?", re.MerchantNum).First(&mer).Error
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			ups := make(map[string]interface{})
			if status == 2 {
				//成功   1 修改订单的状态  2扣除商户的冻结金额   3.生成管理的账变  4.管理生成账变
				err = db.Model(&model.Record{}).Where("id=?", id).Update(&model.Record{
					Status: 2, Certificate: path, ExchangeRate: exchangeRate, ActualAmount: ActualAmount}).Error
				if err != nil {
					tools.ReturnErr101Code(c, err.Error())
					return
				}
				err := db.Model(&model.Admin{}).Where("id=? and  profit=?", whoMap.ID, whoMap.Profit).Update(map[string]interface{}{"Profit": whoMap.Profit + re.WithdrawalCommission}).Error
				if err != nil {
					db.Rollback()
					tools.ReturnErr101Code(c, err.Error())
					return
				}

				//生成账变
				change := model.AdminAccountChange{NowAmount: whoMap.Profit + re.WithdrawalCommission, ChangeAmount: re.WithdrawalCommission, FontAmount: whoMap.Profit, Kinds: 1, RecordId: re.ID}
				err = change.Add(db)
				if err != nil {
					db.Rollback()
					tools.ReturnErr101Code(c, err.Error())
					return
				}
			} else {
				//失败   1 修改订单的状态  2退还商户的额度
				err := db.Model(&model.Record{}).Where("id=?", id).Update(&model.Record{Status: 3}).Error
				if err != nil {
					tools.ReturnErr101Code(c, err.Error())
					return
				}
				ups["AvailableAmount"] = mer.AvailableAmount + re.Amount + re.WithdrawalCommission
				//商户账变
				change := modelPay.AmountChange{MerchantNum: mer.MerchantNum, Amount: re.Amount + re.WithdrawalCommission, Before: mer.AvailableAmount, After: mer.AvailableAmount + re.Amount + re.WithdrawalCommission, RecordId: re.ID, Remark: remark}
				err = change.Add(db)
				if err != nil {
					db.Rollback()
					tools.ReturnErr101Code(c, err.Error())
					return
				}
			}
			//计算
			ups["FreezeAmount"] = mer.FreezeAmount - (re.Amount + re.WithdrawalCommission)
			err = db.Model(&model.Merchant{}).Where("id=?", mer.ID).Update(ups).Error
			if err != nil {
				db.Rollback()
				tools.ReturnErr101Code(c, err.Error())
				return
			}

		} else if re.MerchantNum == "" && re.AgencyRunnerId != 0 {
			//代理
			//查询代理

			agency := model.AgencyRunner{}
			err := db.Where("id=?", re.AgencyRunnerId).First(&agency).Error
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}

			//ups := make(map[string]interface{})
			//if status == 2 {
			//
			//} else {
			//	//失败
			//
			//}

		} else {
			tools.ReturnErr101Code(c, "system is  fail")
			return
		}
		tools.ReturnSuccess2000Code(c, "OK")
		redis.Rdb.Set("CashOrder_"+id, "12233", 5*time.Second)
		return
	}
	//修改
	if action == "update" {
		id := c.PostForm("id")
		record := model.Record{}
		err := mysql.DB.Where("id=?", id).First(&record).Error
		if err != nil {
			tools.ReturnErr101Code(c, "this record  is not exist")
			return
		}

		//修改的内容
		ups := make(map[string]interface{})
		if actualAmount, isE := c.GetPostForm("actual_amount"); isE == true {
			ups["ActualAmount"], _ = strconv.ParseFloat(actualAmount, 64)

		}

		//汇率
		if actualAmount, isE := c.GetPostForm("exchange_rate"); isE == true {
			ups["ExchangeRate"], _ = strconv.ParseFloat(actualAmount, 64)
		}

		//支付凭证
		if certificate, isE := c.FormFile("certificate"); isE != nil {
			//成功 需要上传  转账 凭证
			nameArray := strings.Split(certificate.Filename, ".")
			//判断文件夹是否存在
			path := "./static/upload/" + time.Now().Format("20060102") + "/"
			if noExist, _ := tools.IsFileNotExist(path); noExist {
				if err := os.MkdirAll(path, 0777); err != nil {
					tools.ReturnErr101Code(c, err.Error())
					return
				}
			}
			path = path + time.Now().Format("20060102150405") + nameArray[0] + "." + nameArray[1]
			err = c.SaveUploadedFile(certificate, path)
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			ups["Certificate"] = path

		}
		err = mysql.DB.Model(&model.Record{}).Where("id=?", id).Update(ups).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		tools.ReturnSuccess2000Code(c, "OK")
		return

	}

}
