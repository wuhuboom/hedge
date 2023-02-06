package tools

import (
	"fmt"
	"strings"
	"testing"
)

func FuzzName(f *testing.F) {

	fmt.Println(strings.Replace("https://inpayapi.gvfootball.com", "api", "", 1))

}
