package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wangyi/GinTemplate/tools"
	"time"
)

type Admin struct {
	ID            int    `gorm:"primaryKey" `
	AdminUsername string // 管理用户名
	AdminPassword string //管理密码
	GoogleCode    string //谷歌验证码
	Token         string //  管理员token
	Created       int64
}

func CheckIsExistModelAdminAdmin(db *gorm.DB) {
	if db.HasTable(&Admin{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Admin{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Admin{})
		db.Save(&Admin{AdminUsername: "jiabanjushi", AdminPassword: "8e3db6345135296ba2dde6f45c2519cd", Created: time.Now().Unix(), Token: tools.RandStringRunes(32)})
	}
}
