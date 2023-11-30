package user

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model/modelPay"
	"github.com/wangyi/GinTemplate/tools"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// UploadCertificate 玩家上传支付凭证
func UploadCertificate(c *gin.Context) {

	var lo UploadCertificateVerify
	//检查参数
	if err := c.ShouldBind(&lo); err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}
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

func UploadCertificateForImageRtt(c *gin.Context) {
	var lo UploadCertificateVerifyForImage
	//检查参数
	if err := c.ShouldBind(&lo); err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}
	//判断这个订单是否存在
	col := modelPay.Collection{}
	err := mysql.DB.Where("own_order=?", lo.OrderNum).First(&col).Error
	if err != nil {
		tools.ReturnErr101Code(c, "Sorry, you cannot submit illegally")
		return
	}

	image := lo.ProofOfPayment
	if err != nil {
		tools.ReturnErr101Code(c, "fail")
		return
	}
	//限制图片大小
	if image.Size > 102400 {
		tools.ReturnErr101Code(c, "Sorry, your picture is too big")
		return
	}
	//检查图片
	nameArray := strings.Split(image.Filename, ".")
	if len(nameArray) != 2 {
		tools.ReturnErr101Code(c, "Please upload the correct file")
		return
	}

	checkImage := []string{"png", "jpg", "jpeg"}
	if tools.IsArray(checkImage, strings.ToLower(nameArray[1])) == false {
		tools.ReturnErr101Code(c, "Sorry, the picture format is wrong")
		return
	}

	open, err := image.Open()
	if err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}
	all, err := ioutil.ReadAll(bufio.NewReader(open))
	if err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}
	if tools.CheckPhpInImage(all) == false {
		tools.ReturnErr101Code(c, "Please upload the correct file")
		return
	}
	//判断文件夹是否存在
	path := "static/upload/" + time.Now().Format("20060102") + "/"
	if noExist, _ := tools.IsFileNotExist(path); noExist {
		if err := os.MkdirAll(path, 0777); err != nil {
			tools.ReturnErr101Code(c, err.Error())
			return
		}
	}
	if tools.IsChinese(nameArray[0]) == true {
		tools.ReturnErr101Code(c, "Chinese illegality")
		return
	}
	path = path + time.Now().Format("20060102150405") + nameArray[0] + "." + nameArray[1]
	err = c.SaveUploadedFile(image, path)
	if err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}
	mysql.DB.Model(&modelPay.Collection{}).Where("id=?", col.ID).Update(&modelPay.Collection{
		ProofOfPaymentImageUrl: path,
		Updated:                time.Now().Unix(),
	})
	fmt.Println(path, lo.OrderNum)
	tools.ReturnSuccess2000Code(c, "Upload successful")
	return
}

func GetCertificateForImageRtt(c *gin.Context) {
	//判断这个订单是否存在
	col := modelPay.Collection{}
	err := mysql.DB.Where("own_order=?", c.PostForm("own_order")).First(&col).Error
	if err != nil {
		tools.ReturnErr101Code(c, "Sorry, you cannot submit illegally")
		return
	}

	//获取商户号的 api
	//mer := model.Merchant{}

	//err = mysql.DB.Where("merchant_num=?", col.MerChantNum).First(&mer).Error
	//if err==nil{
	//
	//	mysql.DB.Where("id=?",mer.GatewayId).First(&)
	//}
	//fmt.Println(mer)
	DA := make(map[string]string)
	DA["host"] = "https://admin.oppay.cc"
	DA["imageUrl"] = col.ProofOfPaymentImageUrl
	if col.ProofOfPaymentImageUrl == "" {
		tools.ReturnErr101Code(c, "")
		return
	}

	tools.ReturnSuccess2000DataCode(c, DA, "OK")
	return
}
