package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wangyi/GinTemplate/common"
	eeor "github.com/wangyi/GinTemplate/error"
	"math/rand"
	"strconv"
	"time"
)

type PaidOrder struct {
	ID                 int `gorm:"primaryKey"`
	MerchantId         int
	ThreeOrder         string  //别人的订单
	OwnOrder           string  //自己的订单
	Amount             float64 `gorm:"type:decimal(10,2)"` //玩家要  提现的钱
	AlreadyWithdrawn   float64 `gorm:"type:decimal(10,2)"` //已经 提现的 钱
	FreezeWithdrawn    float64 `gorm:"type:decimal(10,2)"` //冻结已经提现的
	NeedWithdrawn      float64 `gorm:"type:decimal(10,2)"` //需要去提现的钱
	OtherPaidWithdrawn float64 `gorm:"type:decimal(10,2)"` //发起其他代付
	AccountName        string  //账户名
	BankCode           string  //银行编号
	BankIfsc           string  //银行ifsc
	CardNum            string  //银行卡号
	NoticeUrl          string
	Updated            int64
	Created            int64
	ExpirationTime     int64  `gorm:"-"`
	MerchantName       string `gorm:"-"`
}

func CheckIsExistModelPaidOrder(db *gorm.DB) {
	if db.HasTable(&PaidOrder{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&PaidOrder{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&PaidOrder{})
	}
}

func (po *PaidOrder) Add(db *gorm.DB) (bool, error) {
	po.Created = time.Now().Unix()
	po.Updated = time.Now().Unix()

	//进行一个判断? 订单是否重复
	err := db.Where("three_order=?", po.ThreeOrder).First(&PaidOrder{}).Error
	if err == nil {
		return false, eeor.OtherError("Order duplication")
	}
	po.OwnOrder = "Ow" + time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(1000))
	err = db.Save(po).Error
	if err == nil {
		return false, err
	}

	return true, nil

}

// 获取是有可以支付的订单

func (po *PaidOrder) GetOneCanPay(db *gorm.DB, username string) (PaidOrder, error) {
	//读取 是否有订单可以支付
	//读锁
	common.OrderLuck.RLock()
	paid := PaidOrder{}
	err := db.Where("merchant_id=?  and  need_withdrawn   >= ?", po.MerchantId, po.Amount).Order("created asc").First(&paid).Error
	if err != nil {
		//没有这条数据
		common.OrderLuck.RUnlock() //解锁
		return PaidOrder{}, eeor.OtherError("sorry,system busy!")
	}
	//开启事务   并且开启写锁
	common.OrderLuck.RUnlock() //解锁
	common.OrderLuck.Lock()
	defer common.OrderLuck.Unlock() //解锁
	db = db.Begin()
	//更新 paid
	update := make(map[string]interface{})
	update["NeedWithdrawn"] = paid.NeedWithdrawn - po.Amount
	update["FreezeWithdrawn"] = paid.FreezeWithdrawn + po.Amount
	err = db.Model(&PaidOrder{}).Where("id=?", paid.ID).Update(update).Error
	if err != nil {
		return PaidOrder{}, err
	}
	//生成新的支付订单
	order := PayOrder{MerchantId: po.MerchantId, Amount: po.Amount, ThreeOrder: po.ThreeOrder, PaidOrderId: paid.ID, Username: username}
	paid.ExpirationTime, err = order.AddOrder(db)
	if err != nil {
		db.Rollback()
		return PaidOrder{}, err
	}
	db.Commit()
	return paid, nil
}
