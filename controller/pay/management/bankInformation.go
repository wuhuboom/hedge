package management

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

func BankInformation(c *gin.Context) {
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.BankInformation, 0)
		db := mysql.DB
		var total int
		//条件

		//状态
		if status, isExist := c.GetPostForm("status"); isExist == true {
			db = db.Where("status=?", status)
		}

		db.Model(modelPay.BankInformation{}).Count(&total)
		db = db.Model(&modelPay.BankInformation{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}
	if action == "add" {
		BankName := c.PostForm("bank_name")
		bankC := c.PostForm("bank_coding")
		remark := c.PostForm("remark")
		if BankName == "" || bankC == "" {
			tools.ReturnErr101Code(c, "The parameter cannot be empty")
			return
		}
		bi := modelPay.BankInformation{BankName: BankName, BankCoding: bankC, Remark: remark, Country: c.PostForm("country")}
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
		err := mysql.DB.Where("id=?", id).First(&modelPay.BankInformation{}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}

		//状态单独修改
		if status, isExist := c.GetPostForm("status"); isExist == true {
			st, _ := strconv.Atoi(status)
			err := mysql.DB.Model(&modelPay.BankInformation{}).Where("id=?", id).Update(&modelPay.BankInformation{Status: st}).Error
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
		}

		err = mysql.DB.Model(&modelPay.BankInformation{}).Where("id=?", id).Update(&modelPay.BankInformation{BankName: c.PostForm("bank_name"), BankCoding: c.PostForm("bank_coding"), Remark: c.PostForm("remark"), Country: c.PostForm("country")}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}

		tools.ReturnSuccess2000Code(c, "OK")
		return
	}

}
