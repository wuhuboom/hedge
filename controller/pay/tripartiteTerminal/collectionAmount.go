package tripartiteTerminal

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// CollectionAmount 代收订单拉起
func CollectionAmount(c *gin.Context) {
	var cpd CheckPayData
	if err := c.BindJSON(&cpd); err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}

	fmt.Println(cpd.Amount)
	mer := model.Merchant{}
	if err := mysql.DB.Where("merchant_num=?", cpd.MerChantNum).First(&mer).Error; err != nil {
		tools.ReturnErr101Code(c, "Illegal request")
		return
	}

	ga := model.Gateway{}
	err2 := mysql.DB.Where("id=?", mer.GatewayId).First(&ga).Error
	if err2 != nil {
		tools.ReturnErr101Code(c, "Illegal request")
		return
	}
	mer.Gateway = strings.Replace(ga.Gateway, "api", "", 1)
	if err := SignatureCollectionAmount(cpd, mer.ApiKey); err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}
	//签名通过  校验 金额大小
	amountFlot, _ := strconv.ParseFloat(cpd.Amount, 64)
	if amountFlot < mer.MinMoney || amountFlot > mer.MaxMoney {
		tools.ReturnErr101Code(c, "Amount out of range")
		return
	}
	//判断通道是否存在
	//判断通道状态
	ch := modelPay.Channel{}
	if err := mysql.DB.Where("channel_name=?", cpd.ChannelId).First(&ch).Error; err != nil || ch.Status != 1 || ch.Kinds != 1 {
		tools.ReturnErr101Code(c, "Channel under maintenance")
		return
	}
	//检查 通道id是否开通  或者存在
	ChannelId := strconv.Itoa(ch.ID)
	payChannelArray := strings.Split(mer.PayChannel, "@")
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
		//tools.ReturnErr101Code(c, "Order already exists")
		//return
	}

	config := model.Config{}
	err := mysql.DB.Where("id=?", 1).First(&config).Error
	if err != nil {
		config.ExpireTime = 60 * 60
	}
	//添加数据
	collection.Amount = amountFlot
	collection.ActualAmount = amountFlot
	collection.NoticeUrl = cpd.NoticeUrl
	collection.ChannelId = ch.ID
	collection.Currency = cpd.Currency
	collection.Callback = 1
	parseInt, _ := strconv.ParseInt(cpd.Amount, 10, 64)
	rand.Seed(time.Now().Unix() + parseInt + rand.Int63n(4000))
	//strconv.ParseInt(cpd.Amount, 10, 64)
	//collection.OwnOrder = "Mer" + time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(1000))

	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	collection.OwnOrder = "Mer" + timestamp + strconv.Itoa(rand.Intn(1000))
	collection.Date = time.Now().Format("2006-01-02")
	collection.ReleaseTime = time.Now().Unix() + config.ReleaseTime*60
	i := time.Now().Unix() + config.ExpireTime*60
	collection.ExpireTime = i
	is := strconv.FormatInt(i, 10)

	//判断模式
	//对冲优先级最高
	var UpiString string
	if mer.HedgeSwitch == 1 {

	}
	//跑分
	if mer.RunSwitch == 1 && UpiString == "" {
		runner := model.Runner{}
		//collection.Upi = upiBank.Upi
		UpiString, err = runner.SnagTheOrder(mysql.DB, collection)
		if err == nil {

			tools.ReturnSuccess2000DataCode(c, fmt.Sprintf(mer.Gateway+"/#/?upi=%s&amount=%s&order_num=%s&expiration=%s", UpiString, cpd.Amount, collection.OwnOrder, is), "ok")
		}
		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(UpiString)

	}
	//正常三方
	if UpiString == "" {
		ch2 := modelPay.Channel{ID: ch.ID}
		upiBank, err := ch2.GetUpi(mysql.DB)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		collection.Upi = upiBank.Upi
		collection.BankId = upiBank.ID
		err = collection.Add(mysql.DB)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		//每日统计
		mysql.DB.Model(&modelPay.ChannelBank{}).Where("bank_id=?", upiBank.ID).UpdateColumn("frequency", gorm.Expr("frequency + ?", 1))
		tools.ReturnSuccess2000DataCode(c, fmt.Sprintf(mer.Gateway+"/#/?upi=%s&amount=%s&order_num=%s&expiration=%s", upiBank.Upi, cpd.Amount, collection.OwnOrder, is), "ok")
		sta := modelPay.Statistics{TodayAllAmount: amountFlot, MerchantNum: mer.MerchantNum, TodayAllCollection: 1}
		sta.Add(mysql.DB, 3)
		return
	}
	sta := modelPay.Statistics{TodayAllAmount: amountFlot, MerchantNum: mer.MerchantNum, TodayAllCollection: 1}
	sta.Add(mysql.DB, 3)

}
