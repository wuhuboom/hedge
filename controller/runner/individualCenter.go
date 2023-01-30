package runner

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
)

//TODO    个人中心

// GetAnnouncement 获取公告
func GetAnnouncement(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	ann := make([]model.Announcement, 0)
	mysql.DB.Where("agency_runner_id=? and status= ?", whoMap.AgencyRunnerId, 1).Find(&ann)
	tools.ReturnSuccess2000DataCode(c, ann, "OK")
	return
}

// GetCustomerServiceAddress 获取客服地址
func GetCustomerServiceAddress(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	runner := model.AgencyRunner{}
	mysql.DB.Where("id=?", whoMap.AgencyRunnerId).First(&runner)
	tools.ReturnSuccess2000DataCode(c, map[string]interface{}{"CustomerServiceAddress": runner.CustomerServiceAddress}, "OK")
	return
}
