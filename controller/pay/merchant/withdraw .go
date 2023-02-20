package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

//回U记录

func Withdraw(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Merchant)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.Record, 0)
		db := mysql.DB.Where("merchant_num=? ", whoMap.MerchantNum)
		var total int
		//条件
		db.Model(&model.Record{}).Count(&total)
		db = db.Model(&model.Record{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		for i, record := range sl {
			mer := model.Merchant{}
			mysql.DB.Where("merchant_num=?", record.MerchantNum).First(&mer)
			sl[i].TrcAddress = mer.TrcAddress
		}
		tools.ReturnDataLIst2000(c, sl, total)
		return
	}
	if action == "withdraw" {
		//提现金额
		amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
		//判断余额

		if whoMap.AvailableAmount < amount+amount*whoMap.WithdrawCommission || amount <= 0 {
			tools.ReturnErr101Code(c, "Insufficient account balance")
			return
		}

		commission := amount * whoMap.WithdrawCommission
		WithdrawAmount := amount
		amount = amount + amount*whoMap.WithdrawCommission
		//判断是否存在Trc 地址
		if whoMap.TrcAddress == "" {
			tools.ReturnErr101Code(c, "Please bind back to U address first")
			return
		}
		//减少金额
		db := mysql.DB.Begin()
		ups := make(map[string]interface{})
		ups["AvailableAmount"] = whoMap.AvailableAmount - amount
		ups["FreezeAmount"] = whoMap.FreezeAmount + amount

		affected := db.Model(&model.Merchant{}).
			Where("merchant_num=? and  freeze_amount =? and  available_amount=?", whoMap.MerchantNum, whoMap.FreezeAmount, whoMap.AvailableAmount).Update(ups).RowsAffected
		if affected == 0 {
			tools.ReturnErr101Code(c, eeor.OtherError("u f"))
			return
		}
		//生成record
		record := model.Record{
			MerchantNum:      whoMap.MerchantNum,
			Kinds:            1,
			WithdrawalMethod: 3,
			Amount:           WithdrawAmount, WithdrawalCommission: commission}

		err := record.Add(db)
		if err != nil {
			db.Rollback()
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		//账变
		change := modelPay.AmountChange{
			MerchantNum: whoMap.MerchantNum,
			Amount:      -amount,
			Before:      whoMap.AvailableAmount,
			After:       whoMap.AvailableAmount - amount,
			Kinds:       1, RecordId: record.ID, Remark: "提现订单:" + record.OrderNum}
		err = change.Add(db)
		if err != nil {
			db.Rollback()
			tools.ReturnErr101Code(c, err.Error())
			return
		}

		db.Commit()
		tools.ReturnSuccess2000Code(c, "Submission completed, waiting for administrator review")
		return
	}

}
