package user

import (
	"mime/multipart"
)

// UploadCertificateVerify LoginVerify 登录参数
type UploadCertificateVerify struct {
	ProofOfPayment string `form:"proof_of_payment"  binding:"required,max=12,min=12"`
	OrderNum       string `form:"order_num"  binding:"required,max=25"`
}

type UploadCertificateVerifyForImage struct {
	ProofOfPayment *multipart.FileHeader `form:"file" validate:"required,image" `
	OrderNum       string                `form:"order_num"  binding:"required,max=25"`
}
