package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"strconv"
)

// OrderOperation 订单管理
func OrderOperation(c *gin.Context) {
	//who, _ := c.Get("who")
	//whoMap := who.(model.Admin)
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
		//用户名  runner_id(奔跑者)
		if username, isE := c.GetPostForm("username"); isE == true {
			runner := model.Runner{Username: username}
			db = db.Where("runner_id=?", runner.GetRunnerId(mysql.DB))
		}
		//填写代理名字
		if aUsername, isE := c.GetPostForm("agency_runner_name"); isE == true {
			runner := model.AgencyRunner{Username: aUsername}
			db = db.Where("agency_runner_id=?", runner.GetId(mysql.DB))
		}
		//单子类型
		if species, isE := c.GetPostForm("species"); isE == true {
			db = db.Where("species=?", species)
		}
		//merchant_num 商户号
		if species, isE := c.GetPostForm("merchant_num"); isE == true {
			db = db.Where("merchant_num=?", species)
		}
		//merchant_order_num 商家订单号
		if species, isE := c.GetPostForm("merchant_order_num"); isE == true {
			db = db.Where("merchant_order_num=?", species)
		}
		//own_order
		if species, isE := c.GetPostForm("own_order"); isE == true {
			db = db.Where("own_order=?", species)
		}
		//upi
		if species, isE := c.GetPostForm("upi"); isE == true {
			db = db.Where("upi=?", species)
		}
		//channel_id  通道名字
		if species, isE := c.GetPostForm("channel_name"); isE == true {
			atom, _ := strconv.Atoi(species)
			channel := modelPay.Channel{ID: atom}
			channel.GetChannelId(mysql.DB)
			db = db.Where("channel_id=?", species)
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

}
