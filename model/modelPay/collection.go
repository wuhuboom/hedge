package modelPay

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

//代收

type Collection struct {
	ID                     int     `gorm:"primaryKey"`
	ChannelId              int     //  通道id
	MerChantNum            string  // 商户号
	MerchantOrderNum       string  //商家订单号
	Status                 int     //1等待支付  2支付成功  3失败 4超时
	Callback               int     //  回调  1没有回调 2回调成功 3回调失败
	Amount                 float64 `gorm:"type:decimal(10,2)"`
	ActualAmount           float64 `gorm:"type:decimal(10,2)"` //实际支付金额
	Commission             float64 `gorm:"type:decimal(10,2)"` //手续费用
	CallbackReturn         string  //回调返回的数据
	CallbackContent        string  //回调内容
	NoticeUrl              string  //通知地址
	Currency               string  //币种
	OwnOrder               string  `gorm:"unique_index"` //我方平台订单
	Kinds                  int     `gorm:"default:1" `   //1 代收 2代付
	BankName               string
	BankCode               string
	IFSC                   string
	Name                   string
	ProofOfPayment         string //支付凭证
	ProofOfPaymentImageUrl string //支付凭证图片
	BankId                 int    //   这笔代付使用的  银行卡
	Upi                    string //Upi
	//Bank             Bank   `gorm:"-"`
	Remark           string //备注
	BankNum          string `gorm:"-"` //银行卡号
	Date             string //日期
	ReleaseTime      int64  //释放时间
	ExpireTime       int64
	Created          int64
	Updated          int64
	Species          int `gorm:"default:1"` //1三方单子 2对冲  3跑单
	AgencyRunnerId   int
	RunnerId         int
	RunnerName       string `gorm:"-"`
	AgencyRunnerName string `gorm:"-"`
}

func CheckIsExistModelCollection(db *gorm.DB) {
	if db.HasTable(&Collection{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Collection{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Collection{})
	}
}

//商家订单号是否存在

func (coll *Collection) MerchantOrderNumIsExist(db *gorm.DB) error {
	return db.Where("mer_chant_num=? and merchant_order_num=?", coll.MerChantNum, coll.MerchantOrderNum).First(&Collection{}).Error
}

func (coll *Collection) Add(db *gorm.DB) error {
	coll.Created = time.Now().Unix()
	coll.Updated = time.Now().Unix()
	coll.Status = 1
	return db.Save(coll).Error
}
