package example

import (
	"log"
	"testing"

	"github.com/ve-weiyi/ve-admin-store/server/infra/oauth"
)

func TestQQ(t *testing.T) {
	conf := &oauth.AuthConfig{ClientID: "xx", ClientSecret: "xx", RedirectUrl: "xx"}

	auth := oauth.NewAuthQq(conf)
	//获取第三方登录地址
	url := auth.GetRedirectUrl("sate")
	log.Println("url:", url)

	//获取token信息
	//tokenRes, err := auth.GetAccessToken("38F5FECC409FACC5A25B77B1EC9E26AE")
	//log.Println("tokenRes:", err, tokenRes)

	//刷新token
	//refresh, err := auth.RefreshToken("AE6F4302DD7AFB52902F56150FC58D2A")
	//log.Println("refresh:", err, refresh)

	//获取用户信息
	userInfo, err := auth.GetUserInfo("1E1D9C86BDADDD48719D7027BDF9CE48")

	log.Println("userInfo:", err, userInfo)

}
