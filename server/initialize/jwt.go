package initialize

import (
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/jjwt"
)

func JwtToken() {
	global.JWT = &jjwt.JwtToken{
		SigningKey:  []byte(global.CONFIG.JWT.SigningKey),
		TokenPrefix: "",
		Issuer:      "blog",
	}
}
