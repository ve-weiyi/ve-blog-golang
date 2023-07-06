package example

import (
	"log"
	"testing"

	"github.com/ve-weiyi/ve-admin-store/server/infra/oauth"
)

func TestFeishu(t *testing.T) {
	conf := &oauth.AuthConfig{ClientID: "xx", ClientSecret: "xx", RedirectUrl: "xx"}

	auth := oauth.NewAuthFeishu(conf)
	//获取第三方登录地址
	url := auth.GetRedirectUrl("sate")
	log.Println("url:", url)

	//获取token信息
	tokenRes, err := auth.GetAccessToken("064l6ca998d149d3bfa7298858695d8b")
	//{"access_token":"dXOXo.h9leZpCrz_jj1jdA4lnPcl51bHqU0015y82xaM","refresh_token":"cy3fsds_lfMH0cKgAVS9wP4llVAl5119UE005lO82xuB","token_type":"Bearer","expires_in":7200,"refresh_expires_in":2592000,"scope":"auth:user.id:read user_profile"}
	if err != nil {
		return
	}
	log.Println("ssss:", err, tokenRes)
	//获取用户信息
	userInfo, err := auth.GetUserInfo("2TMmI0ryJfxpyWAJxmYOG51h3GuR519xPy00h5w02KF_")
	if err != nil {
		return
	}
	log.Println("ssss:", err, userInfo)

	auth.RefreshToken(tokenRes.RefreshToken)
}
