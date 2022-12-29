package modelPay

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Channel struct {
	ID          int     `gorm:"primaryKey"`
	Kinds       int     //  1代收通道  2代付通道
	Status      int     //1正常  2禁用
	ChannelName int     `gorm:"unique_index"` //通道名字
	Rate        float64 //    费率
	Currency    string  //币种
	Created     int64
	Updated     int64
}

func CheckIsExistModelChannel(db *gorm.DB) {
	if db.HasTable(&Channel{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Channel{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Channel{})
	}
}

func (ch *Channel) Add(db *gorm.DB) error {
	ch.Status = 1
	ch.Created = time.Now().Unix()
	ch.Updated = time.Now().Unix()
	return db.Save(ch).Error
}
