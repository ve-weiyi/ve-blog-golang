package oauth

import (
	"fmt"
	"log"

	"github.com/goccy/go-json"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth/result"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/https"
)

// Feishu授权登录
type AuthFeishu struct {
	BaseRequest
}

func NewAuthFeishu(conf *AuthConfig) *AuthFeishu {
	authRequest := &AuthFeishu{}
	authRequest.Set("", conf)

	authRequest.authorizeUrl = "https://passport.feishu.cn/suite/passport/oauth/authorize"
	authRequest.TokenUrl = "https://passport.feishu.cn/suite/passport/oauth/token"
	authRequest.userInfoUrl = "https://passport.feishu.cn/suite/passport/oauth/userinfo"
	authRequest.RefreshUrl = "https://passport.feishu.cn/suite/passport/oauth/token"

	return authRequest
}

// 获取登录地址
func (a *AuthFeishu) GetRedirectUrl(state string) string {
	url := https.NewHttpBuilder(a.authorizeUrl).
		AddParam("client_id", a.config.ClientID).
		AddParam("redirect_uri", a.config.RedirectUrl).
		AddParam("response_type", "code").
		AddParam("state", a.GetState(state)).
		GetUrl()
	return url
}

// 获取token https://open.weibo.com/apps/2658270041/privilege/oauth
func (a *AuthFeishu) GetAccessToken(code string) (resp *result.TokenResult, err error) {
	body, err := https.NewHttpBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("client_id", a.config.ClientID).
		AddParam("client_secret", a.config.ClientSecret).
		AddParam("redirect_uri", a.config.RedirectUrl).
		AddParam("code", code).
		Post()

	log.Println("err:", err)
	log.Println("body:", body)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 获取用户信息
func (a *AuthFeishu) RefreshToken(refreshToken string) (resp *result.RefreshResult, err error) {
	body, err := https.NewHttpBuilder(a.RefreshUrl).
		AddData("grant_type", "refresh_token").
		AddData("refresh_token", refreshToken).
		Post()

	log.Println("err:", err)
	log.Println("body:", body)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// 获取用户信息
func (a *AuthFeishu) GetUserInfo(accessToken string) (resp *result.UserResult, err error) {
	body, err := https.NewHttpBuilder(a.userInfoUrl).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		AddHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		Get()

	log.Println("err:", err)
	log.Println("body:", body)

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
