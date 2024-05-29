package initialize

import (
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/feishu"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/qq"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/weibo"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
)

func InitOauth(c map[string]config.OauthConf) map[string]oauth.Oauth {
	var om = make(map[string]oauth.Oauth)

	for k, v := range c {
		conf := &oauth.AuthConfig{
			ClientId:     v.ClientId,
			ClientSecret: v.ClientSecret,
			RedirectUri:  v.RedirectUri,
		}
		switch k {
		case "qq":
			auth := qq.NewAuthQq(conf)
			om["qq"] = auth
		case "weibo":
			auth := weibo.NewAuthWb(conf)
			om["weibo"] = auth
		case "feishu":
			auth := feishu.NewAuthFeishu(conf)
			om["feishu"] = auth
		}
	}
	return om
}
