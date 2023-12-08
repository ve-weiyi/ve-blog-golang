package weibo

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/https"
)

// 微博授权登录
type AuthWb struct {
	oauth.AuthOauth
}

func NewAuthWb(conf *oauth.AuthConfig) *AuthWb {
	authRequest := &AuthWb{}
	authRequest.Set("weibo", conf)

	authRequest.AuthorizeUrl = "https://api.weibo.com/oauth2/authorize"
	authRequest.TokenUrl = "https://api.weibo.com/oauth2/access_token"
	authRequest.UserInfoUrl = "https://api.weibo.com/2/users/show.json"

	return authRequest
}

// 获取登录地址
func (a *AuthWb) GetRedirectUrl(state string) string {
	url := https.NewHttpBuilder(a.AuthorizeUrl).
		AddParam("response_type", "code").
		AddParam("client_id", a.Config.ClientID).
		AddParam("redirect_uri", a.Config.RedirectUrl).
		AddParam("state", state).
		GetUrl()

	return url
}

// 获取用户信息
func (a *AuthWb) GetUserOpenInfo(code string) (resp *oauth.UserResult, err error) {
	tk, err := a.GetAccessToken(code)
	if err != nil {
		return nil, err
	}

	user, err := a.GetUserInfo(tk.AccessToken, tk.Uid)
	if err != nil {
		return nil, err
	}

	resp = &oauth.UserResult{
		OpenID:   strconv.FormatInt(user.ID, 10),
		NickName: user.ScreenName,
		Name:     user.Name,
		Avatar:   user.AvatarLarge,
	}

	return resp, nil
}

// 获取token
func (a *AuthWb) GetAccessToken(code string) (resp *TokenResult, err error) {
	body, err := https.NewHttpBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("code", code).
		AddParam("client_id", a.Config.ClientID).
		AddParam("client_secret", a.Config.ClientSecret).
		AddParam("redirect_uri", a.Config.RedirectUrl).
		Post()
	if err != nil {
		return nil, err
	}

	log.Println("body:", string(body))
	var token TokenResult
	err = json.Unmarshal(body, &token)
	if err != nil {
		return nil, err
	}

	return &token, nil
}

// 获取第三方用户信息
func (a *AuthWb) GetUserInfo(accessToken string, openID string) (resp *UserResult, err error) {

	body, err := https.NewHttpBuilder(a.UserInfoUrl).
		AddParam("uid", openID).
		AddParam("access_token", accessToken).
		Get()
	if err != nil {
		return nil, err
	}

	log.Println("body:", string(body))

	user := UserResult{}

	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
