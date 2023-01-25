package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"math"
	"strconv"
	"strings"
	"time"
)

// AgencyOperation 代理操作
func AgencyOperation(c *gin.Context) {
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.AgencyRunner, 0)
		db := mysql.DB
		var total int
		//条件
		if status, IsE := c.GetPostForm("status"); IsE == true {
			db = db.Where("status=?", status)
		}
		db.Model(model.AgencyRunner{}).Count(&total)
		db = db.Model(&model.AgencyRunner{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		ReturnDataLIst2000(c, sl, total)
		return
	}
	if action == "add" {
		//对代收渠道进行验证
		CArray := strings.Split(c.PostForm("collection_channel"), "@")
		for _, s := range CArray {
			err := mysql.DB.Where("id=? and  kinds =?", s, 1).First(&modelPay.Channel{}).Error
			fmt.Println(err)
			if err != nil {
				tools.ReturnErr101Code(c, "Illegal collection channel")
				return
			}
		}
		add := model.AgencyRunner{
			Username:          c.PostForm("username"),
			Password:          tools.MD5(c.PostForm("password")),
			CollectionChannel: c.PostForm("collection_channel"),
		}
		add.CashPledge, _ = strconv.ParseFloat(c.PostForm("cash_pledge"), 64)
		add.CollectionPoint, _ = strconv.ParseFloat(c.PostForm("collection_point"), 64)
		add.PayLimit, _ = strconv.ParseFloat(c.PostForm("pay_limit"), 64)
		add.PayPoint, _ = strconv.ParseFloat(c.PostForm("pay_point"), 64)
		add.CollectionLimit, _ = strconv.ParseFloat(c.PostForm("collection_limit"), 64)
		add.JuniorPoint, _ = strconv.ParseFloat(c.PostForm("junior_point"), 64)
		add.WithdrawCommission, _ = strconv.ParseFloat(c.PostForm("withdraw_commission"), 64)
		_, err := add.Add(mysql.DB, redis.Rdb)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		tools.ReturnSuccess2000Code(c, "Adding agency succeeded")
		return
	}
	if action == "update" {
		id := c.PostForm("id")
		err := mysql.DB.Where("id=?", id).First(&model.AgencyRunner{}).Error
		if err != nil {
			tools.ReturnErr101Code(c, "Illegal modification")
			return
		}
		//状态单独修改
		if status, isExist := c.GetPostForm("status"); isExist == true {
			sta, _ := strconv.Atoi(status)
			if sta != 2 && sta != 1 {
				tools.ReturnErr101Code(c, "Please enter the correct status")
				return
			}
			err := mysql.DB.Model(&model.AgencyRunner{}).Where("id=?", id).Update(&model.AgencyRunner{Status: sta}).Error
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			tools.ReturnSuccess2000Code(c, "Modified successfully")
			return
		}
		//重置谷歌
		if _, isExist := c.GetPostForm("reset"); isExist == true {
			err := mysql.DB.Model(&model.AgencyRunner{}).Where("id=?", id).Update(map[string]interface{}{"GoogleCode": ""}).Error
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			tools.ReturnSuccess2000Code(c, "Modified successfully")
			return
		}
		ups := make(map[string]interface{})
		ups["Password"] = tools.MD5(c.PostForm("password"))
		ups["PayPassword"] = tools.MD5(c.PostForm("pay_password"))
		ups["CollectionChannel"] = c.PostForm("collection_channel")
		ups["CollectionPoint"], _ = strconv.ParseFloat(c.PostForm("collection_point"), 64)
		ups["PayPoint"], _ = strconv.ParseFloat(c.PostForm("pay_point"), 64)
		ups["JuniorPoint"], _ = strconv.ParseFloat(c.PostForm("junior_point"), 64)
		ups["WithdrawCommission"], _ = strconv.ParseFloat(c.PostForm("withdraw_commission"), 64)
		err = mysql.DB.Model(&model.AgencyRunner{}).Where("id=?", id).Update(ups).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		tools.ReturnSuccess2000Code(c, "Modified successfully")
		return
	}
	if action == "money" {
		id := c.PostForm("id")
		remark := c.PostForm("remark")
		if remark == "" {
			remark = "OK"
		}
		AgencyRunner := model.AgencyRunner{}
		err := mysql.DB.Where("id=?", id).First(&AgencyRunner).Error
		if err != nil {
			tools.ReturnErr101Code(c, "Illegal modification")
			return
		}
		//押金操作
		if cashPledge, isExist := c.GetPostForm("cash_pledge"); isExist == true {
			cp, _ := strconv.ParseFloat(cashPledge, 64)
			result, _ := redis.Rdb.Get("CashPledge_" + AgencyRunner.Username).Result()
			if result != "" {
				tools.ReturnErr101Code(c, "Don't do the deposit operation at the specified time")
				return
			}
			if AgencyRunner.CashPledge < math.Abs(cp) && cp < 0 {
				tools.ReturnErr101Code(c, "Sorry, the deposit is not deductible")
				return
			}
			err := AgencyRunner.ChangeCashPledge(mysql.DB, cp, remark)
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			redis.Rdb.Set("CashPledge_"+AgencyRunner.Username, cp, 5*time.Second)
			tools.ReturnSuccess2000Code(c, "OK")
			return
		}
		// collection_limit  代收额度操作
		if cashPledge, isExist := c.GetPostForm("collection_limit"); isExist == true {
			cp, _ := strconv.ParseFloat(cashPledge, 64)
			result, _ := redis.Rdb.Get("CollectionLimit_" + AgencyRunner.Username).Result()
			if result != "" {
				tools.ReturnErr101Code(c, "Don't do the deposit operation at the specified time")
				return
			}
			if AgencyRunner.CollectionLimit < math.Abs(cp) && cp < 0 {
				tools.ReturnErr101Code(c, "Sorry, the deposit is not deductible")
				return
			}
			err := AgencyRunner.ChangeCollectionLimit(mysql.DB, cp, remark)
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
			redis.Rdb.Set("CollectionLimit_"+AgencyRunner.Username, cp, 5*time.Second)
			tools.ReturnSuccess2000Code(c, "OK")
			return
		}

	}

}
