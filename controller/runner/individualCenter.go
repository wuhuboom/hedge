package runner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

//TODO    个人中心

// GetAnnouncement 获取公告
func GetAnnouncement(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	ann := make([]model.Announcement, 0)
	action := c.Query("action")
	if action == "one" {
		id, _ := strconv.Atoi(c.PostForm("id"))
		Announcement := model.Announcement{}
		mysql.DB.Where("id=?", id).First(&Announcement)
		tools.ReturnSuccess2000DataCode(c, Announcement, "ok")
		return
	}
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

// SetUpi 玩家设置 Upi
func SetUpi(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	action := c.Query("action")
	if action == "check" {
		ru := make([]model.RunnerUpi, 0)
		kinds := c.PostForm("kinds")
		mysql.DB.Where("runner_id=? and kind=?", whoMap.ID, kinds).Find(&ru)
		tools.ReturnSuccess2000DataCode(c, ru, "OK")
		return
	}

	if action == "add" {
		kinds, _ := strconv.Atoi(c.PostForm("kinds"))
		if kinds < 0 || kinds > 3 {
			tools.ReturnErr101Code(c, "Parameter violation")
			return
		}
		address := c.PostForm("address")
		if address == "" || len(address) > 100 {
			tools.ReturnErr101Code(c, "Parameter violation")
			return
		}
		upi := model.RunnerUpi{AgencyRunnerId: whoMap.AgencyRunnerId, RunnerId: whoMap.ID, Kind: kinds, Address: address}
		fmt.Println(upi)
		err := upi.Add(mysql.DB)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}

	if action == "update" {
		id := c.PostForm("id")
		upi := model.RunnerUpi{}
		err := mysql.DB.Where("id =? and runner_id =?", id, whoMap.ID).First(&upi).Error
		if err != nil {
			tools.ReturnErr101Code(c, "Illegal modification")
			return
		}
		address := c.PostForm("address")
		if address == "" || len(address) > 100 {
			tools.ReturnErr101Code(c, "Parameter violation")
			return
		}

		err = mysql.DB.Model(&model.RunnerUpi{}).Where("id=?", id).Update(&model.RunnerUpi{Address: address}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err)
			return
		}
		tools.ReturnSuccess2000Code(c, "Modified successfully")
		return

	}
}
