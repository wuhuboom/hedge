package three

import (
	"fmt"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"strings"
)

type CheckPaidData struct {
	MerchantNum string `json:"merchant_num"`
	OrderNum    string `json:"order_num"`
	Amount      string `json:"amount"`
	AccountName string `json:"account_name"` //账户名
	BankCode    string `json:"bank_code"`    //银行编号
	BankIfsc    string `json:"bank_ifsc"`    //银行ifsc
	CardNum     string `json:"card_num"`     //卡号
	NoticeUrl   string `json:"notice_url"`
	Sign        string `json:"sign"`
}

func (cp *CheckPaidData) CheckSignature(db *gorm.DB, ip string) (bool, error) {

	//判断商户是否存在?
	mer := model.Merchant{}
	err := db.Where("merchant_num=?", cp.MerchantNum).First(&mer).Error
	if err != nil {
		return false, eeor.OtherError("Merchant non-existence")
	}
	//检查ip
	ipArray := strings.Split(mer.WhiteIps, "@")
	if tools.IsArray(ipArray, ip) == false {
		return false, eeor.OtherError("Illegal request")
	}

	account, _ := strconv.ParseFloat(cp.Amount, 64)
	if account < mer.MinMoney || account > mer.MaxMoney {
		return false, eeor.OtherError("Illegal account")
	}
	//str := "account=" + cp.Account + "&merchant_num=" + cp.MerchantNum + "&order_num=" + cp.OrderNum + "&key=" + mer.ApiKey
	st := make(map[string]string)
	st["amount"] = cp.Amount
	st["merchant_num"] = cp.MerchantNum
	st["order_num"] = cp.OrderNum
	st["account_name"] = cp.AccountName
	st["bank_code"] = cp.BankCode
	st["bank_ifsc"] = cp.BankIfsc
	st["card_num"] = cp.CardNum
	st["notice_url"] = cp.NoticeUrl
	str := tools.AsciiKey(st)
	str = str + "&key=" + mer.ApiKey
	fmt.Println(str)
	if tools.MD5(str) != cp.Sign {
		return false, eeor.OtherError("wrong signature")
	}

	//创建订单 !
	order := model.PaidOrder{Amount: account,
		ThreeOrder: cp.OrderNum,
		MerchantId: mer.ID, BankIfsc: cp.BankIfsc, BankCode: cp.BankCode, CardNum: cp.CardNum, NoticeUrl: cp.NoticeUrl, AccountName: cp.AccountName, NeedWithdrawn: account}
	_, err = order.Add(db)
	if err != nil {
		return false, eeor.OtherError(err.Error())

	}

	return true, nil
}

type OrderInquiryVer struct {
	MerchantNum string `json:"merchant_num"` // 商户号
	OrderNum    string `json:"order_num"`    //订单号
	Sign        string `json:"sign"`
}

func MerSignAndIpCheck(db *gorm.DB, ip string, merNum string, st map[string]string, sign string) (bool, error) {
	// 商户号 是否存在
	mer := model.Merchant{}
	err := db.Where("merchant_num=?", merNum).First(&mer).Error
	if err != nil {
		return false, eeor.OtherError("Merchant non-existence")
	}
	//检查ip
	ipArray := strings.Split(mer.WhiteIps, "@")
	if tools.IsArray(ipArray, ip) == false {
		return false, eeor.OtherError("Illegal request")

	}
	//签名 检查
	///str := "merchant_num=" + cpd.MerchantNum + "&order_num=" + cpd.OrderNum + "&key=" + mer.ApiKey
	str := tools.AsciiKey(st)
	str = str + "&key=" + mer.ApiKey

	fmt.Println(str)
	if tools.MD5(str) != sign {
		return false, eeor.OtherError("wrong  sign")
	}
	return true, nil
}

type NoticeUseOtherPaidVer struct {
	MerchantNum string `json:"merchant_num"` // 商户号
	OrderNum    string `json:"order_num"`    //订单号
	Amount      string `json:"amount"`       //代付金额
	Sign        string `json:"sign"`
}
