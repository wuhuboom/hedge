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
	ID                    int     `gorm:"primaryKey"`
	MerchantNum           string  `gorm:"unique_index"`       //商户号
	MinMoney              float64 `gorm:"type:decimal(10,2)"` //最小的充值金额
	MaxMoney              float64 `gorm:"type:decimal(10,2)"` //最大的充值金额
	WhiteIps              string  //ip白名单
	Status                int     //状态 1正常  2禁用
	ApiKey                string  // 加密key
	Kinds                 int     `gorm:"default:1"` // 类型 1 对冲账号  2印度支付
	PayChannel            string  // 代收渠道
	PaidChannel           string  //代付渠道
	LoginPassword         string  //登录密码
	GoogleCode            string  //谷歌 code
	GoogleSwitch          int     `gorm:"default:2"` //谷歌开关  //1开  2关
	Token                 string  //登录
	GatewayId             int     // 网关id
	Gateway               string  //网关
	AllAmount             float64 `gorm:"type:decimal(10,2)"` //总金额
	AvailableAmount       float64 `gorm:"type:decimal(10,2)"` //可以用的金额
	FreezeAmount          float64 `gorm:"type:decimal(10,2)"` //冻结的金额
	CollectionCommission  float64 `gorm:"type:decimal(10,2)"` //代收手续费
	PayCommission         float64 `gorm:"type:decimal(10,2)"` //代付手续费
	MinPay                float64 `gorm:"type:decimal(10,2)"` //最小代付
	MaxPay                float64 `gorm:"type:decimal(10,2)"` //最大代付
	HedgeSwitch           int     `gorm:"default:2"`          //对冲开关1开 2关
	RunSwitch             int     `gorm:"default:2"`          //跑分开关1开 2关
	Created               int64
	Updated               int64
	TrcAddress            string              //trc 回U的地址
	WithdrawCommission    float64             `gorm:"type:decimal(10,2);default:0.06"` //提现手续费
	PayC                  []modelPay.Channel  `gorm:"-"`
	ChangeAvailableAmount float64             `gorm:"-"`
	ChangeFreezeAmount    float64             `gorm:"-"`
	Col                   modelPay.Collection `gorm:"-"`
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

// ChangePassword 修改密码
func (m *Merchant) ChangePassword(db *gorm.DB) error {
	m.LoginPassword = tools.MD5(m.LoginPassword)
	return db.Model(&Merchant{}).Where("id=?", m.ID).Update(m).Error
}

// ChangeTrcAddress 修改回U地址
func (m *Merchant) ChangeTrcAddress(db *gorm.DB) error {
	return db.Model(&Merchant{}).Where("id=?", m.ID).Update(m).Error
}

// AmountChange   kind  1代收 2代付(代收,代付)
func (m *Merchant) AmountChange(db *gorm.DB, amount float64, channelId int, collectionId int, merOrder string, species int, col modelPay.Collection) (error, Merchant) {
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
	err = db.Where("id=?", channelId).First(&ch).Error
	if err != nil {
		common.MerchantChangeMoneyLock.RUnlock()
		return err, mer
	}
	common.MerchantChangeMoneyLock.RUnlock()
	common.MerchantChangeMoneyLock.Lock()
	defer common.MerchantChangeMoneyLock.Unlock()

	if species != 3 {
		db = db.Begin()
	}
	update := make(map[string]interface{})
	update["Updated"] = time.Now().Unix()
	if ch.Kinds == 1 { //代收成功的订单(管理员审核)
		update["AllAmount"] = mer.AllAmount + amount
		update["CollectionCommission"] = mer.CollectionCommission + amount*ch.Rate
		update["AvailableAmount"] = mer.AvailableAmount + amount*(1-ch.Rate)
		//更新代收订单
		if err := db.Model(&modelPay.Collection{}).
			Where("id=?  and  status=? and  updated =? and  commission=?", collectionId, col.Status, col.Updated, col.Commission).Update(map[string]interface{}{"ActualAmount": amount, "Commission": amount * ch.Rate, "Updated": time.Now().Unix(), "Status": 2}).Error; err != nil {
			return err, mer
		}
		//更新账变(总余额)
		amountC := modelPay.AmountChange{MerchantNum: m.MerchantNum, Amount: amount * (1 - ch.Rate), Before: mer.AvailableAmount,
			After: mer.AvailableAmount + amount*(1-ch.Rate), Remark: "关联订单: " + merOrder, Kinds: 1, CollectionId: col.ID}
		if err := amountC.Add(db); err != nil {
			db.Rollback()
			return err, mer
		}
		//每日统计
		sta := modelPay.Statistics{TodayCollectionAmount: amount, TodayCollectionCommission: amount * ch.Rate, MerchantNum: m.MerchantNum, TodayCollection: 1}
		err := sta.Add(db, 1)
		if err != nil {
			if species != 3 {
				db.Rollback()
			}
			return err, mer
		}

		// 佣金处理(代理-总代-商户)
		commission := Commission{
			Species:        species,
			RunnerId:       col.RunnerId,
			AgencyRunnerId: col.AgencyRunnerId,
			ActualAmount:   col.ActualAmount,
			CollectionId:   col.ID,
			MerRate:        ch.Rate,
			PayType:        1, Commission: amount * ch.Rate}
		err = commission.ChangeCommission(db)
		if err != nil {
			if species != 3 {
				db.Rollback()
			}
			return err, Merchant{}
		}

	} else if ch.Kinds == 2 { //代付
		update["AvailableAmount"] = mer.AvailableAmount - amount - (amount * ch.Rate) //可用金额-金额-手续费
		update["FreezeAmount"] = mer.FreezeAmount + amount + amount*ch.Rate
		//每日统计
		sta := modelPay.Statistics{
			TodayAllPay:       1,
			TodayPayAllAmount: amount,
			MerchantNum:       m.MerchantNum}
		err = sta.Add(db, 4)
		if err != nil {
			if species != 3 {
				db.Rollback()
			}
			return err, mer
		}
		//创建代付订单
		err = col.Add(db)
		if err != nil {
			if species != 3 {
				db.Rollback()
			}
			return err, Merchant{}
		}
		//形成账变
		amountC := modelPay.AmountChange{
			MerchantNum:  m.MerchantNum,
			Amount:       -(amount + (amount * ch.Rate)),
			Before:       mer.AvailableAmount,
			After:        mer.AvailableAmount - amount - (amount * ch.Rate),
			Remark:       "关联订单: " + merOrder,
			Kinds:        1,
			CollectionId: col.ID}
		if err := amountC.Add(db); err != nil {
			if species != 3 {
				db.Rollback()
			}
			return err, mer
		}

	}
	err = db.Model(&Merchant{}).Where("id=? and available_amount =? and freeze_amount =?", mer.ID, mer.AvailableAmount, mer.FreezeAmount).Update(update).Error
	if err != nil {
		if species != 3 {
			db.Rollback()
		}
		return err, mer
	}
	if species != 3 {
		db.Commit()
	}
	return nil, mer
}

// AvailableAmountChangeAndFreezeAmount 商户 可用金额变化
func (m *Merchant) AvailableAmountChangeAndFreezeAmount(db *gorm.DB) error {
	ups := make(map[string]interface{})

	ups["AvailableAmount"] = m.AvailableAmount + m.ChangeAvailableAmount
	ups["FreezeAmount"] = m.FreezeAmount + m.ChangeFreezeAmount
	ups["Updated"] = time.Now().Unix()
	//更新商户号
	err := db.Model(&Merchant{}).Where("id=? and available_amount =?  and  freeze_amount=?", m.ID, m.AvailableAmount, m.FreezeAmount).Update(ups).Error
	if err != nil {
		return eeor.OtherError("mer update is fail  :" + err.Error())
	}
	//产生账变
	change := modelPay.AmountChange{
		MerchantNum: m.MerchantNum,
		Amount:      m.ChangeAvailableAmount,
		After:       m.AvailableAmount + m.ChangeAvailableAmount,
		Before:      m.AvailableAmount, CollectionId: m.Col.ID, Remark: m.Col.Remark}
	err = change.Add(db)
	if err != nil {
		return err
	}
	return nil
}
