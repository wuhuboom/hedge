package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"math/rand"
	"strconv"
	"time"
)

type Record struct {
	ID                   int    `gorm:"primaryKey"`
	OrderNum             string `gorm:"unique_index"` //订单号
	RunnerId             int
	AgencyRunnerId       int
	MerchantNum          string  //商户号
	Amount               float64 `gorm:"type:decimal(10,2)"`
	Kinds                int     //1提现
	WithdrawalMethod     int     `gorm:"default:2"` //2 提现收款地址 3 提现收款 U 地址
	Created              int64
	Updated              int64
	Status               int `gorm:"default:1"` //状态 1 审核中 2审核通过 3审核失败
	Remark               string
	ActualAmount         float64 `gorm:"type:decimal(10,2);default:0"` //实际提现金额  U
	ExchangeRate         float64 `gorm:"type:decimal(10,2);default:0"`
	WithdrawalCommission float64 `gorm:"type:decimal(10,2);default:0"` //提现手续费
	Date                 string
	Certificate          string //凭证截图
	TrcAddress           string `gorm:"-"`
}

func CheckIsExistModelRecord(db *gorm.DB) {
	if db.HasTable(&Record{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Record{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Record{})
	}
}

// GetId 通过订单号 获取 id
func (rc *Record) GetId(db *gorm.DB) int {
	db.Where("order_num=?", rc.OrderNum).First(rc)
	return rc.ID
}

// GetOrderNum 通过 id  获取订单号
func (rc *Record) GetOrderNum(db *gorm.DB) string {
	db.Where("id=?", rc.ID).First(rc)
	return rc.OrderNum
}

func (rc *Record) Add(db *gorm.DB) error {
	rc.Created = time.Now().Unix()
	rc.Updated = time.Now().Unix()
	rc.Date = time.Now().Format("2006-01-02")
	timestamp := strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	rc.OrderNum = "Tx" + timestamp + strconv.Itoa(rand.Intn(1000))
	return db.Save(rc).Error
}
