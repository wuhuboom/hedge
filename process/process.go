package process

import (
	"github.com/jinzhu/gorm"
	jsoniter "github.com/json-iterator/go"
	"github.com/wangyi/GinTemplate/common"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
)

// CheckExpirationPayOrder 检查过期进程
func CheckExpirationPayOrder(db *gorm.DB) {
	for true {
		pay := make([]model.PayOrder, 0)
		db.Where("expire_time  <? and status= ?", time.Now().Unix(), 1).Find(&pay)
		for _, order := range pay {
			db = db.Begin()
			common.OrderLuck.RLock()
			paid := model.PaidOrder{}
			err := db.Where("id=?", order.PaidOrderId).First(&paid).Error
			if err != nil {
				//没有找到这个代付订单
				common.OrderLuck.RUnlock() //读锁关闭
				db.Rollback()
				continue
			}
			common.OrderLuck.RUnlock() //读锁关闭
			common.OrderLuck.Lock()    //写锁开启
			update := make(map[string]interface{})
			update["NeedWithdrawn"] = paid.NeedWithdrawn + order.Amount
			update["FreezeWithdrawn"] = paid.FreezeWithdrawn - order.Amount
			err = db.Model(&model.PaidOrder{}).Where("id=?", paid.ID).Update(update).Error
			if err != nil {
				common.OrderLuck.Unlock()
				db.Rollback()
				continue
			}
			notice := make(map[string]string)
			notice["PayOrderNum"] = order.ThreeOrder                         //订单号
			notice["amount"] = strconv.FormatFloat(order.Amount, 'f', 2, 64) //金额  退还金额(重新放回)
			notice["action"] = "fail"                                        //  冻结金额
			notice["PaidOrderNum"] = paid.ThreeOrder                         //订单号
			marshal, err := jsoniter.Marshal(notice)
			if err != nil {
				common.OrderLuck.Unlock()
				db.Rollback()
				return
			}
			go tools.HttpJsonPost(paid.NoticeUrl, marshal)
			db.Commit()
			common.OrderLuck.Unlock() //关闭写锁
			time.Sleep(time.Second * 2)
		}
		time.Sleep(time.Second * 30)
	}
}

func OverdueCollection(db *gorm.DB) {
	for true {
		db.Model(&modelPay.Collection{}).Where("status=?  and  species  =1 and expire_time  < ?", 1, time.Now().Unix()).Update(&modelPay.Collection{Status: 4})
		time.Sleep(time.Second * 30)
	}
}

// ExpireCollection 过期订单
func ExpireCollection(db *gorm.DB) {

	for true {
		col := make([]modelPay.Collection, 0)
		db.Where("species= 3  and  status=1 and kinds=1  and expire_time<?", time.Now().Unix()).Find(&col)
		for _, collection := range col {
			runner := model.Runner{ID: collection.RunnerId}
			runner.Col.ID = collection.ID
			runner.CollectionLimit = collection.Amount
			runner.FreezeCollectionLimit = -collection.Amount
			runner.ChangeCollectionLimit(db, false, 3)

		}
		time.Sleep(time.Second * 30)
	}
}

// CheckLastGetOrderTime 离线玩家接单
func CheckLastGetOrderTime(db *gorm.DB) {
	for true {
		db.Model(&model.Runner{}).Where("working=2 and last_get_order_time < ?", time.Now().Unix()).Update(&model.Runner{Working: 1})
		time.Sleep(time.Second * 60)
	}
}
