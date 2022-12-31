package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wangyi/GinTemplate/common"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"time"
)

// Merchant 商户号
type Merchant struct {
	ID                   int     `gorm:"primaryKey"`
	MerchantNum          string  `gorm:"unique_index"`       //商户号
	MinMoney             float64 `gorm:"type:decimal(10,2)"` //最小的充值金额
	MaxMoney             float64 `gorm:"type:decimal(10,2)"` //最大的充值金额
	WhiteIps             string  //ip白名单
	Status               int     //状态 1正常  2禁用
	ApiKey               string  // 加密key
	Kinds                int     `gorm:"default:1"` // 类型 1 对冲账号  2印度支付
	PayChannel           string  // 代收渠道
	PaidChannel          string  //代付渠道
	LoginPassword        string  //登录密码
	GoogleCode           string  //谷歌 code
	Token                string  //登录
	Gateway              string  //网关
	AllAmount            float64 `gorm:"type:decimal(10,2)"` //总金额
	AvailableAmount      float64 `gorm:"type:decimal(10,2)"` //可以用的金额
	FreezeAmount         float64 `gorm:"type:decimal(10,2)"` //冻结的金额
	CollectionCommission float64 `gorm:"type:decimal(10,2)"` //代收手续费
	PayCommission        float64 `gorm:"type:decimal(10,2)"` //代付手续费
	Created              int64
	Updated              int64
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
	err := db.Where("merchant_num=?", m.MerchantNum).First(&Merchant{}).Error
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

// ChannelIsExistOrOpenForCollection 判断支付通道是否存在  如果存在是否开启?
func (m *Merchant) ChannelIsExistOrOpenForCollection(db *gorm.DB) {

}

// AmountChange   kind  1代收 2代付
func (m *Merchant) AmountChange(db *gorm.DB, amount float64, channel int, collectionId int) (error, Merchant) {
	//查询账户余额
	common.MerchantChangeMoneyLock.RLock()
	mer := Merchant{}
	err := db.Where("merchant_num=?", m.MerchantNum).First(&mer).Error
	if err != nil {
		common.MerchantChangeMoneyLock.RUnlock()
		return err, mer
	}
	//  获取渠道
	ch := modelPay.Channel{}
	err = db.Where("channel_name=?", channel).First(&ch).Error
	if err != nil {
		common.MerchantChangeMoneyLock.RUnlock()
		return err, mer
	}
	common.MerchantChangeMoneyLock.RUnlock()
	common.MerchantChangeMoneyLock.Lock()
	defer common.MerchantChangeMoneyLock.Unlock()
	db = db.Begin()
	update := make(map[string]interface{})
	if ch.Kinds == 1 { //代收
		update["AllAmount"] = mer.AllAmount + amount
		update["CollectionCommission"] = mer.CollectionCommission + amount*ch.Rate
		update["AvailableAmount"] = mer.AvailableAmount + amount*(1-ch.Rate)
		//更新代收订单
		if err := db.Model(&modelPay.Collection{}).Where("id=?", collectionId).Update(map[string]interface{}{"ActualAmount": amount, "Commission": amount * ch.Rate, "Updated": time.Now().Unix(), "Status": 2}).Error; err != nil {
			return err, mer
		}
		//更新账变
		amountC := modelPay.AmountChange{MerchantNum: m.MerchantNum, Amount: amount, Before: mer.AllAmount,
			After: mer.AllAmount + amount, Remark: "collection"}
		if err := amountC.Add(db); err != nil {
			db.Rollback()
			return err, mer
		}

	} else if ch.Kinds == 2 { //代付
		update["AvailableAmount"] = mer.AvailableAmount - amount - (amount * ch.Rate)
		update["PayCommission"] = amount*ch.Rate + mer.PayCommission
		update["FreezeAmount"] = mer.FreezeAmount + amount + amount*ch.Rate
	}
	err = db.Model(&Merchant{}).Where("id=?", mer.ID).Update(update).Error
	if err != nil {
		db.Rollback()
		return err, mer
	}

	db.Commit()
	return nil, mer
}
