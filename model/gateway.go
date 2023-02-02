package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Gateway struct {
	ID      int    `gorm:"primaryKey"`
	Name    string `gorm:"unique_index"` // 名字
	Gateway string `gorm:"unique_index"`
}

func CheckIsExistModelGateway(db *gorm.DB) {
	if db.HasTable(&Gateway{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Gateway{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Gateway{})
	}
}

// Add 添加数据
func (ga *Gateway) Add(db *gorm.DB) error {
	return db.Save(ga).Error
}

// GetName 获取名字
func (ga *Gateway) GetName(db *gorm.DB) string {
	db.Where("id=?", ga.ID).First(ga)
	return ga.Name
}
