package model

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"math"
	"strconv"
	"time"
)

type Runner struct {
	ID                    int    `gorm:"primaryKey"`
	Username              string `gorm:"unique_index"`
	Password              string
	PayPassword           string              `gorm:"default:123123"` //提现密码
	AgencyRunnerId        int                 //代理账号
	InvitationCode        string              `gorm:"unique_index"` //邀请码
	RegisterTime          int64               //注册时间
	Status                int                 `gorm:"default:1"` //状态 1正常 2封
	Superior              int                 //上级 玩家
	LeverTree             string              //上级树  @ 分割
	CollectionPoint       float64             `gorm:"type:decimal(10,2)"`           //代收赢利点
	PayPoint              float64             `gorm:"type:decimal(10,2)"`           //代付盈利点
	CashPledge            float64             `gorm:"type:decimal(10,2);default:0"` //押金
	CollectionLimit       float64             `gorm:"type:decimal(10,2);default:0"` //代收额度
	FreezeCollectionLimit float64             `gorm:"type:decimal(10,2);default:0"` //冻结代收额度
	PayLimit              float64             `gorm:"type:decimal(10,2);default:0"` //代付额度
	Commission            float64             `gorm:"type:decimal(10,2);default:0"` //佣金
	FreezeMoney           float64             `gorm:"type:decimal(10,2);default:0"` //提现冻结金额
	JuniorPoint           float64             `gorm:"type:decimal(10,2)"`           //下级税点
	WithdrawCommission    float64             `gorm:"type:decimal(10,2);default:0"` //提现手续费
	Token                 string              `gorm:"unique_index"`                 //Token   唯一标识   长度  48位
	Balance               float64             `gorm:"type:decimal(10,2);default:0"` //玩家余额
	LastLoginTime         int64               //最后一次登陆时间
	LastLoginIp           string              //最后一次登录的ip
	LastLoginRegion       string              //最后一次登录地区
	PaySwitch             int                 `gorm:"default:2"`  //代付开关   1开  2关
	ActiveGrade           int                 `gorm:"default:60"` //活跃分数
	CreditScore           int                 `gorm:"default:60"` //信用分
	Working               int                 `gorm:"default:1"`  //1  待机   2  工作     //工作状态
	CollectionTime        int                 `gorm:"default:0"`  //代收次数
	Remark                string              `gorm:"-"`
	LastGetOrderTime      int64               `gorm:"default:0"` //最后一次获取订单的时间  也就是我要关闭  接单的按钮
	Col                   modelPay.Collection `gorm:"-"`
	ChangeBalance         float64             `gorm:"-"` //变化的余额
	ChangeCommission      float64             `gorm:"-"`
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

func (r *Runner) GetRunnerUsername(db *gorm.DB) string {
	db.Where("id=?", r.ID).First(r)
	return r.Username
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

// CheckInvitationCode 检查邀请码
func (r *Runner) CheckInvitationCode(db *gorm.DB) (*Runner, error) {
	if len(r.InvitationCode) == 8 {
		//上级是会员
		runner := Runner{}
		err := db.Where("invitation_code=?", r.InvitationCode).First(&runner).Error
		if err != nil {
			return r, eeor.OtherError("The invitation code is invalid or expired")
		}
		r.Superior = runner.ID
		r.LeverTree = runner.LeverTree + strconv.Itoa(r.Superior) + "@"
		r.AgencyRunnerId = runner.AgencyRunnerId
		//查询这个代理
		agencyRunner := AgencyRunner{}
		err = db.Where("id=?", runner.AgencyRunnerId).First(&agencyRunner).Error
		if err != nil {
			return r, eeor.OtherError("The invitation code is invalid or expired")
		}
		r.CollectionPoint = agencyRunner.JuniorPoint
		r.PayPoint = agencyRunner.JuniorPoint
		r.WithdrawCommission = agencyRunner.JuniorWithdrawCommission
		return r, nil
	} else if len(r.InvitationCode) == 6 {
		//上级代理
		agencyRunner := AgencyRunner{}
		err := db.Where("invitation_code=?", r.InvitationCode).First(&agencyRunner).Error
		if err != nil {
			return r, eeor.OtherError("The invitation code is invalid or expired")
		}
		r.AgencyRunnerId = agencyRunner.ID
		r.CollectionPoint = agencyRunner.JuniorPoint
		r.PayPoint = agencyRunner.JuniorPoint
		r.WithdrawCommission = agencyRunner.JuniorWithdrawCommission
		return r, nil
	} else {
		return r, eeor.OtherError("The invitation code is invalid")
	}
}

// ChangeCashPledge 修改玩家 的押金
func (r *Runner) ChangeCashPledge(db *gorm.DB) error {
	//判断是否存在这个用户
	rTwo := Runner{}
	err := db.Where("id=? and  agency_runner_id=?", r.ID, r.AgencyRunnerId).First(&rTwo).Error
	if err != nil {
		return eeor.OtherError("The player does not exist")
	}
	ups := map[string]interface{}{}
	if rTwo.CashPledge+r.CashPledge < 0 {
		return eeor.OtherError("Insufficient deposit balance")
	}
	ups["CashPledge"] = rTwo.CashPledge + r.CashPledge
	//账变
	change := RunnerAmountChange{RunnerId: r.ID,
		NowAmount:    rTwo.CashPledge + r.CashPledge,
		ChangeAmount: r.CashPledge,
		FontAmount:   rTwo.CashPledge, Kinds: 1,
		Remark: r.Remark}
	err = change.Add(db)
	if err != nil {
		return err
	}
	affected := db.Model(&Runner{}).Where("id=? and  cash_pledge=? ", r.ID, rTwo.CashPledge).Update(ups).RowsAffected
	if affected == 0 {
		return eeor.OtherError("Failed to update")
	}
	return nil
}

// ChangeCollectionLimit 修改代收的额度     1.管理员操作  2玩家 接单    3订单失效释放订单   4收款少于订单金额    5管理让订单失败(退还玩家额度)
//  6  管理员让订单 支付成功
func (r *Runner) ChangeCollectionLimit(db *gorm.DB, IfSystem bool, kinds int) error {
	db = db.Begin()
	rr2 := Runner{}
	err := db.Where("id=?", r.ID).First(&rr2).Error
	if err != nil {
		return err
	}
	if r.CollectionLimit < 0 {
		if rr2.CollectionLimit < math.Abs(r.CollectionLimit) {
			return eeor.OtherError("The collection line is not enough")
		}
	}
	//自动代理操作 需要扣除代理的 代收额度
	if IfSystem {
		//判断是否够用
		ag := AgencyRunner{}
		db.Where("id=?", r.AgencyRunnerId).First(&ag)
		if r.CollectionLimit > 0 {
			if ag.CollectionLimit < r.CollectionLimit {
				return eeor.OtherError("The collection line is not enough")
			}
		}
		//修改余额
		apps := make(map[string]interface{})
		apps["CollectionLimit"] = ag.CollectionLimit - r.CollectionLimit
		apps["UseCollectionLimit"] = ag.UseCollectionLimit + r.CollectionLimit
		affected := db.Model(&AgencyRunner{}).Where("id=? and  collection_limit=?  and  use_collection_limit=?", r.AgencyRunnerId, ag.CollectionLimit, ag.UseCollectionLimit).Update(apps).RowsAffected
		if affected == 0 {
			return eeor.OtherError("Failed to update")
		}
		//添加账变
		change := AgencyAccountChange{
			ChangeAmount: r.CollectionLimit,
			NowAmount:    ag.CollectionLimit - r.CollectionLimit,
			FontAmount:   ag.CollectionLimit, Kinds: 2, AgencyRunnerId: r.AgencyRunnerId, RunnerId: r.ID, Remark: r.Remark}
		err = change.Add(db)
		if err != nil {
			db.Rollback()
			return err
		}
	}
	ups := make(map[string]interface{})
	ups["CollectionLimit"] = rr2.CollectionLimit + r.CollectionLimit
	//2玩家 接单
	if kinds == 2 {
		if IfSystem == false {
			ups["FreezeCollectionLimit"] = rr2.FreezeCollectionLimit + r.FreezeCollectionLimit
			ups["CollectionTime"] = rr2.CollectionTime + 1
			r.Remark = "关联订单:" + r.Col.OwnOrder + " 减少"
		}
		if IfSystem == false {
			//生成collection
			r.Col.RunnerId = r.ID
			r.Col.AgencyRunnerId = rr2.AgencyRunnerId
			r.Col.Species = 3
			err := r.Col.Add(db)
			if err != nil {
				db.Rollback()
				return err
			}
			//玩家统计
			statistics := RunnerStatistics{RunnerId: r.ID, AgencyRunnerId: r.AgencyRunnerId, CollectionAllAmount: r.Col.Amount, CollectionAllCount: 1}
			err = statistics.Add(db)
			if err != nil {
				db.Rollback()
				return err
			}
		}
	}
	//3订单失效释放订单
	if kinds == 3 {
		ups["FreezeCollectionLimit"] = rr2.FreezeCollectionLimit + r.FreezeCollectionLimit
		r.Remark = "关联订单:" + r.Col.OwnOrder + " 增加,原因订单过期未支付"
		//修改订单 状态
		affected := db.Model(&modelPay.Collection{}).Where("id=? and commission=?  and  status=1", r.Col.ID, r.Col.Commission).Update(&modelPay.Collection{Status: 4}).RowsAffected
		if affected == 0 {
			db.Rollback()
			return eeor.OtherError("Failed to update")
		}
	}
	//4收款少于订单金额
	if kinds == 4 {
		ups["FreezeCollectionLimit"] = rr2.FreezeCollectionLimit + r.FreezeCollectionLimit
		affected := db.Model(&modelPay.Collection{}).Where("id=?", r.Col.ID).Update(&modelPay.Collection{Status: 5, ActualAmount: r.Col.ActualAmount, Updated: time.Now().Unix()}).RowsAffected
		if affected == 0 {
			db.Rollback()
			return eeor.OtherError("Failed to update")
		}

	}
	//5管理让订单失败(退还玩家额度)
	if kinds == 5 {
		ups["FreezeCollectionLimit"] = rr2.FreezeCollectionLimit + r.FreezeCollectionLimit
		affected := db.Model(&modelPay.Collection{}).Where("id=?   and status=?", r.Col.ID, r.Col.Status).
			Update(&modelPay.Collection{
				Status:  3,
				Updated: time.Now().Unix(), Remark: r.Remark}).RowsAffected
		if affected == 0 {
			db.Rollback()
			return eeor.OtherError("Failed to update")
		}
	}
	// 管理员让订单 支付成功
	if kinds == 6 {
		ups["FreezeCollectionLimit"] = rr2.FreezeCollectionLimit + r.FreezeCollectionLimit
		merchant := Merchant{MerchantNum: r.Col.MerChantNum}
		err, _ := merchant.AmountChange(db, r.Col.ActualAmount, r.Col.ChannelId, r.Col.ID, r.Col.MerchantOrderNum, 3, r.Col)
		if err != nil {
			db.Rollback()
			return err
		}
	}
	//更新玩家的额度
	affected := db.Model(&Runner{}).Where("id=? and  collection_limit =? and  status =1 and  freeze_collection_limit=? ", r.ID, rr2.CollectionLimit, rr2.FreezeCollectionLimit).Update(ups).RowsAffected
	if affected == 0 {
		db.Rollback()
		return eeor.OtherError("Failed to update")
	}
	//生成账变
	change := RunnerAmountChange{RunnerId: r.ID,
		NowAmount:    rr2.CollectionLimit + r.CollectionLimit,
		ChangeAmount: r.CollectionLimit,
		FontAmount:   rr2.CollectionLimit,
		Remark:       r.Remark, Kinds: 2, CollectionId: r.Col.ID}
	err = change.Add(db)
	if err != nil {
		db.Rollback()
		return err
	}

	db.Commit()
	return nil
}

// SnagTheOrder 拉群订单代理操作
func (r *Runner) SnagTheOrder(db *gorm.DB, coll modelPay.Collection) (string, error) {
	//判断代收通道的
	upi := ""
	db.Raw("SELECT * FROM runners  LEFT JOIN  agency_runners   on  runners.agency_runner_id=agency_runners.id WHERE   agency_runners.collection_channel LIKE  ?  AND  runners.status=1 AND  agency_runners.status =1  and   runners.working=2   AND  runners.collection_limit   >=?  ORDER BY  runners.collection_time asc  LIMIT  1", "%"+strconv.Itoa(coll.ChannelId)+"%", coll.Amount).Scan(&r)
	if r.ID == 0 {
		return upi, eeor.OtherError("No upi was matched. Procedure")
	}
	//获取这个人的 upi地址
	UPP := RunnerUpi{}
	db.Where("runner_id=? and  kind=1", r.ID).First(&UPP)
	//修改 用户的余额
	r.CollectionLimit = -coll.Amount
	r.FreezeCollectionLimit = coll.Amount
	r.Col = coll
	r.Col.Upi = UPP.Address
	err := r.ChangeCollectionLimit(db, false, 2)
	if err != nil {
		return "", err
	}

	upi = UPP.Address
	return upi, nil
}

// ChangeCommissionAndBalance 修改账户余额
func (r *Runner) ChangeCommissionAndBalance(db *gorm.DB) error {
	ups := make(map[string]interface{})
	db1 := db
	db = db.Model(&Runner{}).Where("id=?", r.ID)
	if r.Balance != 0 {
		ups["Balance"] = r.Balance + r.ChangeBalance
		db = db.Where("balance=?", r.Balance)
		//添加账变(余额)
		change := RunnerAmountChange{RunnerId: r.ID,
			Remark:       r.Remark,
			NowAmount:    r.Balance + r.ChangeBalance,
			ChangeAmount: r.ChangeBalance,
			FontAmount:   r.Balance, Kinds: 6}
		err := change.Add(db1)
		if err != nil {
			return err
		}
	}
	if r.Commission != 0 {
		ups["Commission"] = r.Commission + r.ChangeCommission
		db = db.Where("commission=?", r.Commission)
		//添加账变(佣金)
		change := RunnerAmountChange{RunnerId: r.ID,
			Remark:       r.Remark,
			NowAmount:    r.Balance + r.ChangeBalance,
			ChangeAmount: r.ChangeBalance,
			FontAmount:   r.Balance, Kinds: 6}
		err := change.Add(db1)
		if err != nil {
			return err
		}
	}
	affected := db.Update(ups).RowsAffected
	if affected == 0 {
		return eeor.OtherError("Failed to update")
	}
	return nil
}

type Striking struct {
	RunnerId       int
	AgencyRunnerId int
	Col            modelPay.Collection
	Rate           float64
	Kinds          int //类型  1代收 2代付
	Symbol         float64
}

// Striking 代理冲正处理
func (sk *Striking) Striking(db *gorm.DB) error {
	//判断这个代理是否是顶级代理
	runner := Runner{}
	err := db.Where("id=?", sk.RunnerId).First(&runner).Error
	if err != nil {
		return eeor.OtherError("runnerId  is fail")
	}
	if sk.Kinds == 1 {
		sk.Rate = runner.CollectionPoint
	} else {
		sk.Rate = runner.PayPoint
	}
	//修改账户余额(本账号)
	runner.ChangeBalance = sk.Rate * sk.Col.ActualAmount * sk.Symbol
	runner.ChangeCommission = sk.Rate * sk.Col.ActualAmount * sk.Symbol
	err = runner.ChangeCommissionAndBalance(db)
	if err != nil {
		return err
	}
	//奔跑者 统计发生变化
	statistics := RunnerStatistics{}
	if sk.Kinds == 1 {
		statistics.Commission = sk.Rate * sk.Col.ActualAmount * sk.Symbol
		statistics.CollectionAmount = sk.Rate * sk.Col.ActualAmount * sk.Symbol
		statistics.CollectionCount = int(sk.Symbol) * 1
	}
	if runner.LeverTree != "" { //说明不是顶级代理(要和上级代理进行 费率对比)
		runner2 := Runner{}
		err := db.Where("id=?", runner.Superior).First(&runner2).Error
		if err != nil {
			return eeor.OtherError("runnerId2  is fail")
		}
		if sk.Kinds == 1 {
			sk.Rate = runner2.CollectionPoint - sk.Rate
		} else {
			sk.Rate = runner2.PayPoint - sk.Rate
		}

		if sk.Rate > 0 {
			runner2.ChangeBalance = sk.Rate * sk.Col.ActualAmount * sk.Symbol
			runner2.ChangeCommission = sk.Rate * sk.Col.ActualAmount * sk.Symbol
			err = runner2.ChangeCommissionAndBalance(db)
			if err != nil {
				return err
			}
			//日统计成功==>减去
			statistics := RunnerStatistics{}
			if sk.Kinds == 1 {
				statistics.Commission = sk.Rate * sk.Col.ActualAmount * sk.Symbol
				statistics.CollectionAmount = sk.Rate * sk.Col.ActualAmount * sk.Symbol
				statistics.CollectionCount = int(sk.Symbol) * 1
			}
			err := statistics.Add(db)
			if err != nil {
				return err
			}
		}
		if sk.Kinds == 1 {
			sk.Rate = runner2.CollectionPoint
		} else {
			sk.Rate = runner2.PayPoint
		}
	}
	//代操作内容
	agg := AgencyRunner{}
	err = db.Where("id=?", sk.AgencyRunnerId).First(&agg).Error
	if err != nil {
		return eeor.OtherError("AgencyRunnerId  is fail")
	}
	if sk.Kinds == 1 {
		sk.Rate = agg.CollectionPoint - sk.Rate
		//更新  每日统计
	} else {
		sk.Rate = agg.PayPoint - sk.Rate
	}

	if sk.Rate > 0 { //处理代理的账变
		agencyRunner := AgencyRunner{
			Balance: sk.Symbol * sk.Rate * sk.Col.ActualAmount, Commission: sk.Symbol * sk.Rate * sk.Col.ActualAmount}
		err := agencyRunner.ChangeCommissionAndBalance(db)
		if err != nil {
			return err
		}

	}

	return nil
}
