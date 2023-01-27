package agency

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/controller/hedge/admin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"strconv"
)

func Logger(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.AgencyRunner)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.AgencyLog, 0)
		db := mysql.DB.Where("agency_runner_id=?", whoMap.ID)
		var total int
		////条件
		if kinds, IsE := c.GetPostForm("kinds"); IsE == true {
			db = db.Where("kinds=?", kinds)
		}

		if username, iSE := c.GetPostForm("username"); iSE == true {
			db = db.Where("content like ?", "%"+username+"%")
		}

		db.Model(model.AgencyLog{}).Count(&total)
		db = db.Model(&model.AgencyLog{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		admin.ReturnDataLIst2000(c, sl, total)
		return
	}

}
