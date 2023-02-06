package modelPay

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type AmountChange struct {
	ID           int `gorm:"primaryKey"`
	MerchantNum  string
	Amount       float64 `gorm:"type:decimal(10,2)"`
	Before       float64 `gorm:"type:decimal(10,2)"`
	After        float64 `gorm:"type:decimal(10,2)"`
	Kinds        int     `gorm:"default:1"` //种类  1可用余额
	Status       int     `gorm:"default:1"` //1等待审核  2审核通过  3审核失败
	CollectionId int     //代收代付订单id
	RecordId     int
	Remark       string
	Date         string
	Created      int64
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
	ac.Date = time.Now().Format("2006-01-02")
	return db.Save(ac).Error
}
