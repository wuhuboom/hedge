package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
)

// Statistics 统计
func Statistics(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Merchant)
	action := c.Query("action")
	if action == "list" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.Statistics, 0)
		db := mysql.DB.Where("merchant_num=?", whoMap.MerchantNum)
		var total int
		//条件
		if start, IsE := c.GetPostForm("start"); IsE == true {
			if end, IsE := c.GetPostForm("end"); IsE == true {
				//startInt, _ := strconv.ParseInt(start, 10, 64)
				//endInt, _ := strconv.ParseInt(end, 10, 64)
				db = db.Where("created  <=  ? and created >=  ?", end, start)
			}
		}
		db.Model(&modelPay.Statistics{}).Count(&total)
		db = db.Model(&modelPay.Statistics{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}

	if action == "all" {
		st := modelPay.Statistics{}
		mysql.DB.Where("merchant_num=? and  date= ?", whoMap.MerchantNum, time.Now().Format("2006-01-02")).First(&st)
		st.AllAmount = whoMap.AllAmount
		st.AvailableAmount = whoMap.AvailableAmount
		st.FreezeAmount = whoMap.FreezeAmount
		tools.ReturnSuccess2000DataCode(c, st, "OK")
		return
	}

}
