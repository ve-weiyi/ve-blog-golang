package oauth

import (
	"errors"
	"log"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth/result"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth/utils"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/https"
)

// 微博授权登录
type AuthWb struct {
	BaseRequest
}

func NewAuthWb(conf *AuthConfig) *AuthWb {
	authRequest := &AuthWb{}
	authRequest.Set("weibo", conf)

	authRequest.authorizeUrl = "https://api.weibo.com/oauth2/authorize"
	authRequest.TokenUrl = "https://api.weibo.com/oauth2/access_token"
	authRequest.userInfoUrl = "https://api.weibo.com/2/users/show.json"

	return authRequest
}

// 获取登录地址
func (a *AuthWb) GetRedirectUrl(state string) string {
	url := https.NewHttpBuilder(a.authorizeUrl).
		AddParam("response_type", "code").
		AddParam("client_id", a.config.ClientID).
		AddParam("redirect_uri", a.config.RedirectUrl).
		AddParam("state", a.GetState(state)).
		GetUrl()

	return url
}

// 获取token
func (a *AuthWb) GetAccessToken(code string) (*result.TokenResult, error) {
	body, err := https.NewHttpBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("code", code).
		AddParam("client_id", a.config.ClientID).
		AddParam("client_secret", a.config.ClientSecret).
		AddParam("redirect_uri", a.config.RedirectUrl).
		Post()

	log.Println("err:", err)
	log.Println("body:", body)

	m := utils.JsonToMSS(string(body))
	if _, ok := m["error"]; ok {
		return nil, errors.New(m["error_description"])
	}

	token := &result.TokenResult{
		AccessToken:  m["access_token"],
		RefreshToken: m["refresh_token"],
		ExpiresIn:    cast.ToInt(m["expires_in"]),
		Scope:        m["scope"],
		TokenType:    m["token_type"],
	}
	return token, nil
}

func (a *AuthWb) RefreshToken(refreshToken string) (resp *result.RefreshResult, err error) {
	body, err := https.NewHttpBuilder(a.RefreshUrl).
		AddParam("grant_type", "refresh_token").
		AddParam("client_id", a.config.ClientID).
		AddParam("client_secret", a.config.ClientSecret).
		AddParam("refresh_token", refreshToken).
		Post()

	log.Println("err:", err)
	log.Println("body:", body)

	// 由于QQ的返回值expire_in是字符串，所以不能直接把string解析到int上
	mss := utils.JsonToMSS(string(body))
	resp = &result.RefreshResult{
		AccessToken:  mss["access_token"],
		ExpiresIn:    cast.ToInt(mss["expires_in"]),
		RefreshToken: mss["refresh_token"],
	}
	return resp, nil
}

// 获取用户openid
func (a *AuthWb) GetOpenid(accessToken string) (resp *result.Credentials, err error) {
	body, err := https.NewHttpBuilder(a.openidUrl).
		AddParam("access_token", accessToken).
		AddParam("fmt", "json").
		Get()

	log.Println("err:", err)
	log.Println("body:", body)

	mss := utils.JsonToMSS(string(body))
	resp = &result.Credentials{
		OpenId:  mss["openid"],
		Unionid: mss["unionid"],
	}

	return resp, nil
}

// 获取第三方用户信息
func (a *AuthWb) GetUserInfo(accessToken string) (resp *result.UserResult, err error) {
	openresp, err := a.GetOpenid(accessToken)

	body, err := https.NewHttpBuilder(a.userInfoUrl).
		AddParam("uid", openresp.OpenId).
		AddParam("access_token", accessToken).
		Get()

	log.Println("err:", err)
	log.Println("body:", body)

	m := utils.JsonToMSS(string(body))
	if _, ok := m["error"]; ok {
		return nil, errors.New(m["error_description"])
	}

	user := &result.UserResult{}
	return user, nil
}
