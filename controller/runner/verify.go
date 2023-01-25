package runner

// RegisterVerify 登录参数
type RegisterVerify struct {
	Username       string `form:"username"  binding:"required"`
	Password       string `form:"password"  binding:"required"`
	InvitationCode string `form:"invitation_code"  binding:"required"`
}
