package weibo

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/httpx"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth"
)

// 微博授权登录
type AuthWb struct {
	Config *oauth.AuthConfig
	oauth.AuthOauthURL
}

func NewAuthWb(conf *oauth.AuthConfig) *AuthWb {
	auth := oauth.AuthOauthURL{}

	auth.Name = "weibo"
	auth.AuthorizeUrl = "https://api.weibo.com/oauth2/authorize"
	auth.AccessTokenUrl = "https://api.weibo.com/oauth2/access_token"
	auth.UserInfoUrl = "https://api.weibo.com/2/users/show.json"

	return &AuthWb{
		Config:       conf,
		AuthOauthURL: auth,
	}
}

func (a *AuthWb) GetName() string {
	return a.Name
}

// 获取登录地址
func (a *AuthWb) GetRedirectUrl(state string) string {

	url := httpx.NewClient(
		httpx.WithParam("response_type", "code"),
		httpx.WithParam("client_id", a.Config.ClientId),
		httpx.WithParam("redirect_uri", a.Config.RedirectUri),
		httpx.WithParam("state", state),
	).EncodeURL(a.AuthorizeUrl)

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
		OpenId:   strconv.FormatInt(user.Id, 10),
		NickName: user.ScreenName,
		Name:     user.Name,
		Avatar:   user.AvatarLarge,
	}

	return resp, nil
}

// 获取token
func (a *AuthWb) GetAccessToken(code string) (resp *TokenResult, err error) {

	body, err := httpx.NewClient(
		httpx.WithParam("grant_type", "authorization_code"),
		httpx.WithParam("code", code),
		httpx.WithParam("client_id", a.Config.ClientId),
		httpx.WithParam("client_secret", a.Config.ClientSecret),
		httpx.WithParam("redirect_uri", a.Config.RedirectUri),
		httpx.WithMethod("GET"),
		httpx.WithURL(a.AccessTokenUrl),
	).DoRequest()

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
func (a *AuthWb) GetUserInfo(accessToken string, openId string) (resp *UserResult, err error) {

	body, err := httpx.NewClient(
		httpx.WithParam("uid", openId),
		httpx.WithParam("access_token", accessToken),
		httpx.WithMethod("GET"),
		httpx.WithURL(a.UserInfoUrl),
	).DoRequest()
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
