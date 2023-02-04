package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Record struct {
	ID             int    `gorm:"primaryKey"`
	OrderNum       string //订单号
	RunnerId       int
	AgencyRunnerId int
	MerchantNum    string  //商户号
	Amount         float64 `gorm:"type:decimal(10,2)"`
	Kinds          int     //1提现
	Created        int64
	Updated        int64
	Date           string
}

func CheckIsExistModelRecord(db *gorm.DB) {
	if db.HasTable(&Record{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Record{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Record{})
	}
}

// GetId 通过订单号 获取 id
func (rc *Record) GetId(db *gorm.DB) int {
	db.Where("order_num=?", rc.OrderNum).First(rc)
	return rc.ID
}

// GetOrderNum 通过 id  获取订单号
func (rc *Record) GetOrderNum(db *gorm.DB) string {
	db.Where("id=?", rc.ID).First(rc)
	return rc.OrderNum
}
