package modelPay

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"sync"
	"time"
)

type Statistics struct {
	ID                        int `gorm:"primaryKey"`
	MerchantNum               string
	TodayCollection           int     //今日代收笔数(成功)
	TodayAllCollection        int     //代收 总笔数
	TodayAllAmount            float64 `gorm:"type:decimal(10,2)"` //代收总金额(总)
	TodayCollectionAmount     float64 `gorm:"type:decimal(10,2)"` // 今日代收金额
	TodayCollectionCommission float64 `gorm:"type:decimal(10,2)"` // 今日代收手续费
	TodayPay                  int     //今日代付笔数
	TodayAllPay               int     //今日 总代付笔数
	TodayPayCommission        float64 `gorm:"type:decimal(10,2)"` //今日代付手续费
	TodayPayAmount            float64 `gorm:"type:decimal(10,2)"` //今日代付金额(成功)
	TodayPayAllAmount         float64 `gorm:"type:decimal(10,2)"` //今日代付金额(总)
	Created                   int64
	Updated                   int64
	Date                      string
	AllAmount                 float64 `gorm:"-"` //总金额
	AvailableAmount           float64 `gorm:"-"` //可以用的金额
	FreezeAmount              float64 `gorm:"-"` //冻结的金额
}

func CheckIsExistModelStatistics(db *gorm.DB) {
	if db.HasTable(&Statistics{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Statistics{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Statistics{})
	}
}

var StatisticsLock sync.RWMutex

// Add kind=1  代收 (成功)   kind2 代付(成功)     3代收(总)   4代付(总)
func (sta *Statistics) Add(db *gorm.DB, kind int) error {
	StatisticsLock.Lock()
	defer StatisticsLock.Unlock()
	sta.Date = time.Now().Format("2006-01-02")
	//判断今天今天数据是否存在
	sta2 := Statistics{}
	err := db.Where("date=? and merchant_num   =? ", sta.Date, sta.MerchantNum).First(&sta2).Error
	if err != nil {
		//不存在
		sta.Created = time.Now().Unix()
		sta.Updated = time.Now().Unix()
		return db.Save(sta).Error
	} else {
		//存在
		sta.Updated = time.Now().Unix()
		if kind == 1 {
			sta.TodayCollection = sta2.TodayCollection + 1
			sta.TodayCollectionCommission = sta2.TodayCollectionCommission + sta.TodayCollectionCommission
			sta.TodayCollectionAmount = sta2.TodayCollectionAmount + sta.TodayCollectionAmount
		} else if kind == 3 {
			sta.TodayAllCollection = sta2.TodayAllCollection + 1
			sta.TodayAllAmount = sta2.TodayAllAmount + sta.TodayAllAmount
		} else if kind == 2 {
			sta.TodayPay = sta2.TodayPay + 1
			sta.TodayPayCommission = sta2.TodayPayCommission + sta.TodayPayCommission
			sta.TodayPayAmount = sta2.TodayPayAmount + sta.TodayPayAmount
		} else if kind == 4 {
			sta.TodayPayAllAmount = sta2.TodayPayAllAmount + sta.TodayPayAllAmount
			sta.TodayAllPay = sta2.TodayAllPay + 1
		}
		return db.Model(&Statistics{}).Where("id=?", sta2.ID).Update(sta).Error
	}
}
