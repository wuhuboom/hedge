package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/wangyi/GinTemplate/dao/mmdb"
	"time"
)

type AgencyLog struct {
	ID             int    `gorm:"primaryKey"`
	Content        string `gorm:"type:text"` //日志内容
	AgencyRunnerId int    //属于那个代理的日志
	Kinds          int    //1登录日志  2...
	Created        int64
}

func CheckIsExistModelAgencyLog(db *gorm.DB) {
	if db.HasTable(&AgencyLog{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&AgencyLog{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&AgencyLog{})
	}
}

type AgencyLogData struct {
	Ip             string
	Content        string
	Area           string
	Username       string
	Kinds          int
	AgencyRunnerId int
}

func (al *AgencyLogData) Add(db *gorm.DB) {
	al.Area, _ = mmdb.GetCountryForIp(al.Ip)
	ALog := AgencyLog{}
	ALog.Kinds = al.Kinds
	ALog.Created = time.Now().Unix()
	ALog.Content = al.Username + "|" + al.Ip + "|" + al.Area + "|" + al.Content
	ALog.AgencyRunnerId = al.AgencyRunnerId
	db.Save(&ALog)
}
