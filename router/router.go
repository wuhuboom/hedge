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
	admin2 "github.com/wangyi/GinTemplate/controller/hedge/admin"
	"github.com/wangyi/GinTemplate/controller/hedge/h5"
	"github.com/wangyi/GinTemplate/controller/hedge/three"
	"github.com/wangyi/GinTemplate/controller/pay/management"
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
	r.Use(PermissionToCheck())

	ad := r.Group("/admin/v1")
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

// PermissionToCheck 权限校验
func PermissionToCheck() gin.HandlerFunc {
	whiteUrl := []string{
		"/admin/v1/login",
		"/paid/v1/paid",
		"/modelPay/v1/modelPay",
		"/paid/v1/orderInquiry",
		"/paid/v1/noticeUseOtherPaid", "/modelPay/v1/uploadPayCertificate", "/modelPay/v1/checkImageForUpload"}

	return func(c *gin.Context) {
		if !tools.IsArray(whiteUrl, c.Request.RequestURI) {
			//token  校验
			//判断是用户还是管理员
			fmt.Println(c.Request.URL.Path)
			token := c.Request.Header.Get("token")
			//用户
			if len(token) == 36 {
				//ad := model.User{}
				//err := mysql.DB.Where("token=?", token).First(&ad).Error
				//if err != nil {
				//	tools.JsonWrite(c, client.IllegalityCode, nil, client.IllegalityMsg)
				//	c.Abort()
				//}
				////判断token 是否过期?
				//if redis.Rdb.Get("UserToken_"+token).Val() == "" {
				//	tools.JsonWrite(c, client.TokenExpire, nil, client.LoginExpire)
				//	c.Abort()
				//}
				//c.Set("who", ad)
				//c.Next()
			} else if len(token) == 32 {
				//管理员
				ad := model.Admin{}
				err := mysql.DB.Where("token=?", token).First(&ad).Error
				if err != nil {
					tools.JsonWrite(c, tools.IllegalityCode, nil, "An unlawful request")
					c.Abort()
				}
				//判断token 是否过期?
				if redis.Rdb.Get("AdminToken_"+token).Val() == "" {
					tools.JsonWrite(c, tools.TokenExpire, nil, "An unlawful request")
					c.Abort()
				}

				//设置who
				c.Set("who", ad)
				c.Next()
			} else {
				tools.JsonWrite(c, tools.IllegalityCode, nil, "An unlawful request")
				c.Abort()
			}

		}

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
		LimitTimes = 2
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