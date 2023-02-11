package management

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
)

// CollectionOperation 操作订单
func CollectionOperation(c *gin.Context) {
	action := c.Query("action")
	who, _ := c.Get("who")
	whoMap := who.(model.Admin)
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]modelPay.Collection, 0)

		db := mysql.DB.Where("kinds=?", c.PostForm("kinds"))
		var total int

		if status, IsE := c.GetPostForm("merchant_num"); IsE == true {
			db = db.Where("mer_chant_num=?", status)
		}
		//条件
		if status, IsE := c.GetPostForm("status"); IsE == true {
			db = db.Where("status=?", status)
		}
		if status, IsE := c.GetPostForm("callback"); IsE == true {
			db = db.Where("callback=?", status)
		}

		if start, IsE := c.GetPostForm("start"); IsE == true {
			if end, IsE := c.GetPostForm("end"); IsE == true {
				db = db.Where("created  <=  ? and created >=  ?", end, start)
			}
		}
		//商家订单号
		if status, IsE := c.GetPostForm("merchant_order_num"); IsE == true {
			db = db.Where("merchant_order_num=?", status)
		}
		//平台订单号
		if status, IsE := c.GetPostForm("own_order"); IsE == true {
			db = db.Where("own_order=?", status)
		}

		//支付凭证  proof_of_payment
		if status, IsE := c.GetPostForm("proof_of_payment"); IsE == true {
			db = db.Where("proof_of_payment=?", status)
		}
		//用户名  runner_id(奔跑者)
		if username, isE := c.GetPostForm("username"); isE == true {
			runner := model.Runner{Username: username}
			db = db.Where("runner_id=?", runner.GetRunnerId(mysql.DB))
		}

		//填写代理名字
		if aUsername, isE := c.GetPostForm("agency_runner_name"); isE == true {
			runner := model.AgencyRunner{Username: aUsername}
			db = db.Where("agency_runner_id=?", runner.GetId(mysql.DB))
		}
		//单子类型
		if species, isE := c.GetPostForm("species"); isE == true {
			db = db.Where("species=?", species)
		}

		//upi
		if species, isE := c.GetPostForm("upi"); isE == true {
			db = db.Where("upi=?", species)
		}
		//channel_id  通道名字
		if species, isE := c.GetPostForm("channel_name"); isE == true {
			atom, _ := strconv.Atoi(species)
			channel := modelPay.Channel{ID: atom}
			channel.GetChannelId(mysql.DB)
			db = db.Where("channel_id=?", species)
		}
		db.Model(&modelPay.Collection{}).Count(&total)
		db = db.Model(&modelPay.Collection{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		for i, collection := range sl {
			bank := modelPay.Bank{}
			mysql.DB.Where("id=?", collection.BankId).First(&bank)
			banIn := modelPay.BankInformation{}
			mysql.DB.Where("id=?", bank.BankInformationId).First(&banIn)
			sl[i].BankNum = bank.CardNum
			sl[i].BankName = banIn.BankName
			channel := modelPay.Channel{ID: collection.ChannelId}
			sl[i].ChannelId = channel.GetChannelName(mysql.DB)
			runner := model.Runner{ID: collection.RunnerId}
			sl[i].RunnerName = runner.GetRunnerUsername(mysql.DB)
		}

		tools.ReturnDataLIst2000(c, sl, total)
		return
	}
	//   操作订单状态 (失败/成功)
	if action == "confirmationOfPayment" {
		id := c.PostForm("id")
		col := modelPay.Collection{}
		err := mysql.DB.Where("id=?", id).First(&col).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		result, _ := redis.Rdb.Get("confirmationOfPayment" + id).Result()
		if result != "" {
			tools.ReturnErr101Code(c, "Don't do the deposit operation at the specified time")
			return
		}

		//  status  1等待支付  2支付成功  3失败
		sta, _ := strconv.Atoi(c.PostForm("status"))
		if sta != 2 && col.Kinds == 1 {
			tools.ReturnErr101Code(c, "Illegal request")
			return
		}
		if sta == col.Status {
			tools.ReturnErr101Code(c, "Don't repeat changes")
			return
		}

		//查询商家
		mer := model.Merchant{}
		err = mysql.DB.Where(" merchant_num= ?", col.MerChantNum).First(&mer).Error
		if err != nil {
			tools.ReturnErr101Code(c, "Businesses don't exist")
			return
		}

		//回调给 三方t
		if sta == 3 {
			if tools.IsChinese(c.PostForm("remark")) {
				tools.ReturnErr101Code(c, "remark is  not  Chinese")
				return
			}
			//跑分 处理逻辑  // 充值失败
			if col.Species == 3 {
				//失败  必须要就订单 以及过期
				//订单被管理驳回了    代理玩家要退还额度
				runner := model.Runner{ID: col.RunnerId}
				runner.CollectionLimit = col.ActualAmount
				runner.FreezeCollectionLimit = -col.ActualAmount
				runner.Col.ID = col.ID
				runner.Col.Status = col.Status
				err := runner.ChangeCollectionLimit(mysql.DB, false, 5)
				if err != nil {
					tools.ReturnErr101Code(c, err.Error())
					return
				}
			} else {
				//代付订单
				if col.Kinds == 2 {
					//逻辑处理  1.修改订单状态  2.返回商户的可用额度
					db := mysql.DB.Begin()
					if err := db.Model(&modelPay.Collection{}).Where("id=? and status=?", col.ID, col.Status).Update(&modelPay.Collection{
						Status: 3, Updated: time.Now().Unix()}).Error; err != nil {
						tools.ReturnErr101Code(c, err.Error())
						return
					}
					//修改商户
					ups := make(map[string]interface{})
					ups["AvailableAmount"] = mer.AvailableAmount + col.Commission + col.Amount
					ups["FreezeAmount"] = mer.FreezeAmount - (col.Commission + col.Amount)
					err := db.Model(&model.Merchant{}).Where("id=? and  available_amount  =? and freeze_amount=? ", mer.ID, mer.AvailableAmount, mer.FreezeAmount).Update(ups).Error
					if err != nil {
						db.Rollback()
						tools.ReturnErr101Code(c, err.Error())
						return
					}

					//账变
					change := modelPay.AmountChange{
						MerchantNum:  mer.MerchantNum,
						CollectionId: col.ID,
						Amount:       col.Commission + col.Amount,
						Before:       mer.AvailableAmount,
						After:        mer.AvailableAmount + col.Commission + col.Amount, Kinds: 1, Status: 2, Remark: "关联订单:" + col.OwnOrder}
					err = change.Add(db)
					if err != nil {
						db.Rollback()
						tools.ReturnErr101Code(c, err.Error())
						return
					}

					db.Commit()
				} else {
					// 正常三方 逻辑

					if err := mysql.DB.Model(&modelPay.Collection{}).Where("id=? and status=?", col.ID, col.Status).Update(&modelPay.Collection{
						Status: 3, Updated: time.Now().Unix()}).Error; err != nil {
						tools.ReturnErr101Code(c, err.Error())
						return
					}
				}

			}

		} else { //订单处理成功
			ActualAmount, _ := strconv.ParseFloat(c.PostForm("actual_amount"), 64)
			if ActualAmount <= 0 {
				tools.ReturnErr101Code(c, "actual_amount must  have")
				return
			}
			if ActualAmount > col.Amount {
				tools.ReturnErr101Code(c, "The actual amount cannot be greater than the order amount")
				return
			}

			if col.Status == 3 {
				tools.ReturnErr101Code(c, "the order status is fail")
				return
			}

			//充值成功  代收逻辑
			if col.Kinds == 2 {
				if col.Species == 3 {
					//代理收益逻辑

				} else { //1.修改订单号成功  2.超管加钱 生成账变    3商户号修改冻结金额
					db := mysql.DB.Begin()
					err := db.Model(&modelPay.Collection{}).Where("id=? and status=? and  kinds=2", col.ID, col.Status).Update(&modelPay.Collection{Status: 2, Updated: time.Now().Unix()}).Error

					if err != nil {
						tools.ReturnErr101Code(c, err.Error())
						return
					}
					admin := model.Admin{ID: whoMap.ID, Profit: col.Commission, CollectionId: col.ID}
					err = admin.ChangeProfit(db)
					if err != nil {
						tools.ReturnErr101Code(c, err.Error())

						db.Rollback()
						return
					}
					ActualAmount = col.Amount
					err = db.Model(&model.Merchant{}).Where("id=? and  freeze_amount=?", mer.ID, mer.FreezeAmount).
						Update(map[string]interface{}{"FreezeAmount": mer.FreezeAmount - col.Amount - col.Commission}).Error

					//zap.L().Debug(fmt.Sprintf("原始金额:%f,变化后的金额:%f,变化订单:%s", mer.FreezeAmount, mer.FreezeAmount-col.Amount-col.Commission, col.OwnOrder))
					if err != nil {
						tools.ReturnErr101Code(c, err.Error())
						db.Rollback()
						return
					}
					//修改每天统计
					statistics := modelPay.Statistics{MerchantNum: col.MerChantNum, TodayPay: 1, TodayPayAmount: ActualAmount, TodayPayCommission: col.Commission}
					err = statistics.Add(db, 2)
					if err != nil {
						db.Rollback()
						tools.ReturnErr101Code(c, err.Error())
						return
					}
					db.Commit()

				}

			} else {
				if col.Species == 3 {
					//跑分逻辑
					runner := model.Runner{ID: col.RunnerId}
					runner.CollectionLimit = 0
					runner.FreezeCollectionLimit = -ActualAmount
					runner.Col = col
					err := runner.ChangeCollectionLimit(mysql.DB, false, 6)
					if err != nil {
						tools.ReturnErr101Code(c, err.Error())
						return
					}
				} else { //普通三方逻辑
					merchant := model.Merchant{MerchantNum: col.MerChantNum}
					err, _ := merchant.AmountChange(mysql.DB, ActualAmount, col.ChannelId, col.ID, col.OwnOrder, 1, col)
					if err != nil {
						tools.ReturnErr101Code(c, err.Error())
						return
					}
				}
			}
		}
		callback := map[string]string{}
		callback["merchant_order_num"] = col.MerchantOrderNum
		callback["channel_id"] = strconv.Itoa(col.ChannelId)
		callback["status"] = strconv.Itoa(sta)
		callback["remark"] = c.PostForm("remark")
		callback["amount"] = strconv.FormatFloat(col.Amount, 'f', 2, 64)
		callback["actual_amount"] = c.PostForm("actual_amount")
		callback["currency"] = col.Currency
		callback["commission"] = strconv.FormatFloat(col.Commission, 'f', 2, 64)
		fmt.Println(tools.AsciiKey(callback) + "&key=" + mer.ApiKey)
		callback["sign"] = tools.MD5(tools.AsciiKey(callback) + "&key=" + mer.ApiKey)
		marshal, err := jsoniter.Marshal(callback)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		returnData := tools.HttpJsonPostOne(col.NoticeUrl, marshal)
		if len(returnData) > 200 {
			returnData = "url is  fail"
		}
		up := &modelPay.Collection{
			CallbackContent: string(marshal),
			CallbackReturn:  returnData,
		}
		if returnData == "SUCCESS" {
			up.Callback = 2
		} else {
			up.Callback = 3
		}
		mysql.DB.Model(&modelPay.Collection{}).Where("id=?", col.ID).Update(up)
		redis.Rdb.Set("confirmationOfPayment"+id, "123", 5*time.Second)
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}
	//回调
	if action == "callback" {

		id := c.PostForm("id")
		col := modelPay.Collection{}
		err := mysql.DB.Where("id=?", id).First(&col).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}

		mer := model.Merchant{}
		err = mysql.DB.Where(" merchant_num= ?", col.MerChantNum).First(&mer).Error
		if err != nil {
			tools.ReturnErr101Code(c, "Businesses don't exist")
			return
		}

		if mer.Status != 2 {
			tools.ReturnErr101Code(c, "The order has not been paid")
			return
		}

		callback := map[string]string{}
		callback["merchant_order_num"] = col.MerchantOrderNum
		callback["channel_id"] = strconv.Itoa(col.ChannelId)
		callback["status"] = strconv.Itoa(col.Status)
		callback["remark"] = c.PostForm("remark")
		callback["amount"] = strconv.FormatFloat(col.Amount, 'f', 2, 64)
		callback["actual_amount"] = strconv.FormatFloat(col.ActualAmount, 'f', 2, 64)
		callback["currency"] = col.Currency
		callback["commission"] = strconv.FormatFloat(col.Commission, 'f', 2, 64)
		callback["sign"] = tools.MD5(tools.AsciiKey(callback) + "&key=" + mer.ApiKey)
		marshal, err := jsoniter.Marshal(callback)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		returnData := tools.HttpJsonPostOne(col.NoticeUrl, marshal)
		if len(returnData) > 200 {
			returnData = "url is  fail"
		}
		up := &modelPay.Collection{
			CallbackContent: string(marshal),
			CallbackReturn:  returnData,
			Updated:         time.Now().Unix(),
		}
		if returnData == "SUCCESS" {
			up.Callback = 2
		} else {
			up.Callback = 3
		}
		mysql.DB.Model(&modelPay.Collection{}).Where("id=?", col.ID).Update(up)
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}
	if action == "one" {
		id := c.PostForm("id")
		col := modelPay.Collection{}
		err2 := mysql.DB.Where("id=?", id).First(&col).Error
		if err2 != nil {
			tools.ReturnErr101Code(c, err2.Error())
			return
		}
		CH := modelPay.Channel{}
		err := mysql.DB.Where("id=?", col.ChannelId).First(&CH).Error
		if err == nil {
			col.ChannelId = CH.ChannelName
		}
		tools.ReturnSuccess2000DataCode(c, col, "OK")
		return

	}

	//订单冲正
	if action == "striking" {
		id := c.PostForm("id")
		col := modelPay.Collection{}
		err := mysql.DB.Where("id=?", id).First(&col).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		mer := model.Merchant{}
		err = mysql.DB.Where("merchant_num=?", col.MerChantNum).First(&mer).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		ch := modelPay.Channel{}
		err = mysql.DB.Where("id=?", col.ChannelId).First(&ch).Error
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		if col.Status != 2 && col.Status != 3 {
			tools.ReturnErr101Code(c, "status  is fail")
			return
		}
		//判断是否是跑分订单
		if col.Species == 1 {
			//判断是代付订单 还是代付订单
			if col.Kinds == 1 {
				if col.Status != 2 {
					tools.ReturnErr101Code(c, "the order's status is not  right")
					return
				}
				//减少总金额   可用金额  代收成功数量  代收成功金额   超管的盈利增加?
				db := mysql.DB.Begin()
				err := db.Model(&modelPay.Collection{}).Where("id=? and status=? ", id, col.Status).Update(&modelPay.Collection{Status: 7, Updated: time.Now().Unix()}).Error
				if err != nil {
					tools.ReturnErr101Code(c, err.Error())
					return
				}
				ups := make(map[string]interface{})
				ups["AvailableAmount"] = mer.AvailableAmount - (col.ActualAmount - col.Commission)
				ups["AllAmount"] = mer.AllAmount - col.ActualAmount
				err = db.Model(&model.Merchant{}).Where("id=?  and  available_amount=? and  all_amount=? ", mer.ID, mer.AvailableAmount, mer.AllAmount).Update(ups).Error
				if err != nil {
					db.Rollback()
					tools.ReturnErr101Code(c, err.Error())
					return
				}
				//新增账变
				change := modelPay.AmountChange{
					MerchantNum: mer.MerchantNum,
					Before:      mer.AvailableAmount,
					Amount:      -(col.ActualAmount - col.Commission),
					After:       mer.AvailableAmount - (col.ActualAmount - col.Commission), CollectionId: col.ID, Remark: "冲正:" + col.OwnOrder}
				err = change.Add(db)
				if err != nil {
					db.Rollback()
					tools.ReturnErr101Code(c, err.Error())
					return
				}
				// 回退 每日统计
				statistics := modelPay.Statistics{
					TodayCollection:           -1,
					TodayCollectionCommission: -col.Commission,
					TodayCollectionAmount:     -col.ActualAmount, MerchantNum: mer.MerchantNum}
				err = statistics.Add(db, 1)
				if err != nil {
					db.Rollback()
					tools.ReturnErr101Code(c, err.Error())
					return
				}

				//管理员的
				admin := model.Admin{Profit: -col.Commission, CollectionId: col.ID}
				err = admin.ChangeProfit(db)
				if err != nil {
					db.Rollback()
					tools.ReturnErr101Code(c, err.Error())
					return
				}
				db.Commit()
				tools.ReturnSuccess2000Code(c, "OK")
				return

			} else { //代付订单冲正
				if col.Status == 3 { //失败的冲正(实际上是代付成功了)
					//实际上是代付成功的所以要直接扣除可用余额 1.商户要扣钱,并且生成账变  2.日统计要发现变化  3.超级管理员要账变
					if mer.AvailableAmount-(col.Amount+col.Commission) < 0 {
						tools.ReturnErr101Code(c, "The balance of the account is not enough")
						return
					}
					db := mysql.DB.Begin()
					err := db.Model(&modelPay.Collection{}).Where("id=? and status=? ", id, col.Status).Update(&modelPay.Collection{Status: 8, Updated: time.Now().Unix()}).Error
					if err != nil {
						tools.ReturnErr101Code(c, err.Error())
						return
					}
					mer.ChangeAvailableAmount = -(col.Amount + col.Commission)
					col.Remark = "冲正:" + col.OwnOrder
					mer.Col = col
					err = mer.AvailableAmountChangeAndFreezeAmount(db)
					if err != nil {
						db.Rollback()
						tools.ReturnErr101Code(c, err.Error())
						return
					}
					//统计发现变化?  代付成功要统计上去
					statistics := modelPay.Statistics{MerchantNum: mer.MerchantNum, TodayPay: 1, TodayPayAmount: col.Amount, TodayPayCommission: col.Commission}
					err = statistics.Add(db, 2)
					if err != nil {
						db.Rollback()
						tools.ReturnErr101Code(c, err.Error())
						return
					}

					//管理的账号
					admin := model.Admin{Profit: col.Commission, CollectionId: col.ID}
					err = admin.ChangeProfit(db)
					if err != nil {
						db.Rollback()
						tools.ReturnErr101Code(c, err.Error())
						return
					}

					//订单状态修改

					db.Commit()
					tools.ReturnSuccess2000Code(c, "OK")
					return
				} else if col.Status == 2 { //成功的冲正(按道理来说是失败的)
					// 1.修改商户余额   2.每日数据修改  3.
					db := mysql.DB.Begin()
					err := db.Model(&modelPay.Collection{}).Where("id=? and status=? ", id, col.Status).Update(&modelPay.Collection{Status: 7, Updated: time.Now().Unix()}).Error
					if err != nil {
						tools.ReturnErr101Code(c, err.Error())
						return
					}

					mer.ChangeAvailableAmount = col.Amount + col.Commission
					col.Remark = "冲正:" + col.OwnOrder
					mer.Col = col
					err = mer.AvailableAmountChangeAndFreezeAmount(db)
					if err != nil {
						db.Rollback()
						tools.ReturnErr101Code(c, err.Error())
						return
					}
					//统计发现变化?  减去代付成功的-
					statistics := modelPay.Statistics{MerchantNum: mer.MerchantNum,
						TodayPay: -1, TodayPayAmount: -col.Amount, TodayPayCommission: -col.Commission}
					err = statistics.Add(db, 2)
					if err != nil {
						db.Rollback()
						tools.ReturnErr101Code(c, err.Error())
						return
					}

					//管理的账号
					admin := model.Admin{Profit: -col.Commission, CollectionId: col.ID}
					err = admin.ChangeProfit(db)
					if err != nil {
						db.Rollback()
						tools.ReturnErr101Code(c, err.Error())
						return
					}
					db.Commit()
					tools.ReturnSuccess2000Code(c, "OK")
					return
				}

			}

		} else { //跑分订单

		}

	}

}
