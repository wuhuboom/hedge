package merchant

// LoginVerify 登录参数
type LoginVerify struct {
	MerchantNum   string `form:"merchant_num"  binding:"required"`
	LoginPassword string `form:"login_password"  binding:"required"`
	GoogleCode    string `form:"googleCode"  binding:"omitempty,max=6" `
	GoogleSecret  string `form:"googleSecret"  binding:"omitempty" `
}
