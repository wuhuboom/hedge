package tools

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	eeor "github.com/wangyi/GinTemplate/error"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"sort"
	"strings"
	"time"
	"unicode"
)

var (
	VerifyErrCode     = -1002 //参数错误
	SuccessReturnCode = 200   //成功
	NeedGoogleBind    = 201   //绑定谷歌
	IllegalityCode    = -1001 //非法请求
	TokenExpire       = -1003 //token  过期
	IpLimitWaring     = -1004 //ip 限制
)

// GetRunPath2 获取程序执行目录
func GetRunPath2() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return ret
}

// IsFileNotExist 判断文件文件夹不存在
func IsFileNotExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return true, nil
	}
	return false, err
}

// IsFileExist 判断文件文件夹是否存在(字节0也算不存在)
func IsFileExist(path string) (bool, error) {
	fileInfo, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false, nil
	}
	//我这里判断了如果是0也算不存在
	if fileInfo.Size() == 0 {
		return false, nil
	}
	if err == nil {
		return true, nil
	}
	return false, err
}

// GetRootPath 获取程序根目录
func GetRootPath() string {
	rootPath, _ := os.Getwd()
	if notExist, _ := IsFileNotExist(rootPath); notExist {
		rootPath = GetRunPath2()
		if notExist, _ := IsFileNotExist(rootPath); notExist {
			rootPath = "."
		}
	}
	return rootPath
}

// JsonWrite json返回
func JsonWrite(context *gin.Context, status int, result interface{}, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code":   status,
		"result": result,
		"msg":    msg,
	})
}

func ReturnErr101Code(context *gin.Context, msg interface{}) {
	context.JSON(http.StatusOK, gin.H{
		"code":   -101,
		"result": nil,
		"msg":    msg,
	})
}

// ReturnVerifyErrCode 参数非法的错误
func ReturnVerifyErrCode(context *gin.Context, err error) {
	context.JSON(http.StatusOK, gin.H{
		"code":   VerifyErrCode,
		"result": nil,
		"msg":    err.Error(),
	})
}

func ReturnSuccess2000Code(context *gin.Context, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code":   SuccessReturnCode,
		"result": nil,
		"msg":    msg,
	})
}
func ReturnSuccess2000DataCode(context *gin.Context, data interface{}, msg string) {
	context.JSON(http.StatusOK, gin.H{
		"code":   SuccessReturnCode,
		"result": data,
		"msg":    msg,
	})
}

// RandStringRunes 生成随机字符串
func RandStringRunes(n int) string {
	var letterRunes = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}

	return string(b)
}

// IsArray 判断数组是否存在
func IsArray(array []string, arr string) bool {
	for _, s := range array {
		if s == arr {
			return true
		}
	}
	return false
}

func MD5(v string) string {
	d := []byte(v)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}

// AsciiKey map   排序
func AsciiKey(unSortedMap map[string]string) string {
	str := ""
	keys := make([]string, 0, len(unSortedMap))
	for k := range unSortedMap {
		if k != "sign" {
			keys = append(keys, k)
		}

	}
	sort.Strings(keys)
	for v, k := range keys {
		//fmt.Println(k, unSortedMap[k])
		if v == 0 {
			str = str + k + "=" + unSortedMap[k]
		} else {
			str = str + "&" + k + "=" + unSortedMap[k]
		}

	}

	return str
}

func HttpJsonPost(postUrl string, jsonData []byte) {
	for i := 0; i < 10; i++ {
		request, error := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonData))
		if error != nil {
			continue
		}
		request.Header.Set("Content-Type", "application/json; charset=UTF-8")
		client := &http.Client{}
		response, error := client.Do(request)
		if error != nil {
			continue
		}
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println("response Body:", string(body))
		if string(body) == "SUCCESS" {
			return
		}
		time.Sleep(1 * time.Second)
	}

}

func HttpJsonPostOne(postUrl string, jsonData []byte) string {
	request, error := http.NewRequest("POST", postUrl, bytes.NewBuffer(jsonData))
	defer request.Body.Close()
	if error != nil {
		return error.Error()
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		return error.Error()
	}
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))
	return string(body)
}

// CheckImageFile 检查  图片校验
func CheckImageFile(path, style string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	switch strings.ToUpper(style) {
	case "JPG", "JPEG":
		_, err = jpeg.Decode(f)
	case "PNG":
		_, err = png.Decode(f)
	case "GIF":
		_, err = gif.Decode(f)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}

// IsChinese 判断 是否 存在中文
func IsChinese(str string) bool {
	var count int
	for _, v := range str {
		if unicode.Is(unicode.Han, v) {
			count++
			break
		}
	}
	return count > 0
}
func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}
func ReturnDataLIst2000(context *gin.Context, data interface{}, count int) {
	context.JSON(http.StatusOK, gin.H{
		"code":   200,
		"result": data,
		"count":  count,
	})
}

func GetInvitationCode(num int, redis *redis.Client) (string, error) {
	for i := 0; i < 5; i++ {
		str := RandStringRunes(num)
		result, _ := redis.HExists("InvitationCode", str).Result()
		if !result {
			return str, nil
		}
	}
	return "", eeor.OtherError("Sorry, duplicate invitation code")
}
