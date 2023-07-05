package initialize

import (
	"github.com/ve-weiyi/ve-admin-store/server/global"
	"github.com/ve-weiyi/ve-admin-store/server/infra/jjwt"
)

func JwtToken() {
	global.JWT = &jjwt.JwtToken{
		SigningKey:  []byte(global.CONFIG.JWT.SigningKey),
		TokenPrefix: "",
		Issuer:      "blog",
	}
}
