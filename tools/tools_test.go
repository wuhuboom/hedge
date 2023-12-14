package tools

import (
	"fmt"
	"testing"
)

func For5ToFunc(a func() string) {
	for i := 0; i < 6; i++ {
		fmt.Println(a())
	}
}

func A() string {
	return "------"
}

func B(a int, b ...int) {
	fmt.Println(a, b)
}

func TestName(t *testing.T) {

	UseTelegramNoticeAdmin("6694559343:AAHJ1ipv7hecz5gJINJwh_esQCMCSTA66wI", " @xiaoleng16 %0A @xiaoleng16", "-4039362456", "565656")

}
