package management

import (
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
)

//操作订单

func CollectionOperation(c *gin.Context) {
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.Collection, 0)

		db := mysql.DB.Where("kinds=?", c.PostForm("kinds"))
		var total int

		if status, IsE := c.GetPostForm("merchant_num"); IsE == true {
			db = db.Where("mer_chant_num=?", status)
		}
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
		//商家订单号
		if status, IsE := c.GetPostForm("merchant_order_num"); IsE == true {
			db = db.Where("merchant_order_num=?", status)
		}

		//平台订单号
		if status, IsE := c.GetPostForm("own_order"); IsE == true {
			db = db.Where("own_order=?", status)
		}

		//支付凭证  proof_of_payment
		if status, IsE := c.GetPostForm("proof_of_payment"); IsE == true {
			db = db.Where("proof_of_payment=?", status)
		}

		db.Model(&modelPay.Collection{}).Count(&total)
		db = db.Model(&modelPay.Collection{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}

	//   操作订单状态 (失败/成功)
	if action == "confirmationOfPayment" {
		id := c.PostForm("id")
		col := modelPay.Collection{}
		err := mysql.DB.Where("id=?", id).First(&col).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}

		//  status  1等待支付  2支付成功  3失败
		sta, _ := strconv.Atoi(c.PostForm("status"))
		if sta != 2 && sta != 3 {
			tools.ReturnErr101Code(c, "Illegal request")
			return
		}
		if sta == col.Status {
			tools.ReturnErr101Code(c, "Don't repeat changes")
			return
		}

		//查询商家
		mer := model.Merchant{}
		err = mysql.DB.Where(" merchant_num= ?", col.MerChantNum).First(&mer).Error
		if err != nil {
			tools.ReturnErr101Code(c, "Businesses don't exist")
			return
		}

		//回调给 三方t

		if sta == 3 {
			// 充值失败
			if err := mysql.DB.Model(&modelPay.Collection{}).Where("id=?", col.ID).Update(&modelPay.Collection{
				Status: 3, Updated: time.Now().Unix()}).Error; err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			if tools.IsChinese(c.PostForm("remark")) {
				tools.ReturnErr101Code(c, "remark is  not  Chinese")
				return
			}
		} else {
			ActualAmount, _ := strconv.ParseFloat(c.PostForm("actual_amount"), 64)
			if ActualAmount <= 0 {
				tools.ReturnErr101Code(c, "actual_amount must  have")
				return
			}
			//充值成功
			merchant := model.Merchant{MerchantNum: col.MerChantNum}
			err, _ := merchant.AmountChange(mysql.DB, ActualAmount, col.ChannelId, col.ID, col.MerchantOrderNum)
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}

		}
		callback := map[string]string{}
		callback["merchant_order_num"] = col.MerchantOrderNum
		callback["channel_id"] = strconv.Itoa(col.ChannelId)
		callback["status"] = strconv.Itoa(sta)
		callback["remark"] = c.PostForm("remark")
		callback["amount"] = strconv.FormatFloat(col.Amount, 'f', 2, 64)
		callback["actual_amount"] = c.PostForm("actual_amount")
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

	//回调
	if action == "callback" {

		id := c.PostForm("id")
		col := modelPay.Collection{}
		err := mysql.DB.Where("id=?", id).First(&col).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}

		mer := model.Merchant{}
		err = mysql.DB.Where(" merchant_num= ?", col.MerChantNum).First(&mer).Error
		if err != nil {
			tools.ReturnErr101Code(c, "Businesses don't exist")
			return
		}

		callback := map[string]string{}
		callback["merchant_order_num"] = col.MerchantOrderNum
		callback["channel_id"] = strconv.Itoa(col.ChannelId)
		callback["status"] = strconv.Itoa(col.Status)
		callback["remark"] = c.PostForm("remark")
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
			Updated:         time.Now().Unix(),
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

}
