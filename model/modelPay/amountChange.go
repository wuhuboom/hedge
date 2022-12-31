package modelPay

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type AmountChange struct {
	ID          int `gorm:"primaryKey"`
	MerchantNum string
	Amount      float64 `gorm:"type:decimal(10,2)"`
	Before      float64 `gorm:"type:decimal(10,2)"`
	After       float64 `gorm:"type:decimal(10,2)"`
	Remark      string
	Created     int64
}

func CheckIsExistModelAmountChange(db *gorm.DB) {
	if db.HasTable(&AmountChange{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&AmountChange{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&AmountChange{})
	}
}

func (ac *AmountChange) Add(db *gorm.DB) error {
	ac.Created = time.Now().Unix()
	return db.Save(ac).Error
}
