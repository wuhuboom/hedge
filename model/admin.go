package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/tools"
	"time"
)

type Admin struct {
	ID            int    `gorm:"primaryKey" `
	AdminUsername string // 管理用户名
	AdminPassword string //管理密码
	GoogleCode    string //谷歌验证码
	Token         string //  管理员token
	Created       int64
	Profit        float64 `gorm:"type:decimal(10,2);default:0"` // 盈利金额
	Deposit       float64 `gorm:"type:decimal(10,2);default:0"` //代理押金操作

	CollectionId int `gorm:"-"`
}

func CheckIsExistModelAdminAdmin(db *gorm.DB) {
	if db.HasTable(&Admin{}) {
		fmt.Println("数据库已经存在了!")
		db.AutoMigrate(&Admin{})
	} else {
		fmt.Println("数据不存在,所以我要先创建数据库")
		db.CreateTable(&Admin{})
		db.Save(&Admin{AdminUsername: "jiabanjushi", AdminPassword: "8e3db6345135296ba2dde6f45c2519cd", Created: time.Now().Unix(), Token: tools.RandStringRunes(32)})
	}
}

func (ad *Admin) ChangeProfit(db *gorm.DB) error {
	aa := Admin{}
	db.Where("id=?", 1).First(&aa)
	ups := make(map[string]interface{})
	ups["Profit"] = aa.Profit + ad.Profit
	//账变
	change := AdminAccountChange{NowAmount: aa.Profit + ad.Profit,
		FontAmount:   aa.Profit,
		ChangeAmount: ad.Profit, Kinds: 1,
		CollectionId: ad.CollectionId}
	err := change.Add(db)
	if err != nil {
		return err
	}
	affected := db.Model(&Admin{}).Where("id=? and profit=?", 1, aa.Profit).Update(ups).RowsAffected
	if affected == 0 {
		return eeor.OtherError("Failed to update")
	}
	return nil
}

type Commission struct {
	RunnerId       int
	AgencyRunnerId int
	Commission     float64
	ActualAmount   float64
	CollectionId   int
	MerRate        float64 // 商户的  费率
	Species        int     // 1三方单子 2对冲  3跑单
	PayType        int     //1  代收  1代付
}

func (co *Commission) ChangeCommission(db *gorm.DB) error {
	//三方普通单子
	if co.Species == 1 {
		admin := Admin{Profit: co.Commission, CollectionId: co.CollectionId}
		err := admin.ChangeProfit(db)
		if err != nil {
			return err
		}
	}
	if co.Species == 3 {
		//判断奔跑者是否存在上级
		runner := Runner{}
		err := db.Where("id=?", co.RunnerId).First(&runner).Error
		if err != nil {
			return err
		}
		//不存在(代理就是其直属上级)
		if co.PayType == 1 {
			//更新玩家
			var rate float64
			rate = runner.CollectionPoint
			affected := db.Model(&Runner{}).Where("id=? and  commission=?  and  balance=?", runner.ID, runner.Commission, runner.Balance).
				Update(map[string]interface{}{"Commission": runner.Commission + rate*co.ActualAmount, "Balance": runner.Balance + rate*co.ActualAmount}).RowsAffected

			if affected == 0 {
				return eeor.OtherError("update is fail")
			}

			//生成账变(佣金账变)
			change := RunnerAmountChange{RunnerId: runner.ID,
				Kinds: 4, CollectionId: co.CollectionId,
				NowAmount: runner.Commission + rate*co.ActualAmount, FontAmount: runner.Commission, ChangeAmount: rate * co.ActualAmount}
			err := change.Add(db)
			if err != nil {
				return err
			}
			//生成账变
			amountChange := RunnerAmountChange{RunnerId: runner.ID, Kinds: 6, CollectionId: co.CollectionId,
				Remark: "rebate", NowAmount: runner.Balance + runner.CollectionPoint*co.ActualAmount, FontAmount: runner.Balance, ChangeAmount: runner.CollectionPoint * co.ActualAmount}
			err = amountChange.Add(db)

			if err != nil {
				return err
			}
			statistics := RunnerStatistics{RunnerId: runner.ID,
				AgencyRunnerId: runner.AgencyRunnerId, CollectionAmount: co.ActualAmount, CollectionCount: 1, Commission: rate * co.ActualAmount}
			err = statistics.Add(db)
			if err != nil {
				return err
			}

			//存在上级
			if runner.Superior != 0 {
				runner2 := Runner{}
				rate = runner2.CollectionPoint - rate
				err = db.Where("id=?", runner.Superior).First(&runner2).Error
				if err != nil {
					return err
				}
				affected = db.Model(&Runner{}).Where("id=? and  commission=?  and  balance=?", runner2.ID, runner2.Commission, runner2.Balance).
					Update(map[string]interface{}{"Commission": runner2.Commission + rate*co.ActualAmount, "Balance": runner2.Balance + rate*co.ActualAmount}).RowsAffected
				if affected == 0 {
					return eeor.OtherError("update is fail")
				}
				//生成账变(佣金账变)
				change := RunnerAmountChange{RunnerId: runner2.ID,
					Kinds: 4, CollectionId: co.CollectionId,
					NowAmount: runner.Commission + runner2.CollectionPoint*co.ActualAmount, FontAmount: runner2.Commission, ChangeAmount: runner2.CollectionPoint * co.ActualAmount}
				err = change.Add(db)
				if err != nil {
					return err
				}
				//生成账变
				amountChange := RunnerAmountChange{RunnerId: runner2.ID, Kinds: 6, CollectionId: co.CollectionId,
					Remark: "rebate", NowAmount: runner2.Balance + runner2.CollectionPoint*co.ActualAmount, FontAmount: runner2.Balance, ChangeAmount: runner2.CollectionPoint * co.ActualAmount}
				err = amountChange.Add(db)
				if err != nil {
					return err
				}

				statistics := RunnerStatistics{RunnerId: runner.ID,
					AgencyRunnerId: runner.AgencyRunnerId, CollectionAmount: co.ActualAmount, CollectionCount: 1, Commission: rate * co.ActualAmount}
				err = statistics.Add(db)
				if err != nil {
					return err
				}
				rate = runner2.CollectionPoint
			}

			//更新代理
			agencyRunner := AgencyRunner{}
			err = db.Where("id=?", co.AgencyRunnerId).First(&agencyRunner).Error
			if err != nil {
				return err
			}

			rate = agencyRunner.CollectionPoint - rate
			affected = db.Model(&AgencyRunner{}).Where("id=? and  commission=?", agencyRunner.ID, agencyRunner.Commission).
				Update(map[string]interface{}{"Commission": agencyRunner.Commission + rate*co.ActualAmount}).RowsAffected
			if affected == 0 {
				return eeor.OtherError("update is fail")
			}
			//账变
			accountChange := AgencyAccountChange{
				RunnerId:       co.RunnerId,
				AgencyRunnerId: co.AgencyRunnerId,
				NowAmount:      agencyRunner.Commission + rate*co.ActualAmount,
				FontAmount:     agencyRunner.Commission,
				ChangeAmount:   rate * co.ActualAmount, CollectionId: co.CollectionId, Kinds: 4}
			err = accountChange.Add(db)
			if err != nil {
				return err
			}
			//更新管理员
			rate = co.MerRate - agencyRunner.CollectionPoint
			admin := Admin{Profit: rate * co.ActualAmount, CollectionId: co.CollectionId}
			err = admin.ChangeProfit(db)
			if err != nil {
				return err
			}

		}
	}

	return nil
}
