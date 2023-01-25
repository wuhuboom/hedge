package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type AgencyAccountChange struct {
	ID                      int     `gorm:"primaryKey"`
	AgencyRunnerId          int     //代理id
	RunnerId                int     //奔跑者id
	NowAmount               float64 `gorm:"type:decimal(10,2)"` //变化后的金额
	ChangeAmount            float64 `gorm:"type:decimal(10,2)"` //变化金额
	FontAmount              float64 `gorm:"type:decimal(10,2)"` //变化之前金额
	Kinds                   int     //1押金 2代收额度  3代付额度 4佣金   5提现
	RunnerCollectionOrderId int
	RunnerPayOrderId        int
	RecordId                int
	Remark                  string
	Created                 int64
}

func CheckIsExistModelAgencyAccountChange(db *gorm.DB) {
	if db.HasTable(&AgencyAccountChange{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&AgencyAccountChange{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&AgencyAccountChange{})
	}
}

// Add 创建订单
func (aac *AgencyAccountChange) Add(db *gorm.DB) error {
	aac.Created = time.Now().Unix()
	return db.Save(aac).Error
}
