package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
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
				//startInt, _ := strconv.ParseInt(start, 10, 64)
				//endInt, _ := strconv.ParseInt(end, 10, 64)
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

		db.Model(&modelPay.Collection{}).Count(&total)
		db = db.Model(&modelPay.Collection{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}
	//查看详情
	if action == "one" {
		id := c.PostForm("id")
		col := modelPay.Collection{}
		err2 := mysql.DB.Where("id=?", id).First(&col).Error
		if err2 != nil {
			tools.ReturnErr101Code(c, err2.Error())
			return
		}
		CH := modelPay.Channel{}
		err := mysql.DB.Where("id=?", col.ChannelId).First(&CH).Error
		if err == nil {
			col.ChannelId = CH.ChannelName
		}

		tools.ReturnSuccess2000DataCode(c, col, "OK")
		return

	}
	//回调
	if action == "callback" {
		id := c.PostForm("id")
		col := modelPay.Collection{}
		err := mysql.DB.Where("id=?  and  mer_chant_num  =?", id, whoMap.MerchantNum).First(&col).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		mer := model.Merchant{}
		err = mysql.DB.Where("merchant_num = ?", col.MerChantNum).First(&mer).Error
		if err != nil {
			tools.ReturnErr101Code(c, "Businesses don't exist")
			return
		}

		if col.Status != 2 {
			tools.ReturnErr101Code(c, "Sorry, you do not have permission to call back unpaid orders")
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
