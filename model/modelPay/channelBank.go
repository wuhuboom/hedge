package modelPay

import (
	"fmt"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"time"
)

type ChannelBank struct {
	ID        int `gorm:"primaryKey"`
	ChannelId int //通道id
	BankId    int //   银行卡 id
	Created   int64
	Channel   Channel `gorm:"-"`
	Bank      Bank    `gorm:"-"`
}

func CheckIsExistModelChannelBank(db *gorm.DB) {
	if db.HasTable(&ChannelBank{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&ChannelBank{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&ChannelBank{})
	}
}

func (ch *ChannelBank) Add(db *gorm.DB) error {
	err := db.Where("channel_id=? and  bank_id= ?", ch.ChannelId, ch.BankId).First(&ChannelBank{}).Error
	if err == nil {
		return eeor.OtherError("Repeat addition")
	}
	ch.Created = time.Now().Unix()
	return db.Save(ch).Error

}
