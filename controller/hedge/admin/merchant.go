package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
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
		for i, merchant := range sl {
			gateway := model.Gateway{ID: merchant.GatewayId}
			sl[i].Gateway = gateway.GetName(mysql.DB)
		}

		ReturnDataLIst2000(c, sl, total)
		return
	}

	if action == "add" {
		mer := model.Merchant{GoogleSwitch: 2}
		mer.MerchantNum = c.PostForm("merchant")
		mer.WhiteIps = c.PostForm("ips")
		mer.MinMoney, _ = strconv.ParseFloat(c.PostForm("min_money"), 64)
		mer.MaxMoney, _ = strconv.ParseFloat(c.PostForm("max_money"), 64)
		mer.PayChannel = c.PostForm("pay_channel")
		mer.PaidChannel = c.PostForm("paid_channel")
		mer.Kinds, _ = strconv.Atoi(c.PostForm("kinds"))
		mer.GatewayId, _ = strconv.Atoi(c.PostForm("gateway_id"))
		mer.LoginPassword = c.PostForm("login_password")
		mer.LoginPassword = tools.MD5(mer.LoginPassword)
		mer.MaxPay, _ = strconv.ParseFloat(c.PostForm("max_pay"), 64)
		mer.MinPay, _ = strconv.ParseFloat(c.PostForm("min_pay"), 64)

		mer.Token = tools.RandStringRunes(36)
		//通道币种
		if mer.MerchantNum == "" {
			tools.ReturnErr101Code(c, "The parameter cannot be empty")
			return
		}
		mer.ApiKey = tools.RandStringRunes(48) + mer.MerchantNum
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

		if status, isExist := c.GetPostForm("hedge_switch"); isExist == true {
			atop, _ := strconv.Atoi(status)
			mysql.DB.Model(&model.Merchant{}).Where("id=?", id).Update(&model.Merchant{HedgeSwitch: atop})
			tools.ReturnSuccess2000Code(c, "ok")
			return
		}

		if status, isExist := c.GetPostForm("run_switch"); isExist == true {
			atop, _ := strconv.Atoi(status)
			mysql.DB.Model(&model.Merchant{}).Where("id=?", id).Update(&model.Merchant{RunSwitch: atop})
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

		//max   min  pay
		if minMoney, isExist := c.GetPostForm("min_pay"); isExist == true {
			if maxMoney, isExist := c.GetPostForm("max_pay"); isExist == true {
				Min, _ := strconv.ParseFloat(minMoney, 64)
				Max, _ := strconv.ParseFloat(maxMoney, 64)
				if Min >= Max {
					tools.ReturnErr101Code(c, "sorry ,min_pay must  < max_pay")
					return
				}
				updated["MinPay"] = Min
				updated["MaxPay"] = Max
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
		if status, isExist := c.GetPostForm("gateway_id"); isExist == true {
			updated["GatewayId"], _ = strconv.Atoi(status)
		}
		//登录密码
		if status, isExist := c.GetPostForm("login_password"); isExist == true {
			if mer.LoginPassword != status {
				updated["LoginPassword"] = tools.MD5(status)
			}

		}
		//谷歌开关
		if status, isExist := c.GetPostForm("google_switch"); isExist == true {
			updated["GoogleSwitch"], _ = strconv.Atoi(status)
		}
		//谷歌code
		if status, isExist := c.GetPostForm("google_code"); isExist == true {
			updated["GoogleCode"] = status
		}
		mysql.DB.Model(&model.Merchant{}).Where("id=?", id).Update(updated)
		tools.ReturnSuccess2000Code(c, "ok")
		return
	}
	//重置谷歌
	if action == "resetGoogle" {

		id := c.PostForm("id")
		//判断这个商户号是否存在?
		mer := model.Merchant{}
		err := mysql.DB.Where("id=?", id).First(&mer).Error
		if err != nil {
			tools.ReturnErr101Code(c, "sorry ,the  merchant is not  exist")
			return
		}
		mysql.DB.Model(&model.Merchant{}).Where("id=?", id).Update(map[string]interface{}{"GoogleCode": ""})
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}
	//修改商户余额
	if action == "available_amount" {

		id := c.PostForm("id")
		//判断这个商户号是否存在?
		mer := model.Merchant{}
		err := mysql.DB.Where("id=?", id).First(&mer).Error
		if err != nil {
			tools.ReturnErr101Code(c, "sorry ,the  merchant is not  exist")
			return
		}
		result, _ := redis.Rdb.Get("available_amount_" + id).Result()
		if result != "" {
			tools.ReturnErr101Code(c, "Too fast. Try again later")
			return
		}

		amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
		remark := c.PostForm("remark")
		if remark == "" || amount == 0 {
			tools.ReturnErr101Code(c, "The parameter cannot be empty")
			return
		}

		if mer.AvailableAmount+amount < 0 {

			tools.ReturnErr101Code(c, "sorry  ,available_amount can't be less than 0")
			return
		}

		//修改可用余额
		db := mysql.DB.Begin()
		err = db.Model(&model.Merchant{}).Where("id=? and  available_amount=? ", mer.ID, mer.AvailableAmount).Update(map[string]interface{}{"AvailableAmount": mer.AvailableAmount + amount}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		change := modelPay.AmountChange{MerchantNum: mer.MerchantNum, Amount: amount, After: mer.AvailableAmount + amount, Before: mer.AvailableAmount, Remark: remark}
		err = change.Add(db)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			db.Rollback()
			return
		}
		db.Commit()
		tools.ReturnSuccess2000Code(c, "OK")

		redis.Rdb.Set("available_amount_"+id, "---", 5*time.Second)

		return
	}

	tools.ReturnErr101Code(c, "An unlawful request")
	return

}
