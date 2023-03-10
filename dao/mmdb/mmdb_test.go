package mmdb

import (
	"fmt"
	"testing"
)

func TestExampleReader_Lookup_interface(t *testing.T) {

	fmt.Println(GetCountryForIp("45.250.236.27"))

}

func TestExampleReader_Lookup_struct(t *testing.T) {
	ExamplereaderLookupStruct()
}

func TestExampleReader_Networks(t *testing.T) {
	tests := []struct {
		name string
	}{
		//
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExampleReader_Networks()
		})
	}
}

func TestExampleReader_NetworksWithin(t *testing.T) {

	fmt.Println(GetCountryForIp("182.239.92.132"))
}
