package qq

import (
	"encoding/json"
	"log"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/httpx"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth"
)

// QQ授权登录
type AuthQq struct {
	Config *oauth.AuthConfig
	oauth.AuthOauthURL
}

func NewAuthQq(conf *oauth.AuthConfig) *AuthQq {
	auth := oauth.AuthOauthURL{}

	auth.Name = "qq"
	auth.AuthorizeUrl = "https://graph.qq.com/oauth2.0/authorize"
	auth.AccessTokenUrl = "https://graph.qq.com/oauth2.0/token"
	auth.RefreshTokenUrl = "https://graph.qq.com/oauth2.0/token"
	auth.OpenidUrl = "https://graph.qq.com/oauth2.0/me"
	auth.UserInfoUrl = "https://graph.qq.com/user/get_user_info"

	return &AuthQq{
		Config:       conf,
		AuthOauthURL: auth,
	}
}

func (a *AuthQq) GetName() string {
	return a.Name
}

// 获取登录地址
func (a *AuthQq) GetAuthorizeUrl(state string) string {

	url := httpx.NewClient(
		httpx.WithParam("response_type", "code"),
		httpx.WithParam("client_id", a.Config.ClientId),
		httpx.WithParam("redirect_uri", a.Config.RedirectUri),
		httpx.WithParam("state", state)).
		EncodeURL(a.AuthorizeUrl)

	return url
}

func (a *AuthQq) GetUserOpenInfo(code string) (resp *oauth.UserResult, err error) {
	token, err := a.GetAccessToken(code)
	if err != nil {
		return nil, err
	}

	open, err := a.GetOpenid(token.AccessToken)
	if err != nil {
		return nil, err
	}

	user, err := a.GetUserInfo(token.AccessToken, open.OpenId)
	if err != nil {
		return nil, err
	}

	resp = &oauth.UserResult{
		OpenId:   open.OpenId,
		NickName: user.Nickname,
		Name:     user.Nickname,
		EnName:   user.Nickname,
		Avatar:   user.FigureURLQQ,
		Email:    "",
		Mobile:   "",
	}
	return resp, nil
}

// 获取token
func (a *AuthQq) GetAccessToken(code string) (resp *TokenResult, err error) {

	body, err := httpx.NewClient(
		httpx.WithParam("grant_type", "authorization_code"),
		httpx.WithParam("code", code),
		httpx.WithParam("client_id", a.Config.ClientId),
		httpx.WithParam("client_secret", a.Config.ClientSecret),
		httpx.WithParam("redirect_uri", a.Config.RedirectUri),
		httpx.WithParam("fmt", "json"), // 由于历史原因，加上这个参数则返回json格式数据
		httpx.WithMethod("GET"),
		httpx.WithURL(a.AccessTokenUrl),
	).DoRequest()

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

// 刷新token
func (a *AuthQq) RefreshToken(refreshToken string) (resp *RefreshResult, err error) {

	body, err := httpx.NewClient(
		httpx.WithParam("grant_type", "refresh_token"),
		httpx.WithParam("client_id", a.Config.ClientId),
		httpx.WithParam("client_secret", a.Config.ClientSecret),
		httpx.WithParam("refresh_token", refreshToken),
		httpx.WithParam("fmt", "json"),
		httpx.WithMethod("GET"),
		httpx.WithURL(a.RefreshTokenUrl),
	).DoRequest()
	if err != nil {
		return nil, err
	}

	log.Println("body:", string(body))

	// 由于QQ的返回值expire_in是字符串，所以不能直接把string解析到int上
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 获取用户openid
func (a *AuthQq) GetOpenid(accessToken string) (resp *OpenResult, err error) {

	body, err := httpx.NewClient(
		httpx.WithParam("access_token", accessToken),
		httpx.WithParam("fmt", "json"),
		httpx.WithMethod("GET"),
		httpx.WithURL(a.OpenidUrl),
	).DoRequest()
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

// 获取第三方用户信息 https://wiki.connect.qq.com/get_user_info
func (a *AuthQq) GetUserInfo(accessToken string, openId string) (resp *UserResult, err error) {

	body, err := httpx.NewClient(
		httpx.WithParam("openid", openId),
		httpx.WithParam("access_token", accessToken),
		httpx.WithParam("oauth_consumer_key", a.Config.ClientId),
		httpx.WithMethod("POST"),
		httpx.WithURL(a.UserInfoUrl),
	).DoRequest()

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
