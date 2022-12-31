package tripartiteTerminal

import (
	"encoding/json"
	"fmt"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/tools"
)

type CheckPayData struct {
	ChannelId        string `json:"channel_id"`         //通道 id
	MerChantNum      string `json:"mer_chant_num"`      //商户号
	MerchantOrderNum string `json:"merchant_order_num"` //商家订单号
	Amount           string `json:"amount"`             //代付金额
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
