package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type RunnerCollectionOrder struct {
	ID       int    `gorm:"primaryKey"`
	OrderNum string //订单号
}

func CheckIsExistModelRunnerCollectionOrder(db *gorm.DB) {
	if db.HasTable(&RunnerCollectionOrder{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&RunnerCollectionOrder{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&RunnerCollectionOrder{})
	}
}

// GetId 通过订单号 获取 id
func (rc *RunnerCollectionOrder) GetId(db *gorm.DB) int {
	db.Where("order_num=?", rc.OrderNum).First(rc)
	return rc.ID
}

//通过 id  获取订单号

func (rc *RunnerCollectionOrder) GetOrderNum(db *gorm.DB) string {
	db.Where("id=?", rc.ID).First(rc)
	return rc.OrderNum
}
