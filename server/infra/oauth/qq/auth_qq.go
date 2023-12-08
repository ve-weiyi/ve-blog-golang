package qq

import (
	"log"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/https"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
)

// QQ授权登录
type AuthQq struct {
	oauth.AuthOauth
}

func NewAuthQq(conf *oauth.AuthConfig) *AuthQq {
	authRequest := &AuthQq{}
	authRequest.Set("qq", conf)

	authRequest.AuthorizeUrl = "https://graph.qq.com/oauth2.0/authorize"
	authRequest.TokenUrl = "https://graph.qq.com/oauth2.0/token"
	authRequest.RefreshUrl = "https://graph.qq.com/oauth2.0/token"
	authRequest.OpenidUrl = "https://graph.qq.com/oauth2.0/me"
	authRequest.UserInfoUrl = "https://graph.qq.com/user/get_user_info"

	return authRequest
}

// 获取登录地址
func (a *AuthQq) GetRedirectUrl(state string) string {
	url := https.NewHttpBuilder(a.AuthorizeUrl).
		AddParam("response_type", "code").
		AddParam("client_id", a.Config.ClientID).
		AddParam("redirect_uri", a.Config.RedirectUrl).
		AddParam("state", state).
		GetUrl()

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
		OpenID:   open.OpenId,
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
	body, err := https.NewHttpBuilder(a.TokenUrl).
		AddParam("grant_type", "authorization_code").
		AddParam("code", code).
		AddParam("client_id", a.Config.ClientID).
		AddParam("client_secret", a.Config.ClientSecret).
		AddParam("redirect_uri", a.Config.RedirectUrl).
		AddParam("fmt", "json"). // 由于历史原因，加上这个参数则返回json格式数据
		Get()
	if err != nil {
		return nil, err
	}

	log.Println("body:", string(body))

	err = jsonconv.JsonToObject(string(body), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 刷新token
func (a *AuthQq) RefreshToken(refreshToken string) (resp *RefreshResult, err error) {
	body, err := https.NewHttpBuilder(a.RefreshUrl).
		AddParam("grant_type", "refresh_token").
		AddParam("client_id", a.Config.ClientID).
		AddParam("client_secret", a.Config.ClientSecret).
		AddParam("refresh_token", refreshToken).
		AddParam("fmt", "json").
		Get()
	if err != nil {
		return nil, err
	}

	log.Println("body:", string(body))

	// 由于QQ的返回值expire_in是字符串，所以不能直接把string解析到int上
	err = jsonconv.JsonToObject(string(body), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 获取用户openid
func (a *AuthQq) GetOpenid(accessToken string) (resp *OpenResult, err error) {
	body, err := https.NewHttpBuilder(a.OpenidUrl).
		AddParam("access_token", accessToken).
		AddParam("fmt", "json").
		Get()
	if err != nil {
		return nil, err
	}

	log.Println("body:", string(body))

	err = jsonconv.JsonToObject(string(body), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// 获取第三方用户信息 https://wiki.connect.qq.com/get_user_info
func (a *AuthQq) GetUserInfo(accessToken string, openId string) (resp *UserResult, err error) {
	body, err := https.NewHttpBuilder(a.UserInfoUrl).
		AddParam("openid", openId).
		AddParam("access_token", accessToken).
		AddParam("oauth_consumer_key", a.Config.ClientID).
		Post()
	if err != nil {
		return nil, err
	}

	log.Println("body:", string(body))

	err = jsonconv.JsonToObject(string(body), &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
