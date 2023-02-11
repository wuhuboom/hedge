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

// Add kind=1  代收 (成功)   kind2 代付(成功)     3代收(总)   4代付(总)    5代收冲正
func (sta *Statistics) Add(db *gorm.DB, kind int) error {
	StatisticsLock.Lock()
	defer StatisticsLock.Unlock()
	sta.Date = time.Now().Format("2006-01-02")
	up := make(map[string]interface{})
	up["Date"] = time.Now().Format("2006-01-02")
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
		up["Updated"] = time.Now().Unix()
		if kind == 1 {
			up["TodayCollection"] = sta2.TodayCollection + sta.TodayCollection
			up["TodayCollectionCommission"] = sta2.TodayCollectionCommission + sta.TodayCollectionCommission
			up["TodayCollectionAmount"] = sta2.TodayCollectionAmount + sta.TodayCollectionAmount
			db = db.Where("today_collection=? and  today_collection_commission =? and  today_collection_amount=?", sta2.TodayCollection, sta2.TodayCollectionCommission, sta2.TodayCollectionAmount)
		} else if kind == 3 {
			up["TodayAllCollection"] = sta2.TodayAllCollection + sta.TodayAllCollection
			up["TodayAllAmount"] = sta2.TodayAllAmount + sta.TodayAllAmount
			db = db.Where("today_all_collection=? and  today_all_amount =?", sta2.TodayAllCollection, sta2.TodayAllAmount)

		} else if kind == 2 {
			up["TodayPay"] = sta2.TodayPay + sta.TodayPay
			up["TodayPayCommission"] = sta2.TodayPayCommission + sta.TodayPayCommission
			up["TodayPayAmount"] = sta2.TodayPayAmount + sta.TodayPayAmount
			db = db.Where("today_pay=? and  today_pay_commission =? and today_pay_amount=? ", sta2.TodayPay, sta2.TodayPayCommission, sta2.TodayPayAmount)

		} else if kind == 4 {
			up["TodayPayAllAmount"] = sta2.TodayPayAllAmount + sta.TodayPayAllAmount
			up["TodayAllPay"] = sta2.TodayAllPay + sta.TodayAllPay
			db = db.Where("today_pay_all_amount=? and  today_all_pay =?  ", sta2.TodayPayAllAmount, sta2.TodayAllPay)

		}
		return db.Model(&Statistics{}).Where("id=?", sta2.ID).Update(up).Error
	}
}
