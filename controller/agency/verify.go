package agency

type LoginVerify struct {
	Username     string `form:"username"  binding:"required"`
	Password     string `form:"password"  binding:"required"`
	GoogleCode   string `form:"googleCode"  binding:"omitempty,max=6" `
	GoogleSecret string `form:"googleSecret"  binding:"omitempty" `
}
