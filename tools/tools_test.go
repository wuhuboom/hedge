package tools

import (
	"fmt"
	"testing"
	"time"
)

func FuzzName(f *testing.F) {

	fmt.Println(5 * 60 * time.Second.Milliseconds())
	fmt.Println(5 * 60 * time.Second.Microseconds())
	fmt.Println(5 * 60 * time.Second.Nanoseconds())
}
