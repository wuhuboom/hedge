package agency

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/controller/hedge/admin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

// CollectionOperation  获取  代收订单
func CollectionOperation(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.AgencyRunner)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.Collection, 0)
		db := mysql.DB.Where("agency_runner_id=?  and  kinds=1", whoMap.ID)
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
		db.Model(&modelPay.Collection{}).Count(&total)
		db = db.Model(&modelPay.Collection{}).Offset((page - 1) * limit).Limit(limit).Order("updated desc")
		db.Find(&sl)
		for i, collection := range sl {
			channel := modelPay.Channel{ID: collection.ChannelId}
			sl[i].ChannelId = channel.GetChannelName(mysql.DB)
			runner := model.Runner{ID: collection.RunnerId}
			sl[i].RunnerName = runner.GetRunnerUsername(mysql.DB)
		}
		admin.ReturnDataLIst2000(c, sl, total)
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
