package runner

import (
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/tools"
)

// Login 奔跑者登录
func Login(c *gin.Context) {

}

// Register 注册
func Register(c *gin.Context) {
	var lo RegisterVerify
	//检查参数
	if err := c.ShouldBind(&lo); err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}

}
