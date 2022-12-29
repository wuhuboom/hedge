package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type Logger struct {
	ID       int    `gorm:"primaryKey"` //
	Content  string `gorm:"type:text"`  //日志内容
	Operator string //操作者
	Kinds    int    //1 审核日志
	Grade    int    // 1 管理员日志   2玩家日志  3三方日志
	Created  int64
}

func CheckIsExistModelLogger(db *gorm.DB) {
	if db.HasTable(&Logger{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Logger{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Logger{})
	}
}

func (lo *Logger) AddSystem(db *gorm.DB) {
	lo.Created = time.Now().Unix()
	lo.Grade = 1
	db.Save(lo)
}
func (lo *Logger) AddUser(db *gorm.DB) {
	lo.Created = time.Now().Unix()
	lo.Grade = 2
	db.Save(lo)
}
func (lo *Logger) AddThree(db *gorm.DB) {
	lo.Created = time.Now().Unix()
	lo.Grade = 3
	db.Save(lo)
}
