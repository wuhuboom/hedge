package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"math/rand"
	"strconv"
	"time"
)

// PayOrder 支付
type PayOrder struct {
	ID               int     `gorm:"primaryKey"`
	MerchantId       int     //商户号
	ThreeOrder       string  //别人的订单
	OwnOrder         string  //自己的订单
	Amount           float64 `gorm:"type:decimal(10,2)"` //充值的金额
	ActualAmount     float64 `gorm:"type:decimal(10,2)"` //实际到账多少钱
	PaidOrderId      int     //代付订单 id
	Status           int     //状态   1 支付中  2支付成功  3超时   4 支付失败
	ExpireTime       int64   //过期时间
	CertificateNum   string  //凭证订单号
	CertificateImage string  //凭证截图
	Username         string  //玩家名字
	Updated          int64
	Created          int64
	Remark           string
	MerchantName     string `gorm:"-"`
}

func CheckIsExistModelPayOrder(db *gorm.DB) {
	if db.HasTable(&PayOrder{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&PayOrder{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&PayOrder{})
	}
}

//创建订单

func (po *PayOrder) AddOrder(db *gorm.DB) (int64, error) {

	//判断订单是否重复
	err2 := db.Where("three_order=?", po.ThreeOrder).First(&PayOrder{}).Error
	if err2 == nil {
		return -1, eeor.OtherError("Order duplication")
	}

	//判断这个用户是否有已经存在存在没有支付的订单
	//err2 = db.Where("username=? and  status  =?", po.Username, 1).First(&PayOrder{}).Error
	//if err2 == nil {
	//	//  存在没支付的订单
	//	return -2, eeor.OtherError("You have an outstanding order")
	//
	//}

	po.Updated = time.Now().Unix()
	po.Created = time.Now().Unix()
	po.OwnOrder = "Py" + time.Now().Format("20060102150405") + strconv.Itoa(rand.Intn(1000))
	//获取过期时间

	config := Config{}
	err := db.Where("id=?", 1).First(&config).Error
	po.ExpireTime = po.Created + 30*60
	if err == nil {
		po.ExpireTime = po.Created + config.ExpireTime*60
	}
	po.Status = 1
	err = db.Save(po).Error
	if err != nil {
		return -1, err
	}

	return po.ExpireTime, nil
}

//  判断订单是否存在

func (po *PayOrder) IsExist(db *gorm.DB) (bool, *PayOrder) {
	err := db.Where("id=?", po.ID).First(po).Error
	if err != nil {
		return false, po
	}

	return true, po
}
