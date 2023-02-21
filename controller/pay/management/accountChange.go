package management

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

// AgencyChangeAmount 账变(代理账变)
func AgencyChangeAmount(c *gin.Context) {
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	sl := make([]model.AgencyAccountChange, 0)

	id := c.PostForm("agency_runner_id")

	err := mysql.DB.Where("id=?", id).First(&model.AgencyRunner{}).Error
	if err != nil {
		tools.ReturnErr101Code(c, "the agency is  not exist")
		return
	}

	kinds := c.PostForm("kinds") //1押金 2代收额度  3代付额度  4佣金   5提现
	db := mysql.DB.Where("kinds=?  and  agency_runner_id =?", kinds, id)
	var total int
	//条件
	db.Model(&model.AgencyAccountChange{}).Count(&total)
	db = db.Model(&model.AgencyAccountChange{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
	db.Find(&sl)
	for i, change := range sl {
		mysql.DB.Where("id=?", change.CollectionId).First(&sl[i].Col)
	}

	tools.ReturnDataLIst2000(c, sl, total)
	return

}

// MerChangeAmount 商户账变
func MerChangeAmount(c *gin.Context) {
	//who, _ := c.Get("who")
	//whoMap := who.(model.Merchant)
	//查询bankCard
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	sl := make([]modelPay.AmountChange, 0)

	merchantNum := c.PostForm("merchant_num")

	err := mysql.DB.Where("merchant_num=?", merchantNum).First(&model.Merchant{}).Error
	if err != nil {
		tools.ReturnErr101Code(c, "merchant_num is not   exist")
		return
	}

	db := mysql.DB.Where("merchant_num=?", merchantNum)
	var total int
	//条件
	if start, IsE := c.GetPostForm("start"); IsE == true {
		if end, IsE := c.GetPostForm("end"); IsE == true {
			db = db.Where("created  <=  ? and created >=  ?", end, start)
		}
	}

	db.Model(&modelPay.AmountChange{}).Count(&total)
	db = db.Model(&modelPay.AmountChange{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
	db.Find(&sl)
	tools.ReturnDataLIst2000(c, sl, total)
	return

}

// GetAdminChangeAmount 获取管理的账变
func GetAdminChangeAmount(c *gin.Context) {
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	sl := make([]model.AdminAccountChange, 0)
	db := mysql.DB
	//Kinds        int     //类型 1盈利金额  2累计押金
	if status, isE := c.GetPostForm("kinds"); isE == true {
		db.Where("kinds=?", status)
	}

	var total int
	db.Model(&model.AdminAccountChange{}).Count(&total)
	db = db.Model(&model.AdminAccountChange{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
	db.Find(&sl)
	tools.ReturnDataLIst2000(c, sl, total)
	return

}
