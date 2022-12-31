package merchant

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"time"
)

// 登录

func Login(c *gin.Context) {
	var lo LoginVerify
	//检查参数
	if err := c.ShouldBind(&lo); err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}
	//判断系统是否开启的谷歌验证
	config := model.Config{}
	err := mysql.DB.Where("id=?", 1).First(&config).Error
	if err != nil {
		tools.ReturnErr101Code(c, "System error. Please contact technical")
		return
	}
	lo.LoginPassword = tools.MD5(lo.LoginPassword)
	mer := model.Merchant{}
	err11 := mysql.DB.Where("merchant_num=? and login_password= ?", lo.MerchantNum, lo.LoginPassword).First(&mer).Error
	//谷歌开启
	if config.IfUseGoogleCode == 1 {
		//判断这个用户是否已经绑定了谷歌
		if mer.GoogleCode == "" {
			//没有绑定谷歌  所以要返回谷歌的验证码
			if lo.GoogleSecret == "" {
				secret, _, qrCodeUrl := tools.InitAuth(lo.MerchantNum)
				tools.JsonWrite(c, tools.NeedGoogleBind, map[string]string{"codeUrl": qrCodeUrl, "googleSecret": secret}, "Please bind your Google account first")
				return

			} else {
				verifyCode, _ := tools.NewGoogleAuth().VerifyCode(lo.GoogleSecret, lo.GoogleCode)
				if !verifyCode {
					tools.ReturnErr101Code(c, "Google verification failure")
					return
				}
				err := mysql.DB.Model(&model.Merchant{}).Where("id=?", mer.ID).Update(&model.Merchant{GoogleCode: lo.GoogleSecret}).Error
				if err != nil {
					tools.ReturnErr101Code(c, err.Error())
					return
				}
			}

		} else {
			//校验谷歌验证
			verifyCode, _ := tools.NewGoogleAuth().VerifyCode(mer.GoogleCode, lo.GoogleCode)
			if !verifyCode {
				tools.ReturnErr101Code(c, "Google verification failure")
				return
			}
		}
	}

	if err11 != nil {
		tools.ReturnErr101Code(c, "Sorry, wrong account number or password")
		return
	}
	redis.Rdb.Set("MerchantToken_"+mer.Token, mer.MerchantNum, 7*86400*time.Second)
	tools.ReturnSuccess2000DataCode(c, mer, "Login success")
	return
}
