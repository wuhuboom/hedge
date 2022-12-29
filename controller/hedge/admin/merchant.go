package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

// MerchantOperation 操作商户
func MerchantOperation(c *gin.Context) {
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.Merchant, 0)
		db := mysql.DB
		var total int
		//条件
		if kinds, IsE := c.GetPostForm("kinds"); IsE == true {
			db = db.Where("kinds=?", kinds)
		}

		db.Model(model.Merchant{}).Count(&total)
		db = db.Model(&model.Merchant{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		ReturnDataLIst2000(c, sl, total)
		return
	}

	if action == "add" {
		mer := model.Merchant{}

		mer.MerchantNum = c.PostForm("merchant")
		mer.WhiteIps = c.PostForm("ips")
		mer.MinMoney, _ = strconv.ParseFloat(c.PostForm("min_money"), 64)
		mer.MaxMoney, _ = strconv.ParseFloat(c.PostForm("max_money"), 64)
		mer.PayChannel = c.PostForm("pay_channel")
		mer.PaidChannel = c.PostForm("paid_channel")
		mer.Kinds, _ = strconv.Atoi(c.PostForm("kinds"))
		if mer.MerchantNum == "" {
			tools.ReturnErr101Code(c, "The parameter cannot be empty\n\n")
			return
		}

		_, err := mer.Add(mysql.DB)
		if err != nil {
			tools.ReturnErr101Code(c, err)
			return
		}
		tools.ReturnSuccess2000Code(c, "OK")
		return

	}

	if action == "update" {

		id := c.PostForm("id")
		//判断这个商户号是否存在?
		mer := model.Merchant{}
		err := mysql.DB.Where("id=?", id).First(&mer).Error
		if err != nil {
			tools.ReturnErr101Code(c, "sorry ,the  merchant is not  exist")
			return
		}
		updated := make(map[string]interface{})
		if status, isExist := c.GetPostForm("status"); isExist == true {
			atop, _ := strconv.Atoi(status)
			mysql.DB.Model(&model.Merchant{}).Where("id=?", id).Update(&model.Merchant{Status: atop})
			tools.ReturnSuccess2000Code(c, "ok")
			return
		}

		//存在 mix  max
		if minMoney, isExist := c.GetPostForm("min_money"); isExist == true {
			if maxMoney, isExist := c.GetPostForm("max_money"); isExist == true {
				Min, _ := strconv.ParseFloat(minMoney, 64)
				Max, _ := strconv.ParseFloat(maxMoney, 64)
				if Min >= Max {
					tools.ReturnErr101Code(c, "sorry ,min_money must  < max_money")
					return
				}
				updated["MinMoney"] = Min
				updated["MaxMoney"] = Max
			}
		}

		if status, isExist := c.GetPostForm("ips"); isExist == true {
			updated["WhiteIps"] = status
		}
		if status, isExist := c.GetPostForm("pay_channel"); isExist == true {

			updated["PayChannel"] = status
		}
		if status, isExist := c.GetPostForm("paid_channel"); isExist == true {
			updated["PaidChannel"] = status
		}

		mysql.DB.Model(&model.Merchant{}).Where("id=?", id).Update(updated)
		tools.ReturnSuccess2000Code(c, "ok")
		return
	}

	tools.ReturnErr101Code(c, "An unlawful request")
	return

}
