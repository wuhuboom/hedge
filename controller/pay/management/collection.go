package management

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

//操作订单

func CollectionOperation(c *gin.Context) {
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.Collection, 0)

		db := mysql.DB.Where("kinds=?", c.PostForm("kinds"))
		var total int

		if status, IsE := c.GetPostForm("merchant_num"); IsE == true {
			db = db.Where("mer_chant_num=?", status)
		}
		//条件
		if status, IsE := c.GetPostForm("status"); IsE == true {
			db = db.Where("status=?", status)
		}
		if status, IsE := c.GetPostForm("callback"); IsE == true {
			db = db.Where("callback=?", status)
		}

		if start, IsE := c.GetPostForm("start"); IsE == true {
			if end, IsE := c.GetPostForm("end"); IsE == true {
				db = db.Where("created  <=  ? and created >=  ?", end, start)
			}
		}

		db.Model(&modelPay.Collection{}).Count(&total)
		db = db.Model(&modelPay.Collection{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}

}
