package tripartiteTerminal

import (
	"encoding/json"
	"fmt"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/tools"
)

type CheckPayData struct {
	ChannelId        string `json:"channel_id" `        //通道 id
	MerChantNum      string `json:"mer_chant_num"`      //商户号
	MerchantOrderNum string `json:"merchant_order_num"` //商家订单号
	Amount           string `json:"amount"`             //代收金额
	NoticeUrl        string `json:"notice_url"`         //通知地址
	Currency         string `json:"currency"`           // 币种
	Sign             string `json:"sign"`
}

func SignatureCollectionAmount(cpd CheckPayData, ApiKey string) error {
	b, err := json.Marshal(&cpd)
	if err != nil {
		return err
	}
	var m map[string]string
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	str := tools.AsciiKey(m)
	str = str + "&key=" + ApiKey
	fmt.Println(str)
	fmt.Println(tools.MD5(str))

	if tools.MD5(str) != cpd.Sign {
		return eeor.OtherError("wrong signature")
	}
	return nil
}

//代付参数

type CheckPayDataTwo struct {
	ChannelId        string `json:"channel_id" banding:"request"`         //通道 id
	MerChantNum      string `json:"mer_chant_num" banding:"request"`      //商户号
	MerchantOrderNum string `json:"merchant_order_num" banding:"request"` //商家订单号
	Amount           string `json:"amount"  banding:"request"`            //代收金额
	NoticeUrl        string `json:"notice_url"  banding:"request"`        //通知地址
	Currency         string `json:"currency"  banding:"request"`          // 币种
	BankName         string `json:"bank_name" banding:"request"`          // 卡名  建设银行
	BankCode         string `json:"bank_code"  banding:"request"`         //卡片编码  12345
	IFSC             string `json:"ifsc" banding:"request"`               //
	Name             string `json:"name" banding:"request"`               //持卡人名字
	Sign             string `json:"sign" banding:"request"`
}

func SignatureCollectionAmountTwo(cpd CheckPayDataTwo, ApiKey string) error {
	b, err := json.Marshal(&cpd)
	if err != nil {
		return err
	}
	var m map[string]string
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	str := tools.AsciiKey(m)
	str = str + "&key=" + ApiKey
	fmt.Println(str)
	fmt.Println(tools.MD5(str))

	if tools.MD5(str) != cpd.Sign {
		return eeor.OtherError("wrong signature")
	}
	return nil
}
