package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginVerify 登录参数
type LoginVerify struct {
	Username     string `form:"adminUsername"  binding:"required"`
	Password     string `form:"adminPassword"  binding:"required"`
	GoogleCode   string `form:"googleCode"  binding:"omitempty,max=6" `
	GoogleSecret string `form:"googleSecret"  binding:"omitempty" `
}

func ReturnDataLIst2000(context *gin.Context, data interface{}, count int) {
	context.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": data,
		"count":  count,
	})
}
