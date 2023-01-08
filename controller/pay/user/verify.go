package user

// UploadCertificateVerify LoginVerify 登录参数
type UploadCertificateVerify struct {
	ProofOfPayment string `form:"proof_of_payment"  binding:"required,max=12,min=12"`
	OrderNum       string `form:"order_num"  binding:"required,max=25"`
}
