package runner

// RegisterVerify 登录参数
type RegisterVerify struct {
	Username       string `form:"username"  binding:"required,min=6,max=20"`
	Password       string `form:"password"  binding:"required,min=6,max=20"`
	InvitationCode string `form:"invitation_code"  binding:"required" `
}
