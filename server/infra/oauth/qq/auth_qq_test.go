package qq

import (
	"log"
	"testing"

	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/initest"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth"
)

func TestQQ(t *testing.T) {
	initest.InitConfig()
	conf := &oauth.AuthConfig{
		ClientID:     global.CONFIG.Oauth.QQ.ClientID,
		ClientSecret: global.CONFIG.Oauth.QQ.ClientSecret,
		RedirectUrl:  "https://www.ve77.cn/blog/oauth/login/qq",
	}

	auth := NewAuthQq(conf)
	//获取第三方登录地址
	url := auth.GetRedirectUrl("state")
	log.Println("url:", url)
	//获取用户信息
	//userInfo, err := auth.GetUserOpenInfo("9C68E4021CB7DC5B1E7EC3FA34FE4C3E")
	//log.Println("userInfo:", err, userInfo)

	//获取token信息
	//tokenRes, err := auth.GetAccessToken("9C68E4021CB7DC5B1E7EC3FA34FE4C3E")
	//log.Println("tokenRes:", err, tokenRes)

	//刷新token
	refresh, err := auth.RefreshToken("9AEA28A71B91AF087CB6D3986BA62D24")
	log.Println("refresh:", err, refresh)
}
