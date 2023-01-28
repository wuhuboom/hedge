package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type RunnerPayOrder struct {
	ID       int    `gorm:"primaryKey"`
	OrderNum string //订单号

}

func CheckIsExistModelRunnerPayOrder(db *gorm.DB) {
	if db.HasTable(&RunnerPayOrder{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&RunnerPayOrder{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&RunnerPayOrder{})
	}
}

// GetId 通过订单号 获取 id
func (rc *RunnerPayOrder) GetId(db *gorm.DB) int {
	db.Where("order_num=?", rc.OrderNum).First(rc)
	return rc.ID
}

//通过 id  获取订单号

func (rc *RunnerPayOrder) GetOrderNum(db *gorm.DB) string {
	db.Where("id=?", rc.ID).First(rc)
	return rc.OrderNum
}
