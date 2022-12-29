package h5

import (
	"fmt"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"strings"
)

type CheckPayData struct {
	ThreeOrder  string `json:"three_order"` //   三方订单
	Amount      string `json:"amount"`      //充值金额
	MerchantNum string `json:"merchant_num"`
	Username    string `json:"username"`
	Sign        string `json:"sign"`
	NoticeUrl   string `json:"notice_url"`
}

func (cp *CheckPayData) CheckSign(db *gorm.DB, ip string) (model.PaidOrder, error) {
	//判断商户是否存在?
	mer := model.Merchant{}
	err := db.Where("merchant_num=?", cp.MerchantNum).First(&mer).Error
	if err != nil {
		return model.PaidOrder{}, eeor.OtherError("Merchant non-existence")
	}
	//检查ip
	ipArray := strings.Split(mer.WhiteIps, "@")
	if tools.IsArray(ipArray, ip) == false {
		return model.PaidOrder{}, eeor.OtherError("Illegal request")
	}

	account, _ := strconv.ParseFloat(cp.Amount, 64)
	if account < mer.MinMoney || account > mer.MaxMoney {
		return model.PaidOrder{}, eeor.OtherError("Illegal account")
	}

	//参数校验

	st := make(map[string]string)
	st["amount"] = cp.Amount
	st["merchant_num"] = cp.MerchantNum
	st["three_order"] = cp.ThreeOrder
	st["username"] = cp.Username
	st["notice_url"] = cp.NoticeUrl
	str := tools.AsciiKey(st)
	str = str + "&key=" + mer.ApiKey
	fmt.Println(str)
	if tools.MD5(str) != cp.Sign {
		return model.PaidOrder{}, eeor.OtherError("wrong signature")
	}

	//是否存在没有支付的订单
	PayOrder := model.PayOrder{}
	err2 := db.Where("username=? and  status  =?", cp.Username, 1).First(&PayOrder).Error
	if err2 == nil {
		//  存在没支付的订单
		paid := model.PaidOrder{}
		db.Where("id=?", PayOrder.PaidOrderId).First(&paid)
		paid.ExpirationTime = PayOrder.ExpireTime
		return paid, eeor.OtherError("You have an outstanding order")
	}

	//去获取代付的订单
	order := model.PaidOrder{Amount: account, MerchantId: mer.ID, ThreeOrder: cp.ThreeOrder}
	paid, err := order.GetOneCanPay(db, cp.Username)
	if err != nil {
		return model.PaidOrder{}, err
	}

	return paid, nil
}
