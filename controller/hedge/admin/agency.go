package admin

import (
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

// AgencyOperation 代理操作  (跑分代理)
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
			if s == "" {
				continue
			}
			err := mysql.DB.Where("id=? and  kinds =? and status=1", s, 1).First(&modelPay.Channel{}).Error
			if err != nil {
				tools.ReturnErr101Code(c, "Illegal collection channel")
				return
			}
		}
		//代付渠道进行验证
		PArray := strings.Split(c.PostForm("pay_channel"), "@")
		for _, s := range PArray {
			if s == "" {
				continue
			}
			err := mysql.DB.Where("id=? and  kinds =? and status=1", s, 2).First(&modelPay.Channel{}).Error
			if err != nil {
				tools.ReturnErr101Code(c, "Illegal collection channel")
				return
			}
		}

		add := model.AgencyRunner{
			Username:          c.PostForm("username"),
			Password:          tools.MD5(c.PostForm("password")),
			CollectionChannel: c.PostForm("collection_channel"),
			PayChannel:        c.PostForm("pay_channel"),
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
		//谷歌开关 单独修改  GoogleSwitch

		if status, isExist := c.GetPostForm("google_switch"); isExist == true {
			sta, _ := strconv.Atoi(status)
			if sta != 2 && sta != 1 {
				tools.ReturnErr101Code(c, "Please enter the correct status")
				return
			}
			err := mysql.DB.Model(&model.AgencyRunner{}).Where("id=?", id).Update(&model.AgencyRunner{GoogleSwitch: sta}).Error
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
		ups["PayChannel"] = c.PostForm("pay_channel")
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

// CollectionOperation 获取代理的订单(代收)
func CollectionOperation(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Admin)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.Collection, 0)
		db := mysql.DB.Where(" kinds=1")
		var total int
		//  状态
		if kinds, IsE := c.GetPostForm("status"); IsE == true {
			db = db.Where("status=?", kinds)
		}
		//用户名  runner_id
		if username, isE := c.GetPostForm("username"); isE == true {
			runner := model.Runner{Username: username}
			db = db.Where("runner_id=?", runner.GetRunnerId(mysql.DB))
		}
		//填写代理名字
		if aUsername, isE := c.GetPostForm("agency_runner_name"); isE == true {
			runner := model.AgencyRunner{Username: aUsername}
			db = db.Where("agency_runner_id=?", runner.GetId(mysql.DB))
		}

		db.Model(&modelPay.Collection{}).Count(&total)
		db = db.Model(&modelPay.Collection{}).Offset((page - 1) * limit).Limit(limit).Order("updated desc")
		db.Find(&sl)

		for i, collection := range sl {
			channel := modelPay.Channel{ID: collection.ChannelId}
			sl[i].ChannelId = channel.GetChannelName(mysql.DB)
			runner := model.Runner{ID: collection.RunnerId}
			sl[i].RunnerName = runner.GetRunnerUsername(mysql.DB)
		}
		ReturnDataLIst2000(c, sl, total)
		return
	}

	if action == "update" {
		id := c.PostForm("id")
		col := modelPay.Collection{}
		err := mysql.DB.Where("id=? and agency_runner_id=?", id, whoMap.ID).First(&col).Error
		if err != nil {
			tools.ReturnErr101Code(c, "Order does not exist")
			return
		}
		status, _ := strconv.Atoi(c.PostForm("status"))
		if status == 2 {

		}

	}

}
