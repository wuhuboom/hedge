package h5

import (
	"bufio"
	"fmt"
	"github.com/gin-gonic/gin"
	jsoniter "github.com/json-iterator/go"
	"github.com/spf13/viper"
	"github.com/wangyi/GinTemplate/dao/mysql"
	"github.com/wangyi/GinTemplate/model"
	"github.com/wangyi/GinTemplate/tools"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

// Pay 玩家支付
func Pay(c *gin.Context) {
	//参数校验
	var cpd CheckPayData
	err := c.BindJSON(&cpd)
	if err != nil {
		tools.ReturnVerifyErrCode(c, err)
		return
	}
	//   参数校验
	paid, err := cpd.CheckSign(mysql.DB, c.ClientIP())
	fmt.Println(paid.ExpirationTime)

	webUrl := viper.GetString("config.webUrl")
	if err != nil {
		if err.Error() == "You have an outstanding order" {
			url := fmt.Sprintf(webUrl+"?BankCode=%s&ThreeOrder=%s&Amount=%f&BankIfsc=%s&AccountName=%s&status=%s&expiration=%d", paid.BankCode, paid.ThreeOrder, paid.Amount, paid.BankIfsc, paid.AccountName, "1", paid.ExpirationTime)
			tools.ReturnSuccess2000DataCode(c, url, "ok")
			return
		}

		tools.ReturnErr101Code(c, err.Error())
		return
	}
	url := fmt.Sprintf(webUrl+"?BankCode=%s&ThreeOrder=%s&Amount=%f&BankIfsc=%s&AccountName=%s&expiration=%d", paid.BankCode, paid.ThreeOrder, paid.Amount, paid.BankIfsc, paid.AccountName, paid.ExpirationTime)
	//  支付订单拉起   通知三方冻结金额
	//通知消息
	notice := make(map[string]string)
	notice["PayOrderNum"] = paid.ThreeOrder //订单号
	notice["amount"] = cpd.Amount           //金额
	notice["action"] = "freeze"             //  冻结金额
	marshal, err := jsoniter.Marshal(notice)
	if err != nil {
		return
	}

	go tools.HttpJsonPost(paid.NoticeUrl, marshal)
	tools.ReturnSuccess2000DataCode(c, url, "ok")
	return

}

// UploadPayCertificate 上传支付凭证
func UploadPayCertificate(c *gin.Context) {
	//判断 提交的订单号  是否存在
	ThreeOrder := c.PostForm("ThreeOrder")
	pay := model.PayOrder{}
	err := mysql.DB.Where("three_order=?", ThreeOrder).First(&pay).Error
	if err != nil || pay.Status != 1 {
		tools.ReturnErr101Code(c, "fail")
		return
	}
	image, err := c.FormFile("file")
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
	path := "./static/upload/" + time.Now().Format("20060102") + "/"
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

	//  更新订单图片
	if pay.CertificateImage != "" {
		os.Remove(pay.CertificateImage)
	}

	//更新 图片
	err = mysql.DB.Model(&model.PayOrder{}).Where("id=?", pay.ID).Update(&model.PayOrder{CertificateImage: path}).Error
	if err != nil {
		tools.ReturnErr101Code(c, err.Error())
		return
	}

	//os.Remove(path)
	tools.ReturnSuccess2000Code(c, "ok")
	return

}

// CheckImageForUpload 查看已经上传的  截图
func CheckImageForUpload(c *gin.Context) {
	//判断 提交的订单号  是否存在
	ThreeOrder := c.PostForm("ThreeOrder")
	pay := model.PayOrder{}
	err := mysql.DB.Where("three_order=?", ThreeOrder).First(&pay).Error
	if err != nil {
		tools.ReturnErr101Code(c, "fail")
		return
	}

	tools.ReturnSuccess2000DataCode(c, pay.CertificateImage, "ok")
	return
}
