package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"time"
)

type RunnerUpi struct {
	ID             int    `gorm:"primaryKey"`
	Kind           int    //1 跑分收款地址 2 提现收款地址 3 提现收款 U 地址
	Address        string //
	Status         int    `gorm:"default:1"` //1正常 2地址
	AgencyRunnerId int
	RunnerId       int
	Created        int64
	Updated        int64
}

func CheckIsExistModelRunnerUpi(db *gorm.DB) {
	if db.HasTable(&RunnerUpi{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&RunnerUpi{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&RunnerUpi{})
	}
}

func (ru *RunnerUpi) Add(db *gorm.DB) error {
	//判断这个玩家这个类型是否帮过卡
	err2 := db.Where("runner_id=? and  kind=?", ru.RunnerId, ru.Kind).First(&RunnerUpi{}).Error

	if err2 == nil {
		return eeor.OtherError("This type can be bound to only one card")
	}
	err := db.Where("address=? and  kind=?", ru.Address, ru.Kind).First(&RunnerUpi{}).Error
	if err == nil {
		return eeor.OtherError("The card is already bound")
	}
	ru.Created = time.Now().Unix()
	ru.Updated = time.Now().Unix()
	ru.Status = 1
	return db.Save(ru).Error

}
