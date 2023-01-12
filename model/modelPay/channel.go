package modelPay

import (
	"fmt"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"time"
)

type Channel struct {
	ID          int     `gorm:"primaryKey"`
	Kinds       int     //  1代收通道  2代付通道
	Status      int     //1正常  2禁用
	ChannelName int     `gorm:"unique_index"` //通道名字
	Rate        float64 //    费率
	Currency    string  //币种
	Created     int64
	Updated     int64
}

func CheckIsExistModelChannel(db *gorm.DB) {
	if db.HasTable(&Channel{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Channel{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Channel{})
	}
}

func (ch *Channel) Add(db *gorm.DB) error {
	ch.Status = 1
	ch.Created = time.Now().Unix()
	ch.Updated = time.Now().Unix()
	return db.Save(ch).Error
}

func (ch *Channel) GetUpi(db *gorm.DB) (Bank, error) {
	// 获取upi
	CB := make([]ChannelBank, 0)
	db.Where("channel_id=?", ch.ID).Order("frequency ASC").Find(&CB)
	//判断金额 是否 超标
	if len(CB) == 0 {
		return Bank{}, eeor.OtherError("I'm sorry, I didn't find a match")
	}
	var TallData struct {
		SumPull float64 `json:"sum_pull"`
	}
	for _, bank := range CB {
		//判断今日已经占用的金额
		db.Raw("SELECT sum(amount)  as  sum_pull  FROM collections  WHERE bank_id =?  and  date =  ? and  release_time  > ?  and  status = 1", bank.ID, time.Now().Format("2006-01-02"), time.Now().Unix()).Scan(&TallData)
		//比大小
		Ba := Bank{}
		err := db.Where("id=?", bank.BankId).First(&Ba).Error
		if err != nil {
			continue
		}
		//符合
		fmt.Printf("限制金额:%f", Ba.LimitMoney)
		fmt.Printf("未释放金额:%f", TallData.SumPull)
		fmt.Println(Ba.LimitMoney > TallData.SumPull)
		if Ba.LimitMoney > TallData.SumPull {
			return Ba, nil
		}

	}
	return Bank{}, eeor.OtherError("I'm sorry, I didn't find a match too")
}
