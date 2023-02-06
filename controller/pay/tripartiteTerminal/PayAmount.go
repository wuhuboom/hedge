package tripartiteTerminal

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//代付

func PayAmount(c *gin.Context) {
	var cpd CheckPayDataTwo
	if err := c.BindJSON(&cpd); err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}
	mer := model.Merchant{}
	if err := mysql.DB.Where("merchant_num=?", cpd.MerChantNum).First(&mer).Error; err != nil {
		tools.ReturnErr101Code(c, "Illegal request")
		return
	}
	if err := SignatureCollectionAmountTwo(cpd, mer.ApiKey); err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}

	//签名通过  校验 金额大小
	amountFlot, _ := strconv.ParseFloat(cpd.Amount, 64)
	if amountFlot < mer.MinPay || amountFlot > mer.MaxPay {
		tools.ReturnErr101Code(c, "Amount out of range")
		return
	}

	//判断   IFSC
	if cpd.IFSC == "" {
		tools.ReturnErr101Code(c, "IFSC  is  not  be empty")
		return
	}

	//判断通道状态
	ch := modelPay.Channel{}
	if err := mysql.DB.Where("channel_name=?", cpd.ChannelId).First(&ch).Error; err != nil || ch.Status != 1 || ch.Kinds != 2 {
		tools.ReturnErr101Code(c, "Channel under maintenance")
		return
	}

	//检查 通道id是否开通  或者存在
	ChannelId := strconv.Itoa(ch.ID)
	payChannelArray := strings.Split(mer.PaidChannel, "@")
	if tools.IsArray(payChannelArray, ChannelId) == false {
		tools.ReturnErr101Code(c, "Illegal channel")
		return
	}

	//判断币种是否正确
	if ch.Currency != cpd.Currency {
		tools.ReturnErr101Code(c, "Currency  is  fail")
		return
	}

	//订单是否重复提交
	collection := modelPay.Collection{MerchantOrderNum: cpd.MerchantOrderNum, MerChantNum: cpd.MerChantNum}
	if err := collection.MerchantOrderNumIsExist(mysql.DB); err == nil {
		tools.ReturnErr101Code(c, "Order already exists")
		return
	}
	//判断账户余额是否足够
	if mer.AvailableAmount < amountFlot+amountFlot*ch.Rate {
		tools.ReturnErr101Code(c, "Insufficient merchant balance")
		return
	}

	//添加数据
	collection.Amount = amountFlot
	collection.NoticeUrl = cpd.NoticeUrl
	collection.ChannelId = ch.ID
	collection.Currency = cpd.Currency
	collection.Callback = 1
	collection.BankName = cpd.BankName
	collection.BankCode = cpd.BankCode
	collection.IFSC = cpd.IFSC
	collection.Name = cpd.Name
	collection.Kinds = 2
	collection.OwnOrder = "Mer" + time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(1000))
	////创建订单
	merchant := model.Merchant{MerchantNum: cpd.MerChantNum}
	err, _ := merchant.AmountChange(mysql.DB, amountFlot, ch.ID, collection.ID, collection.OwnOrder, 1, collection)
	if err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}
	tools.ReturnSuccess2000Code(c, "OK")
	return

}
