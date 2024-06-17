package request

type AboutMeReq struct {
	Content string `json:"content" from:"content" example:"about me"`
}

type WebsiteConfigReq struct {
	Key   string `json:"key" from:"key" example:"about"`
	Value string `json:"value" from:"value" example:"about me"`
}

// VoiceVO 表示系统中的音频数据。
type VoiceVO struct {
	Type int64 `json:"type" validate:"required"` // 消息类型
	//File *multipart.FileHeader `json:"file" validate:"required"` // 文件
	//UserId    int64                   `json:"user_id" validate:"required"`    // 用户id
	//Nickname  string                `json:"nickname" validate:"required"`   // 用户昵称
	//Avatar    string                `json:"avatar" validate:"required"`     // 用户头像
	Content string `json:"content" validate:"required"` // 聊天内容
	//CreatedAt time.Time             `json:"created_at" validate:"required"` // 创建时间
	//IPAddress string                `json:"ip_address" validate:"required"` // 用户登录ip
	//IPSource  string                `json:"ip_source" validate:"required"`  // ip来源
}

type WebsiteConfigDTO struct {
	AdminUrl          string      `json:"admin_url"`           // 后台地址
	AlipayQrCode      string      `json:"alipay_qr_code"`      // 支付宝二维码
	Gitee             string      `json:"gitee"`               // Gitee
	Github            string      `json:"github"`              // Github
	IsChatRoom        int64       `json:"is_chat_room"`        // 是否开启聊天室
	IsCommentReview   int64       `json:"is_comment_review"`   // 是否开启评论审核
	IsEmailNotice     int64       `json:"is_email_notice"`     // 是否开启邮件通知
	IsMessageReview   int64       `json:"is_message_review"`   // 是否开启留言审核
	IsMusicPlayer     int64       `json:"is_music_player"`     // 是否开启音乐播放器
	IsReward          int64       `json:"is_reward"`           // 是否开启打赏
	Qq                string      `json:"qq"`                  // QQ
	SocialLoginList   []string    `json:"social_login_list"`   // 社交登录列表
	SocialUrlList     []string    `json:"social_url_list"`     // 社交地址列表
	TouristAvatar     string      `json:"tourist_avatar"`      // 游客头像
	UserAvatar        string      `json:"user_avatar"`         // 用户头像
	WebsiteAuthor     string      `json:"website_author"`      // 网站作者
	WebsiteAvatar     interface{} `json:"website_avatar"`      // 网站头像
	WebsiteCreateTime string      `json:"website_create_time"` // 网站创建时间
	WebsiteIntro      string      `json:"website_intro"`       // 网站介绍
	WebsiteName       string      `json:"website_name"`        // 网站名称
	WebsiteNotice     string      `json:"website_notice"`      // 网站公告
	WebsiteRecordNo   string      `json:"website_record_no"`   // 网站备案号
	WebsocketUrl      string      `json:"websocket_url"`       // websocket地址
	WeixinQrCode      string      `json:"weixin_qr_code"`      // 微信二维码
}
