package runner

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/controller/hedge/admin"
	"github.com/wangyi/GinTemplate/dao/mmdb"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"time"
)

// Login 奔跑者登录
func Login(c *gin.Context) {
	var lo LoginVerify
	if err := c.ShouldBind(&lo); err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}
	//校验验证码
	//if VerifyCaptcha(lo.Id, lo.Value) == false {
	//	tools.ReturnErr101Code(c, "The verification code is expired or incorrect")
	//	return
	//}

	runner := model.Runner{}
	if err := mysql.DB.Where("username=? and  password=?", lo.Username, tools.MD5(lo.Password)).First(&runner).Error; err != nil {
		tools.ReturnErr101Code(c, "The account or password is incorrect")
		return
	}
	if runner.Status != 1 {
		tools.ReturnErr101Code(c, "Your account has been blocked, please contact the administrator to cancel it")
		return
	}

	runner.PayPassword = "******"

	//写登录日志
	data := model.AgencyLogData{Content: "登录成功", Ip: c.ClientIP(), Username: lo.Username, Kinds: 1, AgencyRunnerId: runner.AgencyRunnerId}
	data.Add(mysql.DB)
	ip, _ := mmdb.GetCountryForIp(c.ClientIP())
	ups := model.Runner{LastLoginTime: time.Now().Unix(), LastLoginRegion: ip, LastLoginIp: c.ClientIP()}
	mysql.DB.Model(&model.Runner{}).Where("id=?", runner.ID).Update(&ups)

	redis.Rdb.Set("RunnerToken_"+runner.Token, lo.Username, 7*86400*time.Second)
	tools.ReturnSuccess2000DataCode(c, runner, "OK")
	return

}

// Register 注册
func Register(c *gin.Context) {
	var lo RegisterVerify
	//检查参数
	if err := c.ShouldBind(&lo); err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}

	//校验验证码
	if VerifyCaptcha(lo.Id, lo.Value) == false {
		tools.ReturnErr101Code(c, "The verification code is expired or incorrect")
		return
	}

	//判断邀请码长度
	runner := model.Runner{InvitationCode: lo.InvitationCode, Username: lo.Username, Password: tools.MD5(lo.Password)}
	_, err := runner.CheckInvitationCode(mysql.DB)
	if err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}
	//添加数据
	err = runner.Add(mysql.DB, redis.Rdb)
	if err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}
	tools.ReturnSuccess2000Code(c, "OK")
	return
}

// GetSlideshow 获取轮播图
func GetSlideshow(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)

	sl := make([]model.Slideshow, 0)
	mysql.DB.Where("agency_runner_id=?", whoMap.AgencyRunnerId).Find(&sl)
	tools.ReturnSuccess2000DataCode(c, sl, "OK")
	return
}

// LogOut 登出
func LogOut(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	redis.Rdb.Del("RunnerToken_" + whoMap.Token)
	tools.ReturnSuccess2000Code(c, "ok")
	return
}

// GetMe 获取自己信息
func GetMe(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	whoMap.PayPassword = "******"
	//判断自己是否存在下级
	err := mysql.DB.Where("superior=?", whoMap.ID).First(&model.Runner{}).Error
	if err != nil {
		whoMap.IfExistSubordinate = 2
	}
	tools.ReturnSuccess2000DataCode(c, whoMap, "OK")
	return
}

// GetChangeAmount 获取自己账变
func GetChangeAmount(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	//查询bankCard
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	sl := make([]model.RunnerAmountChange, 0)
	db := mysql.DB.Where("runner_id=?", whoMap.ID)
	var total int
	//   条件  1押金 2代收额度  --3代付额度  4佣金   5提现  6账户余额
	if kinds, IsE := c.GetPostForm("kinds"); IsE == true {
		db = db.Where("kinds=?", kinds)
	}
	db.Model(&model.RunnerAmountChange{}).Count(&total)
	db = db.Model(&model.RunnerAmountChange{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
	db.Find(&sl)
	for i, change := range sl {
		mysql.DB.Where("id=?", change.CollectionId).First(&sl[i].Col)
	}

	admin.ReturnDataLIst2000(c, sl, total)
	return

}

// GetCollectionDetail 代收详情
func GetCollectionDetail(c *gin.Context) {
	action := c.Query("action")
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	if action == "today" {
		statistics := model.RunnerStatistics{}
		mysql.DB.Where("runner_id=?  and date=?", whoMap.ID, time.Now().Format("2006-01-02")).First(&statistics)
		tools.ReturnSuccess2000DataCode(c, statistics, "OK")
		return
	}
	if action == "all" {
		type RunnerStatistics struct {
			CollectionCount     int     `json:"collection_count"`
			CollectionAllCount  int     `json:"collection_all_count"`
			CollectionAmount    float64 `json:"collection_amount"`
			CollectionAllAmount float64 `json:"collection_all_amount"`
			Commission          float64 `json:"commission"`
			PayAmount           float64 `json:"pay_amount"`
			PayAllAmount        float64 `json:"pay_all_amount"`
			PayCount            int     `json:"pay_count"`
			PayAllCount         int     `json:"pay_all_count"`
		}
		//SELECT  sum(commission)  FROM  runner_statistics  where  runner_id=2
		var data RunnerStatistics
		mysql.DB.Raw("SELECT  sum(collection_count) AS collection_count  FROM  runner_statistics  where  runner_id=?", whoMap.ID).Scan(&data)
		mysql.DB.Raw("SELECT  sum(collection_all_count) AS collection_all_count  FROM  runner_statistics  where  runner_id=?", whoMap.ID).Scan(&data)
		mysql.DB.Raw("SELECT  sum(collection_amount) AS collection_amount  FROM  runner_statistics  where  runner_id=?", whoMap.ID).Scan(&data)
		mysql.DB.Raw("SELECT  sum(collection_all_amount) AS collection_all_amount  FROM  runner_statistics  where  runner_id=?", whoMap.ID).Scan(&data)
		mysql.DB.Raw("SELECT  sum(commission) AS commission  FROM  runner_statistics  where  runner_id=?", whoMap.ID).Scan(&data)
		mysql.DB.Raw("SELECT  sum(pay_amount) AS pay_amount  FROM  runner_statistics  where  runner_id=?", whoMap.ID).Scan(&data)
		mysql.DB.Raw("SELECT  sum(pay_all_amount) AS pay_all_amount  FROM  runner_statistics  where  runner_id=?", whoMap.ID).Scan(&data)
		mysql.DB.Raw("SELECT  sum(pay_count) AS pay_count  FROM  runner_statistics  where  runner_id=?", whoMap.ID).Scan(&data)
		mysql.DB.Raw("SELECT  sum(pay_all_count) AS pay_all_count  FROM  runner_statistics  where  runner_id=?", whoMap.ID).Scan(&data)

		tools.ReturnSuccess2000DataCode(c, data, "OK")
		return
	}

}

// GetCollectionRecord 获取自己的接单记录
func GetCollectionRecord(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	sl := make([]modelPay.Collection, 0)
	db := mysql.DB.Where("runner_id=?", whoMap.ID)
	var total int

	//   条件  1押金 2代收额度  --3代付额度  4佣金   5提现  6账户余额
	if kinds, IsE := c.GetPostForm("kinds"); IsE == true {
		db = db.Where("kinds=?", kinds)
	}
	//status
	if kinds, IsE := c.GetPostForm("status"); IsE == true {
		db = db.Where("status=?", kinds)
	}

	//订单号
	if kinds, IsE := c.GetPostForm("own_order"); IsE == true {
		db = db.Where("own_order=?", kinds)
	}

	//时间段搜索
	if start, IsE := c.GetPostForm("start"); IsE == true {
		if end, IsE := c.GetPostForm("end"); IsE == true {
			db = db.Where("created  <=  ? and created >=  ?", end, start)
		}
	}

	db.Model(&modelPay.Collection{}).Count(&total)
	db = db.Model(&modelPay.Collection{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
	db.Find(&sl)

	for i, _ := range sl {
		sl[i].NoticeUrl = ""
		sl[i].CallbackContent = ""
	}

	admin.ReturnDataLIst2000(c, sl, total)
	return

}

// Withdraw 提现
func Withdraw(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Runner)
	//判断提现卡的类型
	kind, _ := strconv.Atoi(c.PostForm("kinds"))
	amount, _ := strconv.ParseFloat(c.PostForm("amount"), 64)
	if kind != 2 && kind != 3 {
		tools.ReturnErr101Code(c, "The withdrawal type is incorrect")
		return
	}
	if amount < 0 {
		tools.ReturnErr101Code(c, "your withdrawal amount mush be greater then zero ")
		return
	}
	if amount > whoMap.Balance {
		tools.ReturnErr101Code(c, "sorry,your balance is not enough")
		return
	}
	//判断是是否绑定
	UPI := model.RunnerUpi{}
	err := mysql.DB.Where("kind= ? and runner_id=?", kind, whoMap.ID).First(&UPI).Error
	if err != nil {
		tools.ReturnErr101Code(c, "Please bind your withdrawal bank card first")
		return
	}

	//判断最小提现金额
	ag := model.AgencyRunner{}
	mysql.DB.Where("id=?", whoMap.AgencyRunnerId).First(&ag)
	if amount < ag.MinWithdraw {
		tools.ReturnErr101Code(c, fmt.Sprintf("Minimum withdrawal amount  is  %2.f", ag.MinWithdraw))
		return

	}
	//更新最新余额
	ups := map[string]interface{}{}
	db := mysql.DB.Begin()
	ups["Balance"] = whoMap.Balance - amount
	ups["FreezeMoney"] = whoMap.FreezeMoney + amount
	affected := db.Model(&model.Runner{}).Where("id=? and  balance=? and freeze_money=?", whoMap.ID, whoMap.Balance, whoMap.FreezeMoney).Update(ups).RowsAffected
	if affected == 0 {
		tools.ReturnErr101Code(c, "Sorry, system error, please withdraw again")
		return
	}
	//生成账单
	record := model.Record{
		RunnerId:       whoMap.ID,
		AgencyRunnerId: whoMap.AgencyRunnerId,
		Kinds:          1,
		Amount:         amount, WithdrawalMethod: kind}
	err = record.Add(db)
	if err != nil {
		db.Rollback()
		tools.ReturnErr101Code(c, "Sorry, system error, please withdraw again")
		return
	}
	//生成账变
	fmt.Println(record)
	//账变
	change := model.RunnerAmountChange{RunnerId: whoMap.ID, NowAmount: whoMap.Balance - amount, FontAmount: whoMap.Balance, ChangeAmount: amount, Kinds: 6, RecordId: record.ID}
	err = change.Add(db)
	if err != nil {
		db.Rollback()
		tools.ReturnErr101Code(c, "Sorry, system error, please withdraw again")
		return
	}
	db.Commit()
	tools.ReturnSuccess2000Code(c, "Successful  withdraw,waiting  for  management review")
	return
}

// RankingList 排行榜
func RankingList(c *gin.Context) {
	action := c.Query("action")
	type Data struct {
		Username   string  `json:"username"`
		Commission float64 `json:"commission"`
		Id         int     `json:"id"`
	}
	if action == "today" {
		var data []Data
		mysql.DB.Raw("SELECT SUM(change_amount),runner_id as id  FROM runner_amount_changes   WHERE  kinds=4   GROUP BY  runner_id   ORDER BY sum(change_amount)  LIMIT  10  ").Scan(&data)

		for i, datum := range data {
			runner := model.Runner{}
			mysql.DB.Where("id=?", datum.Id).First(&runner)
			data[i].Id = 0
			data[i].Username = runner.Username[0:1] + "*****" + runner.Username[len(runner.Username)-2:]
		}

		tools.ReturnSuccess2000DataCode(c, data, "OK")
		return
	}

	if action == "accumulative" {
		data := make([]Data, 0)
		mysql.DB.Raw("select username as username, commission as commission from runners where status=1 order by commission  desc ").Scan(&data)
		for i, datum := range data {
			data[i].Username = datum.Username[0:1] + "*****" + datum.Username[len(datum.Username)-2:]
		}
		tools.ReturnSuccess2000DataCode(c, data, "OK")
		return
	}

}
