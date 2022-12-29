package modelPay

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type BankInformation struct {
	ID         int    `gorm:"primaryKey"`
	BankName   string `gorm:"unique_index"`
	BankCoding string
	Country    string //国家
	Remark     string //备注
	Status     int    //1 正常  2  Bank
	Created    int64
}

func CheckIsExistModelBankInformation(db *gorm.DB) {
	if db.HasTable(&BankInformation{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&BankInformation{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&BankInformation{})
	}
}

func (Bi *BankInformation) Add(db *gorm.DB) error {
	Bi.Created = time.Now().Unix()
	Bi.Status = 1
	return db.Save(Bi).Error
}

func (Bi *BankInformation) IsExist(db *gorm.DB) error {

	return db.Where("id=?", Bi.ID).First(&BankInformation{}).Error

}
