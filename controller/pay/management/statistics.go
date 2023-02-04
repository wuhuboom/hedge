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
	//who, _ := c.Get("who")
	//whoMap := who.(model.Admin)
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
		var data Data
		mysql.DB.Raw("SELECT sum(actual_amount) as pay_all_amount  FROM  collections WHERE status  =2 AND  kinds=2").Scan(&data)
		mysql.DB.Raw("SELECT sum(actual_amount) as collection_all_amount  FROM  collections WHERE status  =2 AND  kinds=1").Scan(&data)
		mysql.DB.Raw("SELECT  SUM(admin_account_changes.change_amount) as  profit_for_mer FROM  admin_account_changes  LEFT JOIN  collections  ON  collections.id=admin_account_changes.collection_id  WHERE  collections.status=2 AND species =1 ").Scan(&data)
		mysql.DB.Raw("SELECT  SUM(admin_account_changes.change_amount) as  profit_for_run FROM  admin_account_changes  LEFT JOIN  collections  ON  collections.id=admin_account_changes.collection_id  WHERE  collections.status=2 AND species =3 ").Scan(&data)

		tools.ReturnSuccess2000DataCode(c, data, "OK")
		return
	}
}
