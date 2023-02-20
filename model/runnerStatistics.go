package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"time"
)

type RunnerStatistics struct {
	Id                  int     `gorm:"primaryKey"`
	AgencyRunnerId      int     //代理id
	RunnerId            int     //奔跑者id
	CollectionCount     int     //代收成功单数
	CollectionAllCount  int     //代收总单数
	CollectionAmount    float64 `gorm:"type:decimal(10,2);default:0"` //代收金额
	CollectionAllAmount float64 `gorm:"type:decimal(10,2);default:0"` //代收总金额
	Commission          float64 `gorm:"type:decimal(10,2);default:0"` //奖励金额
	PayAmount           float64 `gorm:"type:decimal(10,2);default:0"` //代付金额
	PayAllAmount        float64 `gorm:"type:decimal(10,2);default:0"` //代付总金额
	PayCount            int     //代付成功单数
	PayAllCount         int     //代付单数
	Date                string
	Created             int64
	Updated             int64 ///更新时间
}

func CheckIsExistModelRunnerStatistics(db *gorm.DB) {
	if db.HasTable(&RunnerStatistics{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&RunnerStatistics{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&RunnerStatistics{})
	}
}

// Add 新建
func (rt *RunnerStatistics) Add(db *gorm.DB) error {
	rt.Date = time.Now().Format("2006-01-02")
	rt.Created = time.Now().Unix()
	//判断今日数据是否存在
	tr2 := RunnerStatistics{}
	rt.Updated=time.Now().Unix()
	err := db.Where("runner_id=?  and  date=?", rt.RunnerId, rt.Date).First(&tr2).Error
	//不存在
	if err != nil {
		//插入
		return db.Save(rt).Error
	} else {
		//存在
		//代收
		ups := make(map[string]interface{})
		ups["CollectionAllCount"] = tr2.CollectionAllCount + rt.CollectionAllCount
		ups["CollectionAllAmount"] = tr2.CollectionAllAmount + rt.CollectionAllAmount
		ups["CollectionCount"] = tr2.CollectionCount + rt.CollectionCount
		ups["CollectionAmount"] = tr2.CollectionAmount + rt.CollectionAmount
		ups["Commission"] = tr2.Commission + rt.Commission
		ups["PayAllAmount"] = tr2.PayAllAmount + rt.PayAllAmount
		ups["PayAmount"] = tr2.PayAmount + rt.PayAmount
		ups["PayCount"] = tr2.PayCount + rt.PayCount
		ups["PayAllCount"] = tr2.PayAllCount + rt.PayAllCount
		ups["Updated"] = time.Now().Unix()
		affected := db.Model(&RunnerStatistics{}).Where("id=? and  updated=?", tr2.Id, tr2.Updated).Update(ups).RowsAffected
		if affected == 0 {
			return eeor.OtherError("RunnerStatistics add  65 u is fail")
		}
		return nil
	}
}
