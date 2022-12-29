package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/tools"
	"time"
)

// Merchant 商户号
type Merchant struct {
	ID          int     `gorm:"primaryKey"`
	MerchantNum string  //商户号
	MinMoney    float64 `gorm:"type:decimal(10,2)"` //最小的充值金额
	MaxMoney    float64 `gorm:"type:decimal(10,2)"` //最大的充值金额
	WhiteIps    string  //ip白名单
	Status      int     //状态 1正常  2禁用
	ApiKey      string  // 加密key
	Kinds       int     `gorm:"default:1"` // 类型 1 对冲账号  2印度支付
	PayChannel  string  // 支付渠道
	PaidChannel string  //代付渠道
	Created     int64
	Updated     int64
}

func CheckIsExistModelMerchant(db *gorm.DB) {
	if db.HasTable(&Config{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Merchant{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Config{})
	}
}

// Add 添加商户号
func (m *Merchant) Add(db *gorm.DB) (Merchant, error) {
	//判断这个用户是否存在 ?
	err := db.Where("merchant_num").First(&Merchant{}).Error
	if err == nil {
		return Merchant{}, eeor.OtherError("Do not repeat")
	}
	m.Created = time.Now().Unix()
	m.Updated = time.Now().Unix()
	m.Status = 1
	m.ApiKey = tools.RandStringRunes(48) + m.MerchantNum
	err = db.Save(m).Error
	if err != nil {
		return Merchant{}, err
	}
	return *m, nil
}
