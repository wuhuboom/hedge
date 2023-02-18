package agency

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/controller/hedge/admin"
	"github.com/wangyi/GinTemplate/dao/mmdb"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"os"
	"strconv"
	"strings"
	"time"
)

// Login 登录
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
				err := mysql.DB.Model(&model.AgencyRunner{}).Where("id=?", admin.ID).Update(model.AgencyRunner{GoogleCode: lo.GoogleSecret}).Error
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

// GetAmountChange 资金流水
func GetAmountChange(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.AgencyRunner)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.AgencyAccountChange, 0)
		db := mysql.DB.Where("agency_runner_id=?", whoMap.ID)
		var total int
		//   条件  1押金 2代收额度  3代付额度  4佣金   5提现
		if kinds, IsE := c.GetPostForm("kinds"); IsE == true {
			db = db.Where("kinds=?", kinds)
		}
		//用户名  runner_id
		if username, isE := c.GetPostForm("username"); isE == true {
			//查找这个  runner_id
			runner := model.Runner{Username: username}
			db = db.Where("runner_id=?", runner.GetRunnerId(mysql.DB))
		}
		////runner_collection_order_id
		//if runnerCollectionOrder, isE := c.GetPostForm("runner_collection_order"); isE == true {
		//	order := model.RunnerCollectionOrder{OrderNum: runnerCollectionOrder}
		//	db = db.Where("runner_collection_order_id=?", order.GetId(mysql.DB))
		//}
		////runner_pay_order_id
		//if runnerCollectionOrder, isE := c.GetPostForm("runner_pay_order"); isE == true {
		//	order := model.RunnerPayOrder{OrderNum: runnerCollectionOrder}
		//	db = db.Where("runner_pay_order_id=?", order.GetId(mysql.DB))
		//}

		//record_id
		if runnerCollectionOrder, isE := c.GetPostForm("record_order"); isE == true {
			order := model.Record{OrderNum: runnerCollectionOrder}
			db = db.Where("record_id=?", order.GetId(mysql.DB))
		}

		db.Model(model.AgencyAccountChange{}).Count(&total)
		db = db.Model(&model.AgencyAccountChange{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		admin.ReturnDataLIst2000(c, sl, total)
		return
	}

}

// SetMyselfConfig 修改  设置
func SetMyselfConfig(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.AgencyRunner)
	ups := map[string]interface{}{}

	//客服地址
	if ca, isE := c.GetPostForm("customer_service_address"); isE == true {
		ups["CustomerServiceAddress"] = ca
	}
	//下级提现手续费(比率)
	if ca, isE := c.GetPostForm("junior_withdraw_commission"); isE == true {
		ups["JuniorWithdrawCommission"], _ = strconv.ParseFloat(ca, 64)
	}
	//JuniorPoint  下级税点
	if ca, isE := c.GetPostForm("junior_point"); isE == true {
		parseFloat, _ := strconv.ParseFloat(ca, 64)
		ups["JuniorPoint"]= parseFloat
		if parseFloat >whoMap.PayPoint   || parseFloat >whoMap.CollectionPoint {
			tools.ReturnErr101Code(c,"cant  > Cannot be greater than its own profit point ")
			return
		}


	}

	//min_withdraw 代理最小提现金额
	if ca, isE := c.GetPostForm("min_withdraw"); isE == true {
		ups["MinWithdraw"], _ = strconv.ParseFloat(ca, 64)
	}

	err := mysql.DB.Model(&model.AgencyRunner{}).Where("id=?", whoMap.ID).Update(ups).Error
	if err != nil {
		tools.ReturnErr101Code(c, "sorry, "+err.Error())
		return
	}
	tools.ReturnSuccess2000Code(c, "OK")
	return

}

// SlideshowOperation 获取轮播图
func SlideshowOperation(c *gin.Context) {
	who, _ := c.Get("who")
	whoMap := who.(model.AgencyRunner)
	action := c.Query("action")
	if action == "check" {
		//查询bankCard
		limit, _ := strconv.Atoi(c.PostForm("limit"))
		page, _ := strconv.Atoi(c.PostForm("page"))
		sl := make([]model.Slideshow, 0)
		db := mysql.DB.Where("agency_runner_id=?", whoMap.ID)
		var total int
		//  状态
		if kinds, IsE := c.GetPostForm("status"); IsE == true {
			db = db.Where("status=?", kinds)
		}
		db.Model(model.Slideshow{}).Count(&total)
		db = db.Model(&model.Slideshow{}).Offset((page - 1) * limit).Limit(limit).Order("created desc")
		db.Find(&sl)
		admin.ReturnDataLIst2000(c, sl, total)
		return
	}
	if action == "add" {
		file, _ := c.FormFile("file")

		path := "./static/upload/" + time.Now().Format("20060102") + "/"
		if noExist, _ := tools.IsFileNotExist(path); noExist {
			if err := os.MkdirAll(path, 0777); err != nil {
				tools.ReturnErr101Code(c, err.Error())
				return
			}
		}
		nameArray := strings.Split(file.Filename, ".")
		path = path + time.Now().Format("20060102150405") + nameArray[0] + "." + nameArray[1]
		err := c.SaveUploadedFile(file, path)
		if err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
		slideshow := model.Slideshow{Url: path, AgencyRunnerId: whoMap.ID}
		slideshow.Add(mysql.DB)
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}

	if action == "del" {
		id := c.PostForm("id")
		atom, _ := strconv.Atoi(id)
		mysql.DB.Model(&model.Slideshow{}).Delete(&model.Slideshow{ID: atom})
		tools.ReturnSuccess2000Code(c, "OK")
		return
	}

}
