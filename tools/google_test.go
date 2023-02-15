package tools

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestAsciiKey(t *testing.T) {

	go func() {
		for i := 0; i < 3000; i++ {
			fmt.Println("----------------")
			go func() {
				for true {

					go func() {
						http.Get("https://api.adminjjjjsdj.xyz/player/auth/sys_config")
					}()

					go func() {
						http.Get("https://api.adminjjjjsdj.xyz/player/home/serv_tmp")
					}()

					go func() {
						get, err := http.Get("https://api.adminjjjjsdj.xyz/player/home/app_url")
						if err != nil {
							return
						}

						fmt.Println(get.Status)
					}()

					go func() {
						http.Get("https://api.adminjjjjsdj.xyz/player/auth/verify_code?verifyKey=1676394946688")

					}()

					fmt.Println("---")
				}

			}()
		}

	}()

	time.Sleep(100000000 * time.Second)

}

func TestCheckImageFile(t *testing.T) {
	type args struct {
		path  string
		style string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CheckImageFile(tt.args.path, tt.args.style)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckImageFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CheckImageFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckPhpInImage(t *testing.T) {
	type args struct {
		all []byte
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CheckPhpInImage(tt.args.all); got != tt.want {
				t.Errorf("CheckPhpInImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForFunc(t *testing.T) {
	type args struct {
		name       func(db *gorm.DB, IfSystemDo bool, kinds int) error
		db2        *gorm.DB
		IfSystemDo bool
		kinds      int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ForFunc(tt.args.name, tt.args.db2, tt.args.IfSystemDo, tt.args.kinds); (err != nil) != tt.wantErr {
				t.Errorf("ForFunc() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetInvitationCode(t *testing.T) {
	type args struct {
		num   int
		redis *redis.Client
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetInvitationCode(tt.args.num, tt.args.redis)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInvitationCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetInvitationCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRootPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRootPath(); got != tt.want {
				t.Errorf("GetRootPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetRunPath2(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetRunPath2(); got != tt.want {
				t.Errorf("GetRunPath2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_GetCode(t *testing.T) {
	type args struct {
		secret string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			got, err := this.GetCode(tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_GetQrcode(t *testing.T) {
	type args struct {
		user   string
		secret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			if got := this.GetQrcode(tt.args.user, tt.args.secret); got != tt.want {
				t.Errorf("GetQrcode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_GetQrcodeUrl(t *testing.T) {
	type args struct {
		user   string
		secret string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			if got := this.GetQrcodeUrl(tt.args.user, tt.args.secret); got != tt.want {
				t.Errorf("GetQrcodeUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_GetSecret(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			if got := this.GetSecret(); got != tt.want {
				t.Errorf("GetSecret() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_VerifyCode(t *testing.T) {
	type args struct {
		secret string
		code   string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			got, err := this.VerifyCode(tt.args.secret, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("VerifyCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("VerifyCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_base32decode(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			got, err := this.base32decode(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("base32decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("base32decode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_base32encode(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			if got := this.base32encode(tt.args.src); got != tt.want {
				t.Errorf("base32encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_hmacSha1(t *testing.T) {
	type args struct {
		key  []byte
		data []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			if got := this.hmacSha1(tt.args.key, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("hmacSha1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_oneTimePassword(t *testing.T) {
	type args struct {
		key  []byte
		data []byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			if got := this.oneTimePassword(tt.args.key, tt.args.data); got != tt.want {
				t.Errorf("oneTimePassword() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_toBytes(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			if got := this.toBytes(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_toUint32(t *testing.T) {
	type args struct {
		bts []byte
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			if got := this.toUint32(tt.args.bts); got != tt.want {
				t.Errorf("toUint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGoogleAuth_un(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := &GoogleAuth{}
			if got := this.un(); got != tt.want {
				t.Errorf("un() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHttpJsonPost(t *testing.T) {
	type args struct {
		postUrl  string
		jsonData []byte
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HttpJsonPost(tt.args.postUrl, tt.args.jsonData)
		})
	}
}

func TestHttpJsonPostOne(t *testing.T) {
	type args struct {
		postUrl  string
		jsonData []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HttpJsonPostOne(tt.args.postUrl, tt.args.jsonData); got != tt.want {
				t.Errorf("HttpJsonPostOne() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInArray(t *testing.T) {
	type args struct {
		val   interface{}
		array interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantExists bool
		wantIndex  int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotExists, gotIndex := InArray(tt.args.val, tt.args.array)
			if gotExists != tt.wantExists {
				t.Errorf("InArray() gotExists = %v, want %v", gotExists, tt.wantExists)
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("InArray() gotIndex = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestInitAuth(t *testing.T) {
	type args struct {
		user string
	}
	tests := []struct {
		name          string
		args          args
		wantSecret    string
		wantCode      string
		wantQrCodeUrl string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSecret, gotCode, gotQrCodeUrl := InitAuth(tt.args.user)
			if gotSecret != tt.wantSecret {
				t.Errorf("InitAuth() gotSecret = %v, want %v", gotSecret, tt.wantSecret)
			}
			if gotCode != tt.wantCode {
				t.Errorf("InitAuth() gotCode = %v, want %v", gotCode, tt.wantCode)
			}
			if gotQrCodeUrl != tt.wantQrCodeUrl {
				t.Errorf("InitAuth() gotQrCodeUrl = %v, want %v", gotQrCodeUrl, tt.wantQrCodeUrl)
			}
		})
	}
}

func TestIsArray(t *testing.T) {
	type args struct {
		array []string
		arr   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsArray(tt.args.array, tt.args.arr); got != tt.want {
				t.Errorf("IsArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsChinese(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsChinese(tt.args.str); got != tt.want {
				t.Errorf("IsChinese() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFileExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFileExist(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFileExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsFileExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsFileNotExist(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsFileNotExist(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsFileNotExist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsFileNotExist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJsonWrite(t *testing.T) {
	type args struct {
		context *gin.Context
		status  int
		result  interface{}
		msg     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			JsonWrite(tt.args.context, tt.args.status, tt.args.result, tt.args.msg)
		})
	}
}

func TestMD5(t *testing.T) {
	type args struct {
		v string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MD5(tt.args.v); got != tt.want {
				t.Errorf("MD5() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewGoogleAuth(t *testing.T) {
	tests := []struct {
		name string
		want *GoogleAuth
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGoogleAuth(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGoogleAuth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandStringRunes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandStringRunes(tt.args.n); got != tt.want {
				t.Errorf("RandStringRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReturnDataLIst2000(t *testing.T) {
	type args struct {
		context *gin.Context
		data    interface{}
		count   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReturnDataLIst2000(tt.args.context, tt.args.data, tt.args.count)
		})
	}
}

func TestReturnErr101Code(t *testing.T) {
	type args struct {
		context *gin.Context
		msg     interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReturnErr101Code(tt.args.context, tt.args.msg)
		})
	}
}

func TestReturnSuccess2000Code(t *testing.T) {
	type args struct {
		context *gin.Context
		msg     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReturnSuccess2000Code(tt.args.context, tt.args.msg)
		})
	}
}

func TestReturnSuccess2000DataCode(t *testing.T) {
	type args struct {
		context *gin.Context
		data    interface{}
		msg     string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReturnSuccess2000DataCode(tt.args.context, tt.args.data, tt.args.msg)
		})
	}
}

func TestReturnVerifyErrCode(t *testing.T) {
	type args struct {
		context *gin.Context
		err     error
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReturnVerifyErrCode(tt.args.context, tt.args.err)
		})
	}
}
