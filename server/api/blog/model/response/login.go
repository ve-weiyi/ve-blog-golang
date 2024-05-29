package response

// 用户登录信息
type LoginResp struct {
	Token        *Token        `json:"token"`
	UserInfo     *UserInfo     `json:"user_info"`
	LoginHistory *LoginHistory `json:"login_history"`
}

type Token struct {
	UserId           int    `json:"user_id"`            // 用户id
	TokenType        string `json:"token_type"`         // token类型,Bearer
	AccessToken      string `json:"access_token"`       // 访问token,过期时间较短。2h
	ExpiresIn        int64  `json:"expires_in"`         // 访问token过期时间
	RefreshToken     string `json:"refresh_token"`      // 刷新token,过期时间较长。30d
	RefreshExpiresIn int64  `json:"refresh_expires_in"` // 刷新token过期时间
	Scope            string `json:"scope"`              // 作用域
}

type UserInfo struct {
	UserId   int    `json:"user_id"`  // 用户id
	Username string `json:"username"` // 用户名
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
	Intro    string `json:"intro"`    // 个人简介
	Website  string `json:"website"`  // 个人网站
	Email    string `json:"email"`    // 邮箱

	//ArticleLikeSet []string `json:"article_like_set"` // 文章点赞集合
	//CommentLikeSet []string `json:"comment_like_set"` // 评论点赞集合
	//TalkLikeSet    []string `json:"talk_like_set"`    // 说说点赞集合

	Roles []*RoleDTO `json:"roles"` // 角色列表
}

type LoginHistory struct {
	Id        int    `json:"id"`
	LoginType string `json:"login_type"` // 登录类型
	Agent     string `json:"agent"`      // 代理
	IpAddress string `json:"ip_address"` // ip host
	IpSource  string `json:"ip_source"`  // ip 源
	LoginTime string `json:"login_time"` // 创建时间
}

type OauthLoginUrl struct {
	Url string `json:"url" example:""` // 授权地址
}
