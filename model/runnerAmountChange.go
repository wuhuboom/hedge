package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"time"
)

type RunnerAmountChange struct {
	ID           int `gorm:"primaryKey"`
	RunnerId     int
	NowAmount    float64 `gorm:"type:decimal(10,2);default:0"`
	ChangeAmount float64 `gorm:"type:decimal(10,2);default:0"`
	FontAmount   float64 `gorm:"type:decimal(10,2);default:0"`
	Kinds        int     //1押金   2代收额度  3代付额度  4佣金   5提现
	CollectionId int
	Remark       string
	Created      int64
	Col          modelPay.Collection `gorm:"-"`
}

func CheckIsExistModelRunnerAmountChange(db *gorm.DB) {
	if db.HasTable(&RunnerAmountChange{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&RunnerAmountChange{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&RunnerAmountChange{})
	}
}

func (rac *RunnerAmountChange) Add(db *gorm.DB) error {
	rac.Created = time.Now().Unix()
	return db.Save(rac).Error
}
