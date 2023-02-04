package runner

import (
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
	if kind != 2 && kind != 3 {
		tools.ReturnErr101Code(c, "The withdrawal type is incorrect")
		return
	}
	//判断是是否绑定
	UPI := model.RunnerUpi{}
	err := mysql.DB.Where("kinds= ? and runner_id=?", kind, whoMap.ID).First(&UPI).Error
	if err != nil {
		tools.ReturnErr101Code(c, "Please bind your withdrawal bank card first")
		return
	}

	//

}
