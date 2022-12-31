package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Config struct {
	ID              int    `gorm:"primaryKey"`
	IfUseGoogleCode int    //是否使用谷歌验证码  ?  1是  2不
	ExpireTime      int64  //  过期时间    单位分钟
	IndiaUrl        string //印度  api网关
}

func CheckIsExistModelConfig(db *gorm.DB) {
	if db.HasTable(&Config{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Config{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Config{})
		db.Save(&Config{ID: 1, IfUseGoogleCode: 1})
	}
}
