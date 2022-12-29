package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"strconv"
)

// PaidOperation 代付 订单
func PaidOperation(c *gin.Context) {
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.PaidOrder, 0)
		db := mysql.DB
		var total int
		//条件
		db.Model(model.PaidOrder{}).Count(&total)
		db = db.Model(&model.PaidOrder{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)

		for i, order := range sl {
			MER := model.Merchant{}
			mysql.DB.Where("id=?", order.MerchantId).First(&MER)
			sl[i].MerchantName = MER.MerchantNum
		}

		ReturnDataLIst2000(c, sl, total)
		return
	}
}
