package agency

type LoginVerify struct {
	Username     string `form:"username"  binding:"required"`
	Password     string `form:"password"  binding:"required"`
	GoogleCode   string `form:"googleCode"  binding:"omitempty,max=6" `
	GoogleSecret string `form:"googleSecret"  binding:"omitempty" `
}

type RunnerOperationAdd struct {
	Username        string `form:"username"  binding:"required"`
	Password        string `form:"password"  binding:"required"`
	CollectionPoint string `form:"collection_point" binding:"required"`
	PayPoint        string `form:"pay_point" binding:"required"`
	JuniorPoint     string `form:"junior_point"  binding:"required"`
}
