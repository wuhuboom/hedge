package management

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

func Bank(c *gin.Context) {

	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.Bank, 0)
		db := mysql.DB
		var total int
		//条件

		//状态
		if status, isExist := c.GetPostForm("status"); isExist == true {
			db = db.Where("status=?", status)
		}

		db.Model(modelPay.Bank{}).Count(&total)
		db = db.Model(&modelPay.Bank{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)

		for i, bank := range sl {
			in := modelPay.BankInformation{}
			mysql.DB.Where("id=?", bank.BankInformationId).First(&in)
			sl[i].BankName = in.BankName
			sl[i].BankCoding = in.BankCoding
		}

		tools.ReturnDataLIst2000(c, sl, total)
		return
	}
	if action == "add" {
		Bid, _ := strconv.Atoi(c.PostForm("bank_information_id"))
		information := modelPay.BankInformation{ID: Bid}
		if information.IsExist(mysql.DB) != nil {
			tools.ReturnErr101Code(c, "bank_information_id is  not exist")
			return
		}
		cardNumber := c.PostForm("card_number")
		remark := c.PostForm("remark")
		name := c.PostForm("name")
		IFSC := c.PostForm("IFSC")
		upi := c.PostForm("upi")
		bi := modelPay.Bank{CardNum: cardNumber, Remark: remark, Name: name, IFSC: IFSC, BankInformationId: Bid, Upi: upi}
		err := bi.Add(mysql.DB)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}
	if action == "update" {
		id := c.PostForm("id")
		err := mysql.DB.Where("id=?", id).First(&modelPay.Bank{}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		Bid, _ := strconv.Atoi(c.PostForm("bank_information_id"))

		//状态单独修改
		if status, isExist := c.GetPostForm("status"); isExist == true {
			st, _ := strconv.Atoi(status)
			err := mysql.DB.Model(&modelPay.Bank{}).Where("id=?", id).Update(&modelPay.Bank{Status: st}).Error
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			tools.ReturnSuccess2000Code(c, "OK")

			return
		}
		information := modelPay.BankInformation{ID: Bid}
		if information.IsExist(mysql.DB) != nil {
			tools.ReturnErr101Code(c, "bank_information_id is  not exist")
			return
		}
		err = mysql.DB.Model(&modelPay.Bank{}).Where("id=?", id).Update(&modelPay.Bank{
			BankInformationId: Bid,
			CardNum:           c.PostForm("card_number"),
			Name:              c.PostForm("name"),
			Remark:            c.PostForm("remark"),
			IFSC:              c.PostForm("IFSC"), Upi: c.PostForm("upi")}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}

		tools.ReturnSuccess2000Code(c, "OK")
		return
	}
}
