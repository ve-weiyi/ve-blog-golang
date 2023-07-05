package oauth

import (
	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth/result"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/oauth/source"
)

// 基本配置
type AuthConfig struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
}

type BaseRequest struct {
	authorizeUrl   string      //授权登录URL
	TokenUrl       string      //获得令牌URL
	AccessTokenUrl string      //获得访问令牌URL
	RefreshUrl     string      //刷新令牌URL
	openidUrl      string      //获取用户OPENID
	userInfoUrl    string      //获取用户信息URL
	config         *AuthConfig //配置信息
	registerSource string      //注册来源
}

func (b *BaseRequest) Set(sourceName source.RegisterSource, cfg *AuthConfig) {
	b.config = cfg
	b.registerSource = string(sourceName)
}

func (*BaseRequest) GetState(state string) string {
	if state == "" {
		return "state"
	}
	return state
}

// 获取第三方登录地址
type Oauth interface {
	GetRedirectUrl(state string) string
	GetAccessToken(code string) (*result.TokenResult, error)
	GetUserInfo(accessToken string) (*result.UserResult, error)
	RefreshToken(refreshToken string) (*result.RefreshResult, error)
}
