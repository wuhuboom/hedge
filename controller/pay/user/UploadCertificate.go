package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
)

// UploadCertificate 玩家上传支付凭证
func UploadCertificate(c *gin.Context) {

	var lo UploadCertificateVerify
	//检查参数
	if err := c.ShouldBind(&lo); err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}

	fmt.Println(lo)
	//判断这个订单是否存在
	col := modelPay.Collection{}
	err := mysql.DB.Where("own_order=?", lo.OrderNum).First(&col).Error
	if err != nil {
		tools.ReturnErr101Code(c, "Sorry, you cannot submit illegally")
		return
	}
	err = mysql.DB.Model(&modelPay.Collection{}).Where("id=?", col.ID).Update(&modelPay.Collection{ProofOfPayment: lo.ProofOfPayment}).Error
	if err != nil {
		tools.ReturnErr101Code(c, "Sorry, you cannot submit illegally")
		return
	}

	tools.ReturnErr101Code(c, "You have submitted successfully")
	return
}
