package three

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/common"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
)

// Paid 代付
func Paid(c *gin.Context) {
	var cpd CheckPaidData
	err := c.BindJSON(&cpd)
	if err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}
	//检查签名是否合法
	_, err = cpd.CheckSignature(mysql.DB, c.ClientIP())
	if err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}
	tools.ReturnSuccess2000Code(c, "ok")
	return
}

// OrderInquiry 查询订单情况
func OrderInquiry(c *gin.Context) {
	var cpd OrderInquiryVer
	err := c.BindJSON(&cpd)
	if err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}
	//检查签名
	st := make(map[string]string)
	st["merchant_num"] = cpd.MerchantNum
	st["order_num"] = cpd.OrderNum
	_, err = MerSignAndIpCheck(mysql.DB, c.ClientIP(), cpd.MerchantNum, st, cpd.Sign)
	if err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}

	//获取代付账单
	paid := model.PaidOrder{}
	err = mysql.DB.Where("three_order=?", cpd.OrderNum).First(&paid).Error
	if err != nil {
		tools.ReturnErr101Code(c, "order is not exist")
		return
	}
	tools.ReturnSuccess2000DataCode(c, paid, "ok")
	return
}

// NoticeUseOtherPaid 通知使用其他的代付
func NoticeUseOtherPaid(c *gin.Context) {
	var cpd NoticeUseOtherPaidVer
	err := c.BindJSON(&cpd)
	if err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}
	//检查签名
	st := make(map[string]string)
	st["merchant_num"] = cpd.MerchantNum
	st["order_num"] = cpd.OrderNum
	st["amount"] = cpd.Amount
	_, err = MerSignAndIpCheck(mysql.DB, c.ClientIP(), cpd.MerchantNum, st, cpd.Sign)
	if err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}

	//通知   要发起其他的代付
	//上读锁锁
	common.OrderLuck.RLock()
	paid := model.PaidOrder{}
	err = mysql.DB.Where("three_order=?", cpd.OrderNum).First(&paid).Error
	if err != nil {
		common.OrderLuck.RUnlock() //解锁
		tools.ReturnErr101Code(c, "order  is  not exist")
		return
	}
	common.OrderLuck.RUnlock()      //解锁
	common.OrderLuck.Lock()         //读锁
	defer common.OrderLuck.Unlock() //解锁

	//判断 剩下的钱是否可以 提取
	Amount, _ := strconv.ParseFloat(cpd.Amount, 64)
	if Amount > paid.NeedWithdrawn {
		tools.ReturnErr101Code(c, "Not enough money ,Revise the amount and try again,please")
		return
	}

	db := mysql.DB.Begin()
	update := make(map[string]interface{})
	update["NeedWithdrawn"] = paid.NeedWithdrawn - Amount
	update["OtherPaidWithdrawn"] = paid.OtherPaidWithdrawn + Amount
	err = db.Model(&model.PaidOrder{}).Where("id=?", paid.ID).Update(update).Error
	if err != nil {
		db.Rollback()
		tools.ReturnErr101Code(c, err.Error())
		return
	}
	tools.ReturnSuccess2000Code(c, "OK")
	db.Commit()
	return
}
