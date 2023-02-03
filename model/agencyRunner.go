package model

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"time"
)

type AgencyRunner struct {
	ID                          int    `gorm:"primaryKey"`
	Username                    string `gorm:"unique_index"` //用户名 唯一
	Password                    string
	PayPassword                 string             `gorm:"default:123123"` //提现密码
	Status                      int                `gorm:"default:1"`      //  状态 1正常 2封
	InvitationCode              string             //邀请码唯一性
	CashPledge                  float64            `gorm:"type:decimal(10,2);default:0"` //押金
	CollectionChannel           string             //代收通道   @分割
	PayChannel                  string             //代付通道   @分割
	CollectionPoint             float64            `gorm:"type:decimal(10,2);default:0"` //代收盈利点
	CollectionLimit             float64            `gorm:"type:decimal(10,2);default:0"` //代收额度
	UseCollectionLimit          float64            `gorm:"type:decimal(10,2);default:0"` //已经下发代收额度
	PayPoint                    float64            `gorm:"type:decimal(10,2);default:0"` //代付盈利点
	PayLimit                    float64            `gorm:"type:decimal(10,2);default:0"` //代付额度
	Created                     int64              //账号创建时间
	LastLoginTime               int64              //最后一次登录时间
	LastLoginIp                 string             //最后一次登录的ip
	LastLoginRegion             string             //最后一次登录日期
	Commission                  float64            `gorm:"type:decimal(10,2);default:0"` //佣金
	JuniorPoint                 float64            `gorm:"type:decimal(10,2);default:2"` //下级税点
	Token                       string             `gorm:"unique_index"`                 //token
	FreezeMoney                 float64            `gorm:"type:decimal(10,2);default:0"` //提现冻结金额
	WithdrawCommission          float64            `gorm:"type:decimal(10,2);default:0"` //提现手续费 (0.1%)
	FreezeCashPledge            float64            `gorm:"type:decimal(10,2);default:0"` //冻结的押金
	GoogleCode                  string             //谷歌验证码
	AccumulativeCollectionLimit float64            `gorm:"type:decimal(10,2);default:0"` //累计已用代收额度
	GoogleSwitch                int                `gorm:"default:1"`                    //谷歌开关     1开  2  关
	JuniorWithdrawCommission    float64            `gorm:"default:0"`                    //下级提现手续费
	CustomerServiceAddress      string             //客服地址
	NumberOfBindBack            int                `gorm:"default:3"`
	CollectionChannelArray      []modelPay.Channel `gorm:"-"`
	PayChannelArray             []modelPay.Channel `gorm:"-"`
	RunnerId                    int                `gorm:"-"` //奔跑者id
}

func CheckIsExistModelAgencyRunner(db *gorm.DB) {
	if db.HasTable(&AgencyRunner{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&AgencyRunner{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&AgencyRunner{})
	}
}

// GetId 获取id
func (ar *AgencyRunner) GetId(db *gorm.DB) int {
	db.Where("username=?", ar.Username).First(&ar)
	return ar.ID
}

// Add 添加代理
func (ar *AgencyRunner) Add(db *gorm.DB, redis *redis.Client) (bool, error) {
	ar.Created = time.Now().Unix()
	//生成邀请码
	code, err := tools.GetInvitationCode(6, redis)
	if err != nil {
		return false, err
	}
	ar.InvitationCode = code
	ar.Token = tools.RandStringRunes(31) //31位
	err = db.Save(ar).Error
	if err != nil {
		return false, err
	}
	//邀请码 添加到redis中
	redis.HSet("InvitationCode", ar.InvitationCode, "")
	return true, nil
}

// ChangeCashPledge 添加或者  减少押金
func (ar *AgencyRunner) ChangeCashPledge(db *gorm.DB, changeMoney float64, remark string) error {
	ups := make(map[string]interface{})
	ups["CashPledge"] = ar.CashPledge + changeMoney
	db = db.Begin()
	//修改押金
	err := db.Model(&AgencyRunner{}).Where("id=?", ar.ID).Update(ups).Error
	if err != nil {
		return err
	}
	//生成账变记录
	change := AgencyAccountChange{
		ChangeAmount: changeMoney,
		NowAmount:    ar.CashPledge + changeMoney,
		FontAmount:   ar.CashPledge, Kinds: 1, AgencyRunnerId: ar.ID, Remark: remark}
	err = change.Add(db)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// ChangeCollectionLimit 代收额度操作
func (ar *AgencyRunner) ChangeCollectionLimit(db *gorm.DB, changeMoney float64, remark string) error {
	ups := make(map[string]interface{})
	ups["CollectionLimit"] = ar.CollectionLimit + changeMoney
	db = db.Begin()
	//修改押金
	err := db.Model(&AgencyRunner{}).Where("id=?", ar.ID).Update(ups).Error
	if err != nil {
		return err
	}
	//生成账变记录
	change := AgencyAccountChange{
		ChangeAmount: changeMoney,
		NowAmount:    ar.CollectionLimit + changeMoney,
		FontAmount:   ar.CollectionLimit, Kinds: 2, AgencyRunnerId: ar.ID, Remark: remark}
	err = change.Add(db)
	if err != nil {
		db.Rollback()
		return err
	}
	db.Commit()
	return nil
}

// ChangeCommission 修改佣金和余额
func (ar *AgencyRunner) ChangeCommission(db *gorm.DB) error {
	ag := AgencyRunner{}
	err := db.Where("id=?", ar.ID).First(&ag).Error
	if err != nil {
		return eeor.OtherError("ChangeCommission:" + err.Error())
	}
	//修改佣金
	var ups map[string]interface{}
	ups["Commission"] = ag.Commission + ar.Commission
	//生成账变账单
	change := AgencyAccountChange{
		AgencyRunnerId: ar.ID,
		RunnerId:       ar.RunnerId,
		NowAmount:      ag.Commission + ar.Commission,
		ChangeAmount:   ar.Commission,
		FontAmount:     ag.Commission, Kinds: 4}
	err = change.Add(db)
	if err != nil {
		return err
	}

	//修改奔跑者 佣金

	return nil
}
