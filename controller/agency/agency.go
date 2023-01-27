package agency

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mmdb"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strings"
	"time"
)

func Login(c *gin.Context) {
	var lo LoginVerify
	//检查参数
	if err := c.ShouldBind(&lo); err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}
	//判断系统是否开启的谷歌验证
	lo.Password = tools.MD5(lo.Password)
	admin := model.AgencyRunner{}
	err11 := mysql.DB.Where("username=? and password= ?", lo.Username, lo.Password).First(&admin).Error
	if err11 != nil {
		tools.ReturnErr101Code(c, "Sorry, wrong account number or password")
		return
	}

	if admin.Status != 1 {
		tools.ReturnErr101Code(c, "Sorry, the account has been banned illegally, please contact the management for unblocking")
		return
	}

	//谷歌开启
	if admin.GoogleSwitch == 1 {
		//判断这个用户是否已经绑定了谷歌
		if admin.GoogleCode == "" {
			//没有绑定谷歌  所以要返回谷歌的验证码
			if lo.GoogleSecret == "" {
				secret, _, qrCodeUrl := tools.InitAuth(lo.Username)
				tools.JsonWrite(c, tools.NeedGoogleBind, map[string]string{"codeUrl": qrCodeUrl, "googleSecret": secret}, "Please bind your Google account first")
				return

			} else {
				verifyCode, _ := tools.NewGoogleAuth().VerifyCode(lo.GoogleSecret, lo.GoogleCode)
				if !verifyCode {
					tools.ReturnErr101Code(c, "Google verification failure")
					return
				}
				err := mysql.DB.Model(&model.Admin{}).Where("id=?", admin.ID).Update(model.Admin{GoogleCode: lo.GoogleSecret}).Error
				if err != nil {
					tools.ReturnErr101Code(c, err.Error())
					return
				}
			}

		} else {
			//校验谷歌验证
			verifyCode, _ := tools.NewGoogleAuth().VerifyCode(admin.GoogleCode, lo.GoogleCode)
			if !verifyCode {
				tools.ReturnErr101Code(c, "Google verification failure")
				return
			}
		}
	}
	admin.PayPassword = "******"

	//修改登陆成功
	are, _ := mmdb.GetCountryForIp(c.ClientIP())
	ups := model.AgencyRunner{LastLoginIp: c.ClientIP(), LastLoginRegion: are, LastLoginTime: time.Now().Unix()}
	mysql.DB.Model(&model.AgencyRunner{}).Where("id=?", admin.ID).Update(&ups)
	//登录  日志
	data := model.AgencyLogData{Ip: c.ClientIP(), Username: admin.Username, Content: "登陆成功", Kinds: 1, AgencyRunnerId: admin.ID}
	data.Add(mysql.DB)
	redis.Rdb.Set("AgencyToken_"+admin.Token, admin.Username, 7*86400*time.Second)
	tools.ReturnSuccess2000DataCode(c, admin, "Login success")
	return
}

// GetMe 获取自己的基本信息
func GetMe(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.AgencyRunner)
	ccArray := strings.Split(whoMap.CollectionChannel, "@")
	for _, s := range ccArray {
		ch := modelPay.Channel{}
		err := mysql.DB.Where("id=?", s).First(&ch).Error
		if err == nil {
			whoMap.CollectionChannelArray = append(whoMap.CollectionChannelArray, ch)

		}

	}

	pcArray := strings.Split(whoMap.PayChannel, "@")
	for _, s := range pcArray {
		ch := modelPay.Channel{}
		err := mysql.DB.Where("id=?", s).First(&ch).Error
		if err == nil {
			whoMap.PayChannelArray = append(whoMap.PayChannelArray, ch)

		}

	}
	whoMap.PayPassword = "******"
	tools.ReturnSuccess2000DataCode(c, whoMap, "OK")
	return

}
