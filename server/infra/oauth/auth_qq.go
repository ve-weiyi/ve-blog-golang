package oauth

import (
	"log"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/go-sdk/utils/jsonconv"
	"github.com/ve-weiyi/ve-admin-store/server/infra/oauth/https"
	"github.com/ve-weiyi/ve-admin-store/server/infra/oauth/result"
	"github.com/ve-weiyi/ve-admin-store/server/infra/oauth/utils"
)

// QQ授权登录
type AuthQq struct {
	BaseRequest
}

func NewAuthQq(conf *AuthConfig) *AuthQq {
	authRequest := &AuthQq{}
	authRequest.Set("qq", conf)

	authRequest.authorizeUrl = "https://graph.qq.com/oauth2.0/authorize"
	authRequest.TokenUrl = "https://graph.qq.com/oauth2.0/token"
	authRequest.RefreshUrl = "https://graph.qq.com/oauth2.0/token"
	authRequest.openidUrl = "https://graph.qq.com/oauth2.0/me"
	authRequest.userInfoUrl = "https://graph.qq.com/user/get_user_info"

	return authRequest
}

// 获取登录地址
func (a *AuthQq) GetRedirectUrl(state string) string {
	url := https.NewHttpBuilder(a.authorizeUrl).
		AddParam("response_type", "code").
		AddParam("client_id", a.config.ClientID).
		AddParam("redirect_uri", a.config.RedirectUrl).
		AddParam("state", a.GetState(state)).
		GetUrl()

	return url
}

// 获取token
func (a *AuthQq) GetAccessToken(code string) (resp *result.TokenResult, err error) {
	body, status := https.NewHttpBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("code", code).
		AddParam("client_id", a.config.ClientID).
		AddParam("client_secret", a.config.ClientSecret).
		AddParam("redirect_uri", a.config.RedirectUrl).
		AddParam("fmt", "json"). // 由于历史原因，加上这个参数则返回json格式数据
		Get()

	log.Println("status:", status)
	log.Println("body:", body)

	err = jsonconv.JsonToObject(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 刷新token
func (a *AuthQq) RefreshToken(refreshToken string) (resp *result.RefreshResult, err error) {
	body, status := https.NewHttpBuilder(a.RefreshUrl).
		AddParam("grant_type", "refresh_token").
		AddParam("client_id", a.config.ClientID).
		AddParam("client_secret", a.config.ClientSecret).
		AddParam("refresh_token", refreshToken).
		AddParam("fmt", "json").
		Get()

	log.Println("status:", status)
	log.Println("body:", body)

	// 由于QQ的返回值expire_in是字符串，所以不能直接把string解析到int上
	mss := utils.JsonToMSS(body)
	resp = &result.RefreshResult{
		AccessToken:  mss["access_token"],
		ExpiresIn:    cast.ToInt(mss["expires_in"]),
		RefreshToken: mss["refresh_token"],
	}
	return resp, nil
}

// 获取用户openid
func (a *AuthQq) GetOpenid(accessToken string) (resp *result.Credentials, err error) {
	body, status := https.NewHttpBuilder(a.openidUrl).
		AddParam("access_token", accessToken).
		AddParam("fmt", "json").
		Get()

	log.Println("status:", status)
	log.Println("body:", body)

	mss := utils.JsonToMSS(body)
	resp = &result.Credentials{
		OpenId:  mss["openid"],
		Unionid: mss["unionid"],
	}

	return resp, nil
}

// 获取第三方用户信息 https://wiki.connect.qq.com/get_user_info
func (a *AuthQq) GetUserInfo(accessToken string) (resp *result.UserResult, err error) {
	openresp, err := a.GetOpenid(accessToken)

	body, status := https.NewHttpBuilder(a.userInfoUrl).
		AddParam("openid", openresp.OpenId).
		AddParam("access_token", accessToken).
		AddParam("oauth_consumer_key", a.config.ClientID).
		Post()

	log.Println("status:", status)
	log.Println("body:", body)

	err = jsonconv.JsonToObject(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
