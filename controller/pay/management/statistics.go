package management

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
)

// GetStatistics 管理统计
func GetStatistics(c *gin.Context) {
	action := c.Query("action")

	if action == "list" {
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.Statistics, 0)
		db := mysql.DB.Where("date=?", time.Now().Format("2006-01-02"))
		var total int
		//条件
		if start, IsE := c.GetPostForm("start"); IsE == true {
			if end, IsE := c.GetPostForm("end"); IsE == true {
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
		//总的代收
		type Data struct {
			CollectionAllAmount float64 `json:"collection_all_amount"` //总代收
			PayAllAmount        float64 `json:"pay_all_amount"`        //总代付
			ProfitForMer        float64 `json:"profit_for_mer"`        //盈利金额(三方)
			ProfitForRun        float64 `json:"profit_for_run"`        //盈利金额(跑分)
		}
		mysql.DB.Raw("SELECT SUM(amount) FROM collections  WHERE  status=2")
		return
	}
}
