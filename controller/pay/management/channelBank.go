package management

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"strings"
)

func ChannelBank(c *gin.Context) {
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.ChannelBank, 0)
		db := mysql.DB
		var total int
		//条件
		channelId, _ := strconv.Atoi(c.PostForm("channel_id"))
		db = db.Where("channel_id=?", channelId)
		db.Model(modelPay.ChannelBank{}).Count(&total)
		db = db.Model(&modelPay.ChannelBank{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		for i, bank := range sl {

			mysql.DB.Where("id=?", bank.BankId).First(&sl[i].Bank)
			bi := modelPay.BankInformation{}
			mysql.DB.Where("id=?", &sl[i].Bank.BankInformationId).First(&bi)
			sl[i].Bank.BankName = bi.BankName
			sl[i].Bank.BankCoding = bi.BankCoding

		}
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}

	if action == "add" {
		channelId, _ := strconv.Atoi(c.PostForm("channel_id"))
		err2 := mysql.DB.Where("id=?", channelId).First(&modelPay.Channel{}).Error
		if err2 != nil {
			tools.ReturnErr101Code(c, err2.Error())
			return
		}
		banks := c.PostForm("bank_id")
		BArray := strings.Split(banks, "@")
		for _, s := range BArray {
			bankId, _ := strconv.Atoi(s)
			err2 = mysql.DB.Where("id=?", bankId).First(&modelPay.Bank{}).Error
			if err2 == nil {
				bank := modelPay.ChannelBank{ChannelId: channelId, BankId: bankId}
				bank.Add(mysql.DB)
			}
		}
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}
	if action == "del" {
		err := mysql.DB.Model(&modelPay.ChannelBank{}).Where("id=?", c.PostForm("id")).Delete(&modelPay.ChannelBank{}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		tools.ReturnSuccess2000Code(c, "ok")
		return
	}

}
