package merchant

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"strconv"
	"strings"
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

	lo.LoginPassword = tools.MD5(lo.LoginPassword)
	mer := model.Merchant{}
	err11 := mysql.DB.Where("merchant_num=? and login_password= ?", lo.MerchantNum, lo.LoginPassword).First(&mer).Error
	if err11 != nil {
		tools.ReturnErr101Code(c, "Sorry, wrong account number or password")
		return
	}
	//谷歌开启
	if mer.GoogleSwitch == 1 {
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

	redis.Rdb.Set("MerchantToken_"+mer.Token, mer.MerchantNum, 7*86400*time.Second)
	tools.ReturnSuccess2000DataCode(c, mer, "Login success")
	return
}

//获取个人信息

func GetMe(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Merchant)
	// 代收通道
	payArray := strings.Split(whoMap.PayChannel, "@")
	for _, s := range payArray {
		pay := modelPay.Channel{}
		err := mysql.DB.Where("id =?", s).First(&pay).Error
		if err == nil {
			whoMap.PayC = append(whoMap.PayC, pay)
		}

	}
	//代付
	paidArray := strings.Split(whoMap.PaidChannel, "@")
	for _, s := range paidArray {
		pay := modelPay.Channel{}
		err := mysql.DB.Where("id =?", s).First(&pay).Error
		if err == nil {
			whoMap.PayC = append(whoMap.PayC, pay)
		}
	}
	gateway := model.Gateway{ID: whoMap.GatewayId}

	whoMap.Gateway = gateway.Gateway
	fmt.Println(whoMap)
	tools.ReturnSuccess2000DataCode(c, whoMap, "OK")
	return
}

// GetFlowOfFunds 获取账户资金流水
func GetFlowOfFunds(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.Merchant)
	//查询bankCard
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	page, _ := strconv.Atoi(c.PostForm("page"))
	sl := make([]modelPay.AmountChange, 0)
	db := mysql.DB.Where("merchant_num=?", whoMap.MerchantNum)
	var total int
	//条件
	if start, IsE := c.GetPostForm("start"); IsE == true {
		if end, IsE := c.GetPostForm("end"); IsE == true {
			//startInt, _ := strconv.ParseInt(start, 10, 64)
			//endInt, _ := strconv.ParseInt(end, 10, 64)
			db = db.Where("created  <=  ? and created >=  ?", end, start)
		}
	}
	//商家订单号
	//if status, IsE := c.GetPostForm("merchant_order_num"); IsE == true {
	//	db = db.Where("", status)
	//}

	db.Model(&modelPay.AmountChange{}).Count(&total)
	db = db.Model(&modelPay.AmountChange{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
	db.Find(&sl)
	tools.ReturnDataLIst2000(c, sl, total)
	return

}
