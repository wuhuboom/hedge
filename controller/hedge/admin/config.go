package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

func ConfigOperation(c *gin.Context) {
	action := c.Query("action")

	//获取配置
	if action == "check" {
		config := model.Config{}
		mysql.DB.Where("id=?", 1).First(&config)
		tools.ReturnSuccess2000DataCode(c, config, "ok")
		return
	}

	//	ID              int   `gorm:"primaryKey"`
	//	IfUseGoogleCode int   //是否使用谷歌验证码  ?  1是  2不
	//	ExpireTime      int64 //  过期时间    单位分钟

	if action == "update" {
		if status, isE := c.GetPostForm("IfUseGoogleCode"); isE == true {
			IfUseGoogleCode, _ := strconv.Atoi(status)
			mysql.DB.Model(&model.Config{}).Where("id=?", 1).Update(&model.Config{IfUseGoogleCode: IfUseGoogleCode})
			tools.ReturnSuccess2000Code(c, "OK")
			return
		}

		update := make(map[string]interface{})
		if status, isE := c.GetPostForm("ExpireTime"); isE == true {
			update["ExpireTime"], _ = strconv.ParseFloat(status, 64)

		}
		mysql.DB.Model(&model.Config{}).Where("id=?", 1).Update(update)
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}

}
