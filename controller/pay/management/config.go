package management

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

// GatewayOperation 网关管理
func GatewayOperation(c *gin.Context) {
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.Gateway, 0)
		db := mysql.DB
		var total int
		db.Model(&model.Gateway{}).Count(&total)
		db = db.Model(&model.Gateway{}).Offset((page - 1) * limit).Limit(limit)
		db.Find(&sl)
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}

	if action == "add" {
		name := c.PostForm("name")
		gateway := c.PostForm("gateway")
		if name == "" || gateway == "" {
			tools.ReturnErr101Code(c, "The parameter cannot be null")
			return
		}
		m := model.Gateway{Gateway: gateway, Name: name}
		err := m.Add(mysql.DB)
		if err != nil {
			tools.ReturnErr101Code(c, err)
			return
		}
		tools.ReturnSuccess2000Code(c, "Add successfully")
		return
	}

	if action == "update" {
		id := c.PostForm("id")
		name := c.PostForm("name")
		gateway := c.PostForm("gateway")
		if name == "" || gateway == "" {
			tools.ReturnErr101Code(c, "The parameter cannot be null")
			return
		}
		err := mysql.DB.Where("id=?", id).First(&model.Gateway{}).Error
		if err != nil {
			tools.ReturnErr101Code(c, "The new gateway does not exist")
			return

		}
		err = mysql.DB.Model(&model.Gateway{}).Where("id=?", id).Update(&model.Gateway{Name: name, Gateway: gateway}).Error
		if err != nil {
			tools.ReturnErr101Code(c, err)
			return
		}
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}

}
