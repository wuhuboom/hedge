package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/wangyi/GinTemplate/common"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
)

func PayOperation(c *gin.Context) {
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.PayOrder, 0)
		db := mysql.DB
		var total int
		//条件
		db.Model(model.PayOrder{}).Count(&total)
		db = db.Model(&model.PayOrder{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		for i, order := range sl {
			MER := model.Merchant{}
			mysql.DB.Where("id=?", order.MerchantId).First(&MER)
			sl[i].MerchantName = MER.MerchantNum
		}
		ReturnDataLIst2000(c, sl, total)
		return
	}

	if action == "update" {
		id := c.PostForm("id")
		PId, _ := strconv.Atoi(id)
		Amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
		order1 := model.PayOrder{ID: PId}
		isE, order := order1.IsExist(mysql.DB)
		if isE == false {
			tools.ReturnErr101Code(c, "sorry,order is not exist")
			return
		}
		//审核订单   //状态   1 支付中  2支付成功  3超时   4 支付失败
		status := c.PostForm("status")
		St, _ := strconv.Atoi(status)
		if St < 1 || St > 4 {
			tools.ReturnErr101Code(c, "fail status")
			return
		}
		if order.Status == St {
			tools.ReturnErr101Code(c, "Do not double submit")
			return
		}
		remark := c.PostForm("remark")
		if tools.IsChinese(remark) {
			tools.ReturnErr101Code(c, "The remarks cannot be in Chinese")
			return
		}
		db := mysql.DB.Begin()
		//事务开始  	//修改代付订单
		err := db.Model(&model.PayOrder{}).Where("id=?", id).Update(&model.PayOrder{
			Status: St, Updated: time.Now().Unix(), Remark: remark, ActualAmount: Amount}).Error
		if err != nil {
			tools.ReturnErr101Code(c, "1"+err.Error())
			return
		}
		//支付失败
		if St == 4 || St == 2 {
			//读锁开启
			common.OrderLuck.RLock()
			paid := model.PaidOrder{}
			err = db.Where("id=?", order.PaidOrderId).First(&paid).Error

			fmt.Println(order.PaidOrderId)
			if err != nil {
				//没有找到这个代付订单
				common.OrderLuck.RUnlock() //读锁关闭
				db.Rollback()
				tools.ReturnErr101Code(c, "2"+err.Error())
				return
			}

			common.OrderLuck.RUnlock()      //读锁关闭
			common.OrderLuck.Lock()         //写锁开启
			defer common.OrderLuck.Unlock() //关闭写锁

			update := make(map[string]interface{})

			if St == 2 {
				update["FreezeWithdrawn"] = paid.FreezeWithdrawn - Amount
				update["AlreadyWithdrawn"] = paid.AlreadyWithdrawn + Amount
			} else {
				update["NeedWithdrawn"] = paid.NeedWithdrawn + Amount
				update["FreezeWithdrawn"] = paid.FreezeWithdrawn - Amount
			}
			err = db.Model(&model.PaidOrder{}).Where("id=?", paid.ID).Update(update).Error
			if err != nil {
				db.Rollback()
				tools.ReturnErr101Code(c, "3"+err.Error())
				return
			}

			notice := make(map[string]string)
			notice["PayOrderNum"] = order.ThreeOrder //订单号
			notice["amount"] = c.PostForm("amount")  //金额  退还金额(重新放回)
			notice["action"] = "fail"                //  冻结金额
			if St == 2 {
				//审核 成功  通知三方  上分
				notice["action"] = "success"
			}
			notice["PaidOrderNum"] = paid.ThreeOrder //订单号
			marshal, err := jsoniter.Marshal(notice)
			if err != nil {
				db.Rollback()
				tools.ReturnErr101Code(c, "4"+err.Error())
				return
			}

			go tools.HttpJsonPost(paid.NoticeUrl, marshal)
			db.Commit()
			tools.ReturnSuccess2000Code(c, "OK")
			return
		} else {

			tools.ReturnErr101Code(c, "sorry ,status  is  wrong")
			return
		}

	}

}
