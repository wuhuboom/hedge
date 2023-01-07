package merchant

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

//代收

func CollectionAmount(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Merchant)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.Collection, 0)
		db := mysql.DB.Where("mer_chant_num=?", whoMap.MerchantNum).Where("kinds=?", c.PostForm("kinds"))
		var total int
		//条件
		if status, IsE := c.GetPostForm("status"); IsE == true {
			db = db.Where("status=?", status)
		}
		if status, IsE := c.GetPostForm("callback"); IsE == true {
			db = db.Where("callback=?", status)
		}

		if start, IsE := c.GetPostForm("start"); IsE == true {
			if end, IsE := c.GetPostForm("end"); IsE == true {
				db = db.Where("created  <=  ? and created >=  ?", end, start)
			}
		}

		db.Model(&modelPay.Collection{}).Count(&total)
		db = db.Model(&modelPay.Collection{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}

	if action == "update" {
		id := c.PostForm("id")
		col := modelPay.Collection{}
		err := mysql.DB.Where("id=?", id).First(&col).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}

		//  status  1等待支付  2支付成功  3失败
		if status, isE := c.GetPostForm("status"); isE == true {
			sta, _ := strconv.Atoi(status)
			if sta != 2 && sta != 3 {
				tools.ReturnErr101Code(c, "Illegal request")
				return
			}
			if sta == col.Status {
				tools.ReturnErr101Code(c, "Don't repeat changes")
				return
			}
			// 充值失败
			if sta == 3 {
				if err := mysql.DB.Model(&modelPay.Collection{}).Where("id=?", col.ID).Update(&modelPay.Collection{Status: 3}).Error; err != nil {
					tools.ReturnErr101Code(c, err.Error())
					return
				}
				tools.ReturnErr101Code(c, "OK")
				return
			}
			ActualAmount, _ := strconv.ParseFloat(c.PostForm("ActualAmount"), 64)
			//充值成功
			merchant := model.Merchant{MerchantNum: col.MerChantNum}
			err, mer := merchant.AmountChange(mysql.DB, ActualAmount, col.ChannelId, col.ID)
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}

			//回调给 三方t
			callback := map[string]string{}
			callback["merchant_order_num"] = col.MerchantOrderNum
			callback["channel_id"] = strconv.Itoa(col.ChannelId)
			callback["status"] = strconv.Itoa(col.Status)
			callback["amount"] = strconv.FormatFloat(col.Amount, 'f', 2, 64)
			callback["actual_amount"] = strconv.FormatFloat(col.ActualAmount, 'f', 2, 64)
			callback["currency"] = col.Currency
			callback["commission"] = strconv.FormatFloat(col.Commission, 'f', 2, 64)
			callback["sign"] = tools.MD5(tools.AsciiKey(callback) + "&key=" + mer.ApiKey)
			marshal, err := jsoniter.Marshal(callback)
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			returnData := tools.HttpJsonPostOne(col.NoticeUrl, marshal)
			if len(returnData) > 200 {
				returnData = "url is  fail"
			}

			up := &modelPay.Collection{
				CallbackContent: string(marshal),
				CallbackReturn:  returnData,
			}
			if returnData == "SUCCESS" {
				up.Callback = 2
			} else {
				up.Callback = 3
			}
			mysql.DB.Model(&modelPay.Collection{}).Where("id=?", col.ID).Update(up)
			tools.ReturnSuccess2000Code(c, "OK")
			return
		}

		//回调   1没有回调 2回调成功 3回调失败
		if status, isE := c.GetPostForm("callback"); isE == true {
			sta, _ := strconv.Atoi(status)
			if sta != 2 {
				tools.ReturnErr101Code(c, "illegal request")
				return
			}
			//回调
			returnData := tools.HttpJsonPostOne(col.NoticeUrl, []byte(col.CallbackContent))
			if len(returnData) > 200 {
				returnData = "url is  fail"
			}
			up := &modelPay.Collection{
				CallbackReturn: returnData,
			}
			if returnData == "SUCCESS" {
				up.Callback = 2
			} else {
				up.Callback = 3
			}
			mysql.DB.Model(&modelPay.Collection{}).Where("id=?", col.ID).Update(up)
			tools.ReturnSuccess2000Code(c, "OK")

		}

	}

}
