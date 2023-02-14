package tools

import (
	"fmt"
	"net/http"
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

	get, _ := http.Get("https://api.adminjjjjsdj.xyz/player/auth/sys_config")

	fmt.Println(get.Status)

}
