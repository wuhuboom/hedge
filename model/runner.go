package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

type Runner struct {
	ID          int    `gorm:"primaryKey"`
	Username    string `gorm:"unique_index"`
	Password    string
	PayPassword string `gorm:"default:123123"` //提现密码

	InvitationCode     string  `gorm:"unique_index"` //邀请码
	RegisterTime       int64   //注册时间
	Status             int     //状态 1正常 2封
	Superior           int     //上级 玩家
	LeverTree          string  //上级树  @ 分割
	CollectionPoint    float64 `gorm:"type:decimal(10,2)"` //代收赢利点
	PayPoint           float64 `gorm:"type:decimal(10,2)"` //代付盈利点
	CashPledge         float64 `gorm:"type:decimal(10,2)"` //押金
	CollectionLimit    float64 `gorm:"type:decimal(10,2)"` //代收额度
	PayLimit           float64 `gorm:"type:decimal(10,2)"` //代付额度
	LastLoginTime      int64   //最后一次登陆时间
	LastLoginIp        string  //最后一次登录的ip
	LastLoginRegion    string  //最后一次登录地区
	Commission         float64 `gorm:"type:decimal(10,2)"`           //佣金
	FreezeMoney        float64 `gorm:"type:decimal(10,2)"`           //提现冻结金额
	JuniorPoint        float64 `gorm:"type:decimal(10,2)"`           //下级税点
	Token              string  `gorm:"unique_index"`                 //Token   唯一标识   长度  48位
	WithdrawCommission float64 `gorm:"type:decimal(10,2);default:0"` //提现手续费
	PaySwitch          int     //代付开关   1开  2关
	ActiveGrade        int     //活跃分数
	CreditScore        int     //信用分

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
