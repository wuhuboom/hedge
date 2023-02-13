package agency

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
)

// HomePage 首页获取数据
func HomePage(c *gin.Context) {

	who, _ := c.Get("who")
	whoMap := who.(model.AgencyRunner)
	action := c.Query("action")
	if action == "list" {
		type Data struct {
			CollectionCount     int     `json:"collection_count"`      //代收成功单数
			CollectionAllCount  int     `json:"collection_all_count"`  //代收总单数
			CollectionAmount    float64 `json:"collection_amount"`     //代收金额
			CollectionAllAmount float64 `json:"collection_all_amount"` //代收总金额
			Commission          float64 `json:"commission"`            //奖励金额

			PayAmount    float64 `json:"pay_amount"`     //代付金额
			PayAllAmount float64 `json:"pay_all_amount"` //代付总金额
			PayCount     int     `json:"pay_count"`      //代付成功单数
			PayAllCount  int     `json:"pay_all_count"`  //代付单数
			Date         string  `json:"date"`
		}

		var data []Data
		mysql.DB.Raw("SELECT SUM(collection_all_amount) as collection_all_amount,SUM(commission) as commission,SUM(pay_amount) as pay_amount,SUM(pay_all_amount) as pay_all_amount,SUM(pay_count) as pay_count,SUM(pay_all_count) as pay_all_count,SUM(collection_all_count) as collection_all_count , SUM(collection_count) AS  collection_count ,SUM(collection_amount) as collection_amount, date  as date FROM runner_statistics    WHERE  agency_runner_id=?   GROUP BY  agency_runner_id,date  ", whoMap.ID).Scan(&data)

		tools.ReturnSuccess2000DataCode(c, data, "OK")
		return
	}

	if action == "all" {
		type Data struct {
			Balance     float64 `json:"balance"`      //可用余额
			Commission  float64 `json:"commission"`   //累计佣金
			FreezeMoney float64 `json:"freeze_money"` //冻结金额
		}

		var data Data

		data.Balance = whoMap.Balance
		data.FreezeMoney = whoMap.FreezeMoney
		data.Commission = whoMap.Commission
		tools.ReturnSuccess2000DataCode(c, data, "OK")
		return
	}

}
