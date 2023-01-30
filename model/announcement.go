package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Announcement struct {
	ID             int    `gorm:"primaryKey"`
	Content        string `gorm:"type:text"` //日志内容
	Status         int    `gorm:"default:1"`
	AgencyRunnerId int    //代理id
	Created        int64
}

func CheckIsExistModelAnnouncement(db *gorm.DB) {
	if db.HasTable(&Announcement{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Announcement{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Announcement{})
	}
}

func (an *Announcement) Add(db *gorm.DB) error {
	an.Created = time.Now().Unix()

	return db.Save(an).Error
}
