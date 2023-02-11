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
	//For5ToFunc(A)

	B(1, 10)
}
