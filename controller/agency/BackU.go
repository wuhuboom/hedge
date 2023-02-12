package agency

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/controller/hedge/admin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

// GetU 代付 提现
func GetU(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.AgencyRunner)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.Record, 0)
		db := mysql.DB.Where("agency_runner_id=?", whoMap.ID)
		var total int
		//  状态
		if kinds, IsE := c.GetPostForm("status"); IsE == true {
			db = db.Where("status=?", kinds)
		}
		db.Model(model.Record{}).Count(&total)
		db = db.Model(&model.Record{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		admin.ReturnDataLIst2000(c, sl, total)
		return
	}
	//提现
	if action == "withdraw" {
		//提现金额
		amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
		//判断余额
		if whoMap.Commission < amount+amount*whoMap.WithdrawCommission || amount <= 0 {
			tools.ReturnErr101Code(c, "Insufficient account balance")
			return
		}
		//更新余额
		db := mysql.DB.Begin()
		ups := make(map[string]interface{})
		ups["Commission"] = whoMap.Commission - amount - whoMap.WithdrawCommission*amount
		ups["FreezeMoney"] = whoMap.FreezeMoney + amount + whoMap.WithdrawCommission*amount
		affected := db.Model(&model.AgencyRunner{}).Where("id=? and commission=?  and freeze_money", whoMap.ID, whoMap.Commission).Update(ups).RowsAffected
		if affected == 0 {
			tools.ReturnErr101Code(c, eeor.OtherError("u is fail"))
			return
		}
		//代理佣金账变
		change := model.AgencyAccountChange{
			AgencyRunnerId: whoMap.ID,
			NowAmount:      whoMap.Commission - amount - whoMap.WithdrawCommission*amount,
			ChangeAmount:   amount + whoMap.WithdrawCommission*amount,
			FontAmount:     whoMap.Commission, Kinds: 4}

		err := change.Add(db)
		if err != nil {
			db.Rollback()
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		//生成订单
		record := model.Record{AgencyRunnerId: whoMap.ID, WithdrawalMethod: 3, Kinds: 1, Amount: amount, WithdrawalCommission: whoMap.WithdrawCommission * amount}
		err = record.Add(db)
		if err != nil {
			db.Rollback()
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		db.Commit()
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}

}
