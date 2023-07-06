package example

import (
	"log"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth"
)

func TestWeibo(t *testing.T) {
	conf := &oauth.AuthConfig{ClientID: "xx", ClientSecret: "xx", RedirectUrl: "xx"}

	auth := oauth.NewAuthWb(conf)
	//获取第三方登录地址
	url := auth.GetRedirectUrl("sate")
	log.Println("url:", url)

	//获取token信息
	//tokenRes, err := auth.GetAccessToken("6fcac5ca6bed2ef17e3907b9f88589ca")
	//log.Println("tokenRes:", err, tokenRes)
	//  {"access_token":"2.00OYpWYGPTpttCaf929b916cL6FMXD","remind_in":"157679999","expires_in":157679999,"uid":"6007017078","isRealName":"true"}
	//刷新token
	//refresh, err := auth.RefreshToken("AE6F4302DD7AFB52902F56150FC58D2A")
	//log.Println("refresh:", err, refresh)

	//获取用户信息
	//userInfo, err := auth.GetUserInfo("2.00OYpWYGPTpttCaf929b916cL6FMXD", "6007017078")
	//log.Println("userInfo:", err, userInfo)

}
