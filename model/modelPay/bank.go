package modelPay

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Bank struct {
	ID                int    `gorm:"primaryKey"`
	BankInformationId int    // 银行信息id
	CardNum           string `gorm:"unique_index"` // 卡号
	Name              string //持卡人姓名
	IFSC              string
	Remark            string
	Status            int
	Created           int64
	BankCoding        string `gorm:"-"`
	BankName          string `gorm:"-"`
}

func CheckIsExistModelBank(db *gorm.DB) {
	if db.HasTable(&Bank{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Bank{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Bank{})
	}
}
func (b *Bank) Add(db *gorm.DB) error {
	b.Created = time.Now().Unix()
	b.Status = 1
	return db.Save(b).Error
}
