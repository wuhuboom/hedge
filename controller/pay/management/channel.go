package management

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
)

// Channel 通道处理
func Channel(c *gin.Context) {
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.Channel, 0)
		db := mysql.DB
		var total int
		//条件
		//状态
		if status, isExist := c.GetPostForm("status"); isExist == true {
			db = db.Where("status=?", status)
		}
		//1 代收通道  2代付通道
		if status, isExist := c.GetPostForm("kinds"); isExist == true {
			db = db.Where("kinds=?", status)
		}
		//currency   币种
		if status, isExist := c.GetPostForm("currency"); isExist == true {
			db = db.Where("currency=?", status)
		}

		db.Model(modelPay.Channel{}).Count(&total)
		db = db.Model(&modelPay.Channel{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}
	if action == "add" {
		channelName, _ := strconv.Atoi(c.PostForm("channel_name"))
		if channelName == 0 {

			return
		}

		kinds, _ := strconv.Atoi(c.PostForm("kinds"))
		//status, _ := strconv.Atoi(c.PostForm("status"))
		rate, _ := strconv.ParseFloat(c.PostForm("rate"), 64)
		currency := c.PostForm("currency")
		bi := modelPay.Channel{ChannelName: channelName, Kinds: kinds, Currency: currency, Rate: rate}
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
		err := mysql.DB.Where("id=?", id).First(&modelPay.Channel{}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		//状态单独修改
		if status, isExist := c.GetPostForm("status"); isExist == true {
			st, _ := strconv.Atoi(status)
			err := mysql.DB.Model(&modelPay.Channel{}).Where("id=?", id).Update(&modelPay.Channel{Status: st}).Error
			if err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
		}
		channelName, _ := strconv.Atoi(c.PostForm("channel_name"))
		kinds, _ := strconv.Atoi(c.PostForm("kinds"))
		rate, _ := strconv.ParseFloat(c.PostForm("rate"), 64)
		currency := c.PostForm("currency")
		err = mysql.DB.Model(&modelPay.Channel{}).Where("id=?", id).Update(&modelPay.Channel{
			ChannelName: channelName,
			Kinds:       kinds,
			Rate:        rate,
			Currency:    currency, Updated: time.Now().Unix()}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}

		tools.ReturnSuccess2000Code(c, "OK")
		return

	}

}
