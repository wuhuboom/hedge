package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type AdminAccountChange struct {
	ID           int     `gorm:"primaryKey" `
	NowAmount    float64 `gorm:"type:decimal(10,2)"` //变化后的金额
	ChangeAmount float64 `gorm:"type:decimal(10,2)"` //变化金额
	FontAmount   float64 `gorm:"type:decimal(10,2)"` //变化之前金额
	Kinds        int     //类型 1盈利金额  2累计押金
	CollectionId int     //订单金额
	RecordId     int
	Created      int64
}

func CheckIsExistModelAdminAccountChange(db *gorm.DB) {
	if db.HasTable(&AdminAccountChange{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&AdminAccountChange{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&AdminAccountChange{})
	}
}

func (aac *AdminAccountChange) Add(db *gorm.DB) error {
	aac.Created = time.Now().Unix()
	return db.Save(aac).Error
}
