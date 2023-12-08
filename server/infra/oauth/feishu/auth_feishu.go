package feishu

import (
	"fmt"
	"log"

	"github.com/goccy/go-json"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/https"
)

// Feishu授权登录
type AuthFeishu struct {
	oauth.AuthOauth
}

func NewAuthFeishu(conf *oauth.AuthConfig) *AuthFeishu {
	authRequest := &AuthFeishu{}
	authRequest.Set("", conf)

	authRequest.AuthorizeUrl = "https://passport.feishu.cn/suite/passport/oauth/authorize"
	authRequest.TokenUrl = "https://passport.feishu.cn/suite/passport/oauth/token"
	authRequest.UserInfoUrl = "https://passport.feishu.cn/suite/passport/oauth/userinfo"
	authRequest.RefreshUrl = "https://passport.feishu.cn/suite/passport/oauth/token"

	return authRequest
}

// 获取登录地址
func (a *AuthFeishu) GetRedirectUrl(state string) string {
	url := https.NewHttpBuilder(a.AuthorizeUrl).
		AddParam("client_id", a.Config.ClientID).
		AddParam("redirect_uri", a.Config.RedirectUrl).
		AddParam("response_type", "code").
		AddParam("state", state).
		GetUrl()
	return url
}

// 获取用户信息
func (a *AuthFeishu) GetUserOpenInfo(code string) (resp *oauth.UserResult, err error) {
	token, err := a.GetAccessToken(code)
	if err != nil {
		return nil, err
	}

	user, err := a.GetUserInfo(token.AccessToken)
	if err != nil {
		return nil, err
	}

	resp = &oauth.UserResult{
		OpenID:   user.OpenID,
		NickName: user.EnName,
		Name:     user.Name,
		EnName:   user.EnName,
		Avatar:   user.AvatarURL,
		Email:    user.Email,
		Mobile:   user.Mobile,
	}

	return resp, nil
}

// 获取token https://open.weibo.com/apps/2658270041/privilege/oauth
func (a *AuthFeishu) GetAccessToken(code string) (resp *TokenResult, err error) {
	body, err := https.NewHttpBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("client_id", a.Config.ClientID).
		AddParam("client_secret", a.Config.ClientSecret).
		AddParam("redirect_uri", a.Config.RedirectUrl).
		AddParam("code", code).
		Post()
	if err != nil {
		return nil, err
	}

	log.Println("body:", string(body))

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 获取用户信息
func (a *AuthFeishu) RefreshToken(refreshToken string) (resp *RefreshResult, err error) {
	body, err := https.NewHttpBuilder(a.RefreshUrl).
		AddData("grant_type", "refresh_token").
		AddData("refresh_token", refreshToken).
		Post()
	if err != nil {
		return nil, err
	}

	log.Println("body:", string(body))

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *AuthFeishu) GetUserInfo(accessToken string) (resp *UserResult, err error) {
	body, err := https.NewHttpBuilder(a.UserInfoUrl).
		AddHeader("Content-Type", "application/json;charset=UTF-8").
		AddHeader("Authorization", fmt.Sprintf("Bearer %s", accessToken)).
		Get()
	if err != nil {
		return nil, err
	}

	log.Println("body:", string(body))

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
