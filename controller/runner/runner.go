package runner

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mmdb"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
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
