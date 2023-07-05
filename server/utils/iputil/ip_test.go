package iputil

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//info, err := GetIpInfoByBaidu("24.48.0.1")
	info, err := GetIpInfoByBaidu("119.23.144.144")
	if err != nil {
		return
	}
	fmt.Printf("%+v", info)
}
