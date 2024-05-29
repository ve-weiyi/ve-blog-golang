package initialize

//func InitOauth(c config.Config) map[string]oauth.Oauth {
//	var om = make(map[string]oauth.Oauth)
//
//	for k, v := range c.OauthConf {
//		conf := &oauth.AuthConfig{
//			ClientId:     v.ClientId,
//			ClientSecret: v.ClientSecret,
//			RedirectUri:  v.RedirectUri,
//		}
//		switch k {
//		case "qq":
//			auth := qq.NewAuthQq(conf)
//			om["qq"] = auth
//		case "weibo":
//			auth := weibo.NewAuthWb(conf)
//			om["weibo"] = auth
//		case "feishu":
//			auth := feishu.NewAuthFeishu(conf)
//			om["feishu"] = auth
//		}
//	}
//	return om
//}
