/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wangyi/GinTemplate/controller/agency"
	admin2 "github.com/wangyi/GinTemplate/controller/hedge/admin"
	"github.com/wangyi/GinTemplate/controller/hedge/h5"
	"github.com/wangyi/GinTemplate/controller/hedge/three"
	"github.com/wangyi/GinTemplate/controller/pay/management"
	"github.com/wangyi/GinTemplate/controller/pay/merchant"
	"github.com/wangyi/GinTemplate/controller/pay/tripartiteTerminal"
	"github.com/wangyi/GinTemplate/controller/pay/user"
	"github.com/wangyi/GinTemplate/controller/runner"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/dao/redis"
	eeor "github.com/wangyi/GinTemplate/error"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"log"
	"net/http"
	"time"
)

func Setup() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Use(Cors())
	r.Use(eeor.ErrHandler())
	r.NoMethod(eeor.HandleNotFound)
	r.NoRoute(eeor.HandleNotFound)
	r.Static("/static", "./static")
	//r.Use(PermissionToCheck())

	ad := r.Group("/admin/v1", PermissionToCheckForAdmin())
	{
		//管理登录
		ad.POST("/login", admin2.Login)
		//商户号管理
		ad.POST("/merchant", admin2.MerchantOperation)
		//PaidOperation  代付订单处理
		ad.POST("/paidOperation", admin2.PaidOperation)
		//ConfigOperation
		ad.POST("/configOperation", admin2.ConfigOperation)
		//PayOperation
		ad.POST("/payOperation", admin2.PayOperation)

		//代理处理  AgencyOperation
		ad.POST("/agencyOperation", admin2.AgencyOperation)
		//CollectionOperation
		ad.POST("/agencyCollectionOperation", admin2.CollectionOperation)
		//总订单  OrderOperation
		ad.POST("/orderOperation", admin2.OrderOperation)

		/***
		  印度支付
		*/
		ad.POST("/bankInformation", management.BankInformation)
		//Bank
		ad.POST("/bank", management.Bank)
		//Channel
		ad.POST("/channel", management.Channel)
		//ChannelBank
		ad.POST("/channelBank", management.ChannelBank)
		//CollectionOperation    //  获取代收 或者代付 订单
		ad.POST("/collectionOperation", management.CollectionOperation)
		//GatewayOperation
		ad.POST("/gatewayOperation", management.GatewayOperation)
		//GetStatistics
		ad.POST("/getStatistics", management.GetStatistics)
	}

	paid := r.Group("/paid/v1")
	{
		paid.POST("paid", three.Paid)
		//OrderInquiry
		paid.POST("/orderInquiry", three.OrderInquiry)
		//NoticeUseOtherPaid
		paid.POST("/noticeUseOtherPaid", three.NoticeUseOtherPaid)

	}
	pay := r.Group("/modelPay/v1")
	{
		pay.POST("modelPay", h5.Pay)
		//上传 支付凭证
		pay.POST("uploadPayCertificate", LimitIpRequestSameUrlForUser(), h5.UploadPayCertificate)
		//CheckImageForUpload  查看凭证
		pay.POST("checkImageForUpload", LimitIpRequestSameUrlForUser(), h5.CheckImageForUpload)

	}
	//印度支付 拉起订单
	tripartite := r.Group("/tripartiteTerminal/v2")
	{
		//   来气代收
		tripartite.POST("/createCollection", tripartiteTerminal.CollectionAmount)

		//  拉起 代付 PayAmount
		tripartite.POST("/payAmount", tripartiteTerminal.PayAmount)

	}
	//CollectionAmount
	//Login
	mer := r.Group("/merchant/v2").Use(PermissionToCheckForMerchant())
	{
		//登录
		mer.POST("/login", merchant.Login)
		//代收订单  collection
		mer.POST("/collection", merchant.CollectionAmount)
		//GetMe
		mer.POST("/getMe", merchant.GetMe)
		//GetFlowOfFunds
		mer.POST("/getFlowOfFunds", merchant.GetFlowOfFunds)
		//Statistics
		mer.POST("/statistics", merchant.Statistics)
		//ChangeLoginPassword 修改密码
		mer.POST("/changeLoginPassword", merchant.ChangeLoginPassword)
		//GetLoginLogger
		mer.POST("/getLoginLogger", merchant.GetLoginLogger)
		//ChangeTrcAddress
		mer.POST("/changeTrcAddress", merchant.ChangeTrcAddress)

	}
	//UploadCertificate
	us := r.Group("/user/v2", LimitIpRequestSameUrlForUser())
	{
		us.POST("/uploadCertificate", user.UploadCertificate)
	}
	//Runner
	run := r.Group("/runner/v2", PermissionToCheckForRunner(), LimitIpRequestSameUrlForUser())
	{
		run.POST("/register", runner.Register)
		run.POST("/login", runner.Login)
		run.POST("/generateCaptcha", runner.GenerateCaptcha)
		run.POST("/getAnnouncement", runner.GetAnnouncement)
		run.POST("/getCustomerServiceAddress", runner.GetCustomerServiceAddress)
		run.POST("/getSlideshow", runner.GetSlideshow)
		run.POST("/getReceiveCollectionOrder", runner.GetReceiveCollectionOrder)
		run.POST("/setUpi", runner.SetUpi)
		run.POST("/imWorking", runner.ImWorking)
		run.POST("/confirmTheOrder", runner.ConfirmTheOrder)
		run.POST("/logOut", runner.LogOut)
		run.POST("/getMe", runner.GetMe)
	}

	ay := r.Group("/agency/v2", PermissionToCheckForAgency())
	{
		ay.POST("/login", agency.Login)
		ay.POST("/getMe", agency.GetMe)
		ay.POST("/logger", agency.Logger)
		ay.POST("/getAmountChange", agency.GetAmountChange) //获取账户变化
		ay.POST("/runnerOperation", agency.RunnerOperation)
		ay.POST("/announcementOperation", agency.AnnouncementOperation)
		//SetMyselfConfig
		ay.POST("/setMyselfConfig", agency.SetMyselfConfig)
		//SlideshowOperation
		ay.POST("/slideshowOperation", agency.SlideshowOperation)
		//CollectionOperation
		ay.POST("/collectionOperation", agency.CollectionOperation)

	}

	r.Run(fmt.Sprintf(":%d", viper.GetInt("app.port")))
	return r
}

// Cors /**  跨域

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}

		defer func() {
			if err := recover(); err != nil {
				log.Printf("Panic info is: %v", err)
			}
		}()

		c.Next()
	}
}

func LimitIpRequestSameUrlForUser() gin.HandlerFunc {
	return func(context *gin.Context) {
		//获取请求的地址
		urlPath := context.Request.URL.Path
		ip := context.ClientIP()
		//回去系统设置的ip限制次数
		var LimitTimes int
		LimitTimes = 3
		key := ip + urlPath
		curr := redis.Rdb.LLen(key).Val()
		if int(curr) >= LimitTimes {
			//超出了限制
			tools.JsonWrite(context, tools.IpLimitWaring, nil, "Don't ask too fast")
			context.Abort()
		}
		if v := redis.Rdb.Exists(key).Val(); v == 0 {
			pipe := redis.Rdb.TxPipeline()
			pipe.RPush(key, key)
			//设置过期时间
			pipe.Expire(key, 1*time.Second)
			_, _ = pipe.Exec()
		} else {
			redis.Rdb.RPushX(key, key)
		}

		context.Next()

	}
}

// PermissionToCheckForAdmin 管理控制   token  32
func PermissionToCheckForAdmin() gin.HandlerFunc {
	whiteUrl := []string{
		"/admin/v1/login",
	}
	return func(c *gin.Context) {
		if !tools.IsArray(whiteUrl, c.Request.RequestURI) {
			token := c.Request.Header.Get("token")
			if len(token) != 32 {
				tools.JsonWrite(c, tools.IllegalityCode, nil, "An unlawful request")
				c.Abort()
				return
			}
			ad := model.Admin{}
			err := mysql.DB.Where("token=?", token).First(&ad).Error
			if err != nil {
				tools.JsonWrite(c, tools.IllegalityCode, nil, "An unlawful request")
				c.Abort()
				return
			}
			//判断token 是否过期?
			if redis.Rdb.Get("AdminToken_"+token).Val() == "" {
				tools.JsonWrite(c, tools.TokenExpire, nil, "An unlawful request")
				c.Abort()
				return
			}
			//设置who
			c.Set("who", ad)
			c.Next()
		}
	}
}

// PermissionToCheckForMerchant 商户 控制  token 36
func PermissionToCheckForMerchant() gin.HandlerFunc {
	whiteUrl := []string{
		"/merchant/v2/login"}
	return func(c *gin.Context) {
		if !tools.IsArray(whiteUrl, c.Request.RequestURI) {
			token := c.Request.Header.Get("token")
			//商户号
			if len(token) != 36 {
				tools.JsonWrite(c, tools.IllegalityCode, nil, "An unlawful request")
				c.Abort()
				return
			}

			ad := model.Merchant{}
			err := mysql.DB.Where("token=?", token).First(&ad).Error
			if err != nil {
				tools.JsonWrite(c, tools.IllegalityCode, nil, "An unlawful request")
				c.Abort()
				return
			}
			//判断token 是否过期?
			if redis.Rdb.Get("MerchantToken_"+token).Val() == "" {
				tools.JsonWrite(c, tools.TokenExpire, nil, "An unlawful request")
				c.Abort()
				return
			}

			if ad.Status != 1 {
				tools.JsonWrite(c, -101, nil, "The account has been disabled")
				c.Abort()
				return
			}

			c.Set("who", ad)
			c.Next()
		}
	}
}

func PermissionToCheckForAgency() gin.HandlerFunc {
	whiteUrl := []string{
		"/agency/v2/login"}
	return func(c *gin.Context) {
		if !tools.IsArray(whiteUrl, c.Request.RequestURI) {
			token := c.Request.Header.Get("token")
			//商户号
			if len(token) != 31 {
				tools.JsonWrite(c, tools.IllegalityCode, nil, "An unlawful request")
				c.Abort()
				return
			}

			ad := model.AgencyRunner{}
			err := mysql.DB.Where("token=?", token).First(&ad).Error
			if err != nil {
				tools.JsonWrite(c, tools.IllegalityCode, nil, "An unlawful request")
				c.Abort()
				return
			}
			//判断token 是否过期?
			if redis.Rdb.Get("AgencyToken_"+token).Val() == "" {
				tools.JsonWrite(c, tools.TokenExpire, nil, "An unlawful request")
				c.Abort()
				return
			}
			c.Set("who", ad)
			c.Next()
		}
	}
}

func PermissionToCheckForRunner() gin.HandlerFunc {
	whiteUrl := []string{
		"/runner/v2/login", "/runner/v2/register", "/runner/v2/generateCaptcha"}
	return func(c *gin.Context) {
		if !tools.IsArray(whiteUrl, c.Request.RequestURI) {
			token := c.Request.Header.Get("token")
			//商户号
			if len(token) != 48 {
				tools.JsonWrite(c, tools.IllegalityCode, nil, "An unlawful request")
				c.Abort()
				return
			}

			ad := model.Runner{}
			err := mysql.DB.Where("token=?", token).First(&ad).Error
			if err != nil {
				tools.JsonWrite(c, tools.IllegalityCode, nil, "An unlawful request")
				c.Abort()
				return
			}
			//判断token 是否过期?
			if redis.Rdb.Get("RunnerToken_"+token).Val() == "" {
				tools.JsonWrite(c, tools.TokenExpire, nil, "An unlawful request")
				c.Abort()
				return
			}
			c.Set("who", ad)
			c.Next()
		}
	}
}
