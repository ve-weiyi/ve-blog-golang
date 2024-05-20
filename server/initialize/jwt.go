package initialize

import (
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
)

func JwtToken() {
	//duration, err := time.ParseDuration(daysToHours(global.CONFIG.JWT.ExpiresTime))
	//if err != nil {
	//	panic(err)
	//}

	global.JWT = &jjwt.JwtToken{
		SigningKey: []byte(global.CONFIG.JWT.SigningKey),
		//Issuer:      global.CONFIG.JWT.Issuer,
		//ExpiresTime: duration,
		//TokenPrefix: "",
	}
}

func daysToHours(days string) string {
	daysValue := 0
	fmt.Sscanf(days, "%dd", &daysValue)
	hours := daysValue * 24
	return fmt.Sprintf("%dh", hours)
}
