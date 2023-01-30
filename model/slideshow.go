package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Slideshow struct {
	ID             int `gorm:"primaryKey"`
	AgencyRunnerId int
	Url            string
	Status         int `gorm:"default:1"` //状态 1显示 2不显示
	Created        int64
}

func CheckIsExistModelSlideshow(db *gorm.DB) {
	if db.HasTable(&Slideshow{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Slideshow{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Slideshow{})
	}
}

func (s *Slideshow) Add(db *gorm.DB) error {
	s.Created = time.Now().Unix()
	return db.Save(s).Error
}
