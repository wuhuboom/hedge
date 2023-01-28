package model

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/tools"
	"time"
)

type Runner struct {
	ID                 int    `gorm:"primaryKey"`
	Username           string `gorm:"unique_index"`
	Password           string
	PayPassword        string  `gorm:"default:123123"` //提现密码
	AgencyRunnerId     int     //代理账号
	InvitationCode     string  `gorm:"unique_index"` //邀请码
	RegisterTime       int64   //注册时间
	Status             int     `gorm:"default:1"` //状态 1正常 2封
	Superior           int     //上级 玩家
	LeverTree          string  //上级树  @ 分割
	CollectionPoint    float64 `gorm:"type:decimal(10,2)"`           //代收赢利点
	PayPoint           float64 `gorm:"type:decimal(10,2)"`           //代付盈利点
	CashPledge         float64 `gorm:"type:decimal(10,2);default:0"` //押金
	CollectionLimit    float64 `gorm:"type:decimal(10,2);default:0"` //代收额度
	PayLimit           float64 `gorm:"type:decimal(10,2);default:0"` //代付额度
	Commission         float64 `gorm:"type:decimal(10,2);default:0"` //佣金
	FreezeMoney        float64 `gorm:"type:decimal(10,2);default:0"` //提现冻结金额
	JuniorPoint        float64 `gorm:"type:decimal(10,2)"`           //下级税点
	WithdrawCommission float64 `gorm:"type:decimal(10,2);default:0"` //提现手续费
	Token              string  `gorm:"unique_index"`                 //Token   唯一标识   长度  48位
	LastLoginTime      int64   //最后一次登陆时间
	LastLoginIp        string  //最后一次登录的ip
	LastLoginRegion    string  //最后一次登录地区
	PaySwitch          int     `gorm:"default:2"`  //代付开关   1开  2关
	ActiveGrade        int     `gorm:"default:60"` //活跃分数
	CreditScore        int     `gorm:"default:60"` //信用分

}

func CheckIsExistModelRunner(db *gorm.DB) {
	if db.HasTable(&Runner{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Runner{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Runner{})
	}
}

func (r *Runner) GetRunnerId(db *gorm.DB) int {
	db.Where("username=?", r.Username).First(r)
	return r.ID
}

// Add 创建奔跑者
func (r *Runner) Add(db *gorm.DB, redis *redis.Client) error {
	err2 := db.Where("username=?", r.Username).First(r).Error
	if err2 == nil {
		return eeor.OtherError("Sorry, the user already exists")
	}
	r.InvitationCode, _ = tools.GetInvitationCode(8, redis)
	r.Token = tools.RandStringRunes(48)
	r.RegisterTime = time.Now().Unix()
	err := db.Save(r).Error
	if err == nil {
		redis.HSet("InvitationCode", r.InvitationCode, "")
	}
	return err
}

// ChangePassword 修改密码
func (r *Runner) ChangePassword(db *gorm.DB) error {
	return db.Model(&Runner{}).Where("id=?", r.ID).Update(&Runner{Password: r.Password}).Error
}

// ChangePayPassword 修改支付密码
func (r *Runner) ChangePayPassword(db *gorm.DB) error {
	return db.Model(&Runner{}).Where("id=?", r.ID).Update(&Runner{PayPassword: r.PayPassword}).Error
}

// IsExist 判断是否存在
func (r *Runner) IsExist(db *gorm.DB) (bool, *Runner) {
	err := db.Where("id=?", r.ID).First(r).Error
	if err != nil {
		return false, r
	}
	return true, r
}

func (r *Runner) CheckInvitationCode(db *gorm.DB) error {

	if len(r.InvitationCode) == 8 {
		//上级是会员

	} else if len(r.InvitationCode) == 6 {
		//上级代理

	} else {
		return eeor.OtherError("The invitation code is invalid")
	}
}
