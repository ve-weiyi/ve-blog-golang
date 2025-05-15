package dto

type Album struct {
	Id         int64  `json:"id"`          // 主键
	AlbumName  string `json:"album_name"`  // 相册名
	AlbumDesc  string `json:"album_desc"`  // 相册描述
	AlbumCover string `json:"album_cover"` // 相册封面
}

type AlbumQueryReq struct {
	PageQuery
}

type ArticleArchivesQueryReq struct {
	PageQuery
}

type ArticleClassifyQueryReq struct {
	PageQuery
	ClassifyName string `json:"classify_name,optional"` // 分类名
}

type ArticleDetails struct {
	ArticleHome
	LastArticle          *ArticlePreview   `json:"last_article"`           // 上一篇文章
	NextArticle          *ArticlePreview   `json:"next_article"`           // 下一篇文章
	RecommendArticleList []*ArticlePreview `json:"recommend_article_list"` // 推荐文章列表
	NewestArticleList    []*ArticlePreview `json:"newest_article_list"`    // 最新文章列表
}

type ArticleHome struct {
	Id             int64    `json:"id"`              // 文章ID
	ArticleCover   string   `json:"article_cover"`   // 文章缩略图
	ArticleTitle   string   `json:"article_title"`   // 标题
	ArticleContent string   `json:"article_content"` // 内容
	ArticleType    int64    `json:"article_type"`    // 文章类型
	OriginalUrl    string   `json:"original_url"`    // 原文链接
	IsTop          int64    `json:"is_top"`          // 是否置顶
	Status         int64    `json:"status"`          // 状态值 1 公开 2 私密 3 草稿 4 已删除
	CreatedAt      int64    `json:"created_at"`      // 发表时间
	UpdatedAt      int64    `json:"updated_at"`      // 更新时间
	CategoryName   string   `json:"category_name"`   // 文章分类名
	TagNameList    []string `json:"tag_name_list"`   // 文章标签列表
	LikeCount      int64    `json:"like_count"`      // 点赞量
	ViewsCount     int64    `json:"views_count"`     // 浏览量
}

type ArticleHomeQueryReq struct {
	PageQuery
	ArticleTitle string `json:"article_title,optional"` // 标题
}

type ArticlePreview struct {
	Id           int64  `json:"id"`            // 文章ID
	ArticleCover string `json:"article_cover"` // 文章缩略图
	ArticleTitle string `json:"article_title"` // 标题
	LikeCount    int64  `json:"like_count"`    // 点赞量
	ViewsCount   int64  `json:"views_count"`   // 浏览量
	CreatedAt    int64  `json:"created_at"`    // 创建时间
}

type BatchResp struct {
	SuccessCount int64 `json:"success_count"`
}

type Category struct {
	Id           int64  `json:"id"`
	CategoryName string `json:"category_name"` // 分类名
	ArticleCount int64  `json:"article_count"`
	CreatedAt    int64  `json:"created_at"` // 创建时间
	UpdatedAt    int64  `json:"updated_at"` // 更新时间
}

type CategoryQueryReq struct {
	PageQuery
	CategoryName string `json:"category_name,optional"` // 分类名
}

type ChatRecordResp struct {
	Id          int64  `json:"id"`           // 主键
	UserId      string `json:"user_id"`      // 用户id
	DeviceId    string `json:"device_id"`    // 设备id
	Nickname    string `json:"nickname"`     // 昵称
	Avatar      string `json:"avatar"`       // 头像
	ChatContent string `json:"chat_content"` // 消息内容
	IpAddress   string `json:"ip_address"`   // ip地址
	IpSource    string `json:"ip_source"`    // ip来源
	Type        string `json:"type"`         // 类型
	CreatedAt   int64  `json:"created_at"`   // 创建时间
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间
}

type ClientInfoResp struct {
	ClientId  string `json:"client_id"`  // 客户端id
	UserId    string `json:"user_id"`    // 用户id
	DeviceId  string `json:"device_id"`  // 设备id
	Nickname  string `json:"nickname"`   // 昵称
	IpAddress string `json:"ip_address"` // ip地址
	IpSource  string `json:"ip_source"`  // ip来源
}

type Comment struct {
	Id               int64           `json:"id"`                 // 评论id
	TopicId          int64           `json:"topic_id"`           // 主题id
	ParentId         int64           `json:"parent_id"`          // 父评论id
	ReplyMsgId       int64           `json:"reply_msg_id"`       // 会话id
	UserId           string          `json:"user_id"`            // 用户id
	ReplyUserId      string          `json:"reply_user_id"`      // 被回复用户id
	CommentContent   string          `json:"comment_content"`    // 评论内容
	Type             int64           `json:"type"`               // 评论类型 1.文章 2.友链 3.说说
	CreatedAt        int64           `json:"created_at"`         // 评论时间
	LikeCount        int64           `json:"like_count"`         // 点赞数
	User             *UserInfoVO     `json:"user"`               // 评论用户
	ReplyUser        *UserInfoVO     `json:"reply_user"`         // 被回复评论用户
	ReplyCount       int64           `json:"reply_count"`        // 回复量
	CommentReplyList []*CommentReply `json:"comment_reply_list"` // 评论回复列表
}

type CommentNewReq struct {
	TopicId        int64  `json:"topic_id,optional"`      // 主题id
	ParentId       int64  `json:"parent_id,optional"`     // 父评论id
	ReplyMsgId     int64  `json:"reply_msg_id,optional"`  // 会话id
	ReplyUserId    string `json:"reply_user_id,optional"` // 回复用户id
	CommentContent string `json:"comment_content"`        // 评论内容
	Type           int64  `json:"type"`                   // 评论类型 1.文章 2.友链 3.说说
	Status         int64  `json:"status,optional"`        // 状态 0.正常 1.已编辑 2.已删除
}

type CommentQueryReq struct {
	PageQuery
	TopicId  int64 `json:"topic_id,optional"`  // 主题id
	ParentId int64 `json:"parent_id,optional"` // 父评论id
	Type     int64 `json:"type,optional"`      // 评论类型 1.文章 2.友链 3.说说
}

type CommentReply struct {
	Id             int64       `json:"id"`              // 评论id
	TopicId        int64       `json:"topic_id"`        // 主题id
	ParentId       int64       `json:"parent_id"`       // 父评论id
	ReplyMsgId     int64       `json:"reply_msg_id"`    // 会话id
	UserId         string      `json:"user_id"`         // 用户id
	ReplyUserId    string      `json:"reply_user_id"`   // 被回复用户id
	CommentContent string      `json:"comment_content"` // 评论内容
	Type           int64       `json:"type"`            // 评论类型 1.文章 2.友链 3.说说
	CreatedAt      int64       `json:"created_at"`      // 评论时间
	LikeCount      int64       `json:"like_count"`      // 点赞数
	User           *UserInfoVO `json:"user"`            // 用户信息
	ReplyUser      *UserInfoVO `json:"reply_user"`      // 被回复评论用户
}

type DeleteUserBindThirdPartyReq struct {
	Platform string `json:"platform"` // 平台
}

type EmailLoginReq struct {
	Email       string `json:"email"`                 // 邮箱
	Password    string `json:"password"`              // 密码
	CaptchaKey  string `json:"captcha_key,optional"`  // 验证码key
	CaptchaCode string `json:"captcha_code,optional"` // 验证码
}

type EmptyReq struct {
}

type EmptyResp struct {
}

type FileBackVO struct {
	Id        int64  `json:"id,optional"` // 文件目录ID
	UserId    string `json:"user_id"`     // 用户id
	FilePath  string `json:"file_path"`   // 文件路径
	FileName  string `json:"file_name"`   // 文件名称
	FileType  string `json:"file_type"`   // 文件类型
	FileSize  int64  `json:"file_size"`   // 文件大小
	FileMd5   string `json:"file_md5"`    // 文件md5值
	FileUrl   string `json:"file_url"`    // 上传路径
	CreatedAt int64  `json:"created_at"`  // 创建时间
	UpdatedAt int64  `json:"updated_at"`  // 更新时间
}

type Friend struct {
	Id          int64  `json:"id"`           // id
	LinkName    string `json:"link_name"`    // 链接名
	LinkAvatar  string `json:"link_avatar"`  // 链接头像
	LinkAddress string `json:"link_address"` // 链接地址
	LinkIntro   string `json:"link_intro"`   // 链接介绍
	CreatedAt   int64  `json:"created_at"`   // 创建时间
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间
}

type FriendQueryReq struct {
	PageQuery
}

type GetAboutMeReq struct {
}

type GetAboutMeResp struct {
	Content string `json:"content"`
}

type GetBlogHomeInfoReq struct {
}

type GetBlogHomeInfoResp struct {
	ArticleCount       int64           `json:"article_count"`         // 文章数量
	CategoryCount      int64           `json:"category_count"`        // 分类数量
	TagCount           int64           `json:"tag_count"`             // 标签数量
	TotalUserViewCount int64           `json:"total_user_view_count"` // 总服务量
	TotalPageViewCount int64           `json:"total_page_view_count"` // 总浏览量
	WebsiteConfig      WebsiteConfigVO `json:"website_config"`        // 网站配置
	PageList           []*PageVO       `json:"page_list"`             // 页面列表
}

type GetCaptchaCodeReq struct {
	Width  int64 `json:"width,optional"`  // 宽度
	Height int64 `json:"height,optional"` // 高度
}

type GetCaptchaCodeResp struct {
	CaptchaKey    string `json:"captcha_key"`    // 验证码key
	CaptchaBase64 string `json:"captcha_base64"` // 验证码base64
	CaptchaCode   string `json:"captcha_code"`   // 验证码
}

type GetOauthAuthorizeUrlReq struct {
	Platform string `json:"platform"`       // 平台
	State    string `json:"state,optional"` // 状态
}

type GetOauthAuthorizeUrlResp struct {
	AuthorizeUrl string `json:"authorize_url"` // 授权地址
}

type GetTouristInfoResp struct {
	TouristId string `json:"tourist_id"` // 游客id
}

type IdReq struct {
	Id int64 `json:"id"`
}

type IdsReq struct {
	Ids []int64 `json:"ids"`
}

type LoginReq struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	CaptchaKey  string `json:"captcha_key,optional"`  // 验证码key
	CaptchaCode string `json:"captcha_code,optional"` // 验证码
}

type LoginResp struct {
	Token *Token `json:"token"`
}

type MultiUploadFileReq struct {
	Files    []interface{} `form:"files,optional"`     // 文件列表
	FilePath string        `form:"file_path,optional"` // 文件路径
}

type OnlineCountResp struct {
	Msg   string `json:"msg"`   // 消息
	Count int    `json:"count"` // 在线人数
}

type Page struct {
	Id         int64  `json:"id"`                   // 页面id
	PageName   string `json:"page_name"`            // 页面名
	PageLabel  string `json:"page_label"`           // 页面标签
	PageCover  string `json:"page_cover"`           // 页面封面
	IsCarousel int64  `json:"is_carousel,optional"` // 是否轮播
	CreatedAt  int64  `json:"created_at"`           // 创建时间
	UpdatedAt  int64  `json:"updated_at"`           // 更新时间
}

type PageQuery struct {
	Page     int64    `json:"page,optional"`
	PageSize int64    `json:"page_size,optional"`
	Sorts    []string `json:"sorts,optional"`
}

type PageQueryReq struct {
	PageQuery
}

type PageResp struct {
	Page     int64       `json:"page,omitempty"`
	PageSize int64       `json:"page_size,omitempty"`
	Total    int64       `json:"total,omitempty"`
	List     interface{} `json:"list,omitempty"`
}

type PageVO struct {
	Id         int64  `json:"id"`                   // 页面ID
	PageName   string `json:"page_name"`            // 页面名称
	PageLabel  string `json:"page_label"`           // 页面标签
	PageCover  string `json:"page_cover"`           // 页面封面
	IsCarousel int64  `json:"is_carousel,optional"` // 是否轮播
}

type PhoneLoginReq struct {
	Phone      string `json:"phone"`       // 手机号
	VerifyCode string `json:"verify_code"` // 验证码
}

type Photo struct {
	Id       int64  `json:"id"`        // 主键
	PhotoUrl string `json:"photo_url"` // 照片地址
}

type PhotoQueryReq struct {
	AlbumId int64 `json:"album_id"` // 相册ID
}

type PingReq struct {
}

type PingResp struct {
	Env         string   `json:"env"`
	Name        string   `json:"name"`
	Version     string   `json:"version"`
	Runtime     string   `json:"runtime"`
	Description string   `json:"description"`
	RpcStatus   []string `json:"rpc_status"`
}

type RecallMessageReq struct {
	Id int64 `json:"id"` // 消息id
}

type RecallMessageResp struct {
	Id int64 `json:"id"` // 消息id
}

type ReceiveMsg struct {
	Type      int    `json:"type"`      // 类型
	Data      string `json:"data"`      // 数据
	Timestamp int64  `json:"timestamp"` //时间戳
}

type RegisterReq struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"` // 确认密码
	Email           string `json:"email"`            // 邮箱
	VerifyCode      string `json:"verify_code"`      // 验证码
}

type Remark struct {
	Id             int64       `json:"id,optional"`     // 主键id
	UserId         string      `json:"user_id"`         // 用户id
	TerminalId     string      `json:"terminal_id"`     // 终端id
	MessageContent string      `json:"message_content"` // 留言内容
	IpAddress      string      `json:"ip_address"`      // 用户ip
	IpSource       string      `json:"ip_source"`       // 用户地址
	IsReview       int64       `json:"is_review"`       // 是否审核
	CreatedAt      int64       `json:"created_at"`      // 发布时间
	UpdatedAt      int64       `json:"updated_at"`      // 更新时间
	User           *UserInfoVO `json:"user"`            // 用户信息
}

type RemarkNewReq struct {
	MessageContent string `json:"message_content"` // 留言内容
}

type RemarkQueryReq struct {
	PageQuery
}

type ReplyMsg struct {
	Type      int    `json:"type"`      // 类型
	Data      string `json:"data"`      // 数据
	Timestamp int64  `json:"timestamp"` //时间戳
}

type ResetPasswordReq struct {
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"` // 确认密码
	Email           string `json:"email"`
	VerifyCode      string `json:"verify_code"` // 验证码
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	TraceId string      `json:"trace_id"`
}

type RestHeader struct {
	HeaderCountry    string `header:"Country,optional"`
	HeaderLanguage   string `header:"Language,optional"`
	HeaderTimezone   string `header:"Timezone,optional"`
	HeaderAppName    string `header:"App-name,optional"`
	HeaderXUserId    string `header:"X-User-Id,optional"`
	HeaderXAuthToken string `header:"X-Auth-Token,optional"`
	HeaderTerminalId string `header:"X-Terminal-Id,optional"`
}

type SendEmailVerifyCodeReq struct {
	Email string `json:"email"` // 邮箱
	Type  string `json:"type"`  // 类型 register,reset_password,bind_email,bind_phone
}

type SendMessageReq struct {
	Type    string `json:"type"`    // 消息类型 1: 文本消息 2: 图片消息 3: 文件消息 4: 语音消息 5: 视频消息
	Content string `json:"content"` // 消息内容
}

type SendPhoneVerifyCodeReq struct {
	Phone string `json:"phone"` // 手机号
	Type  string `json:"type"`  // 类型 register,reset_password,bind_email,bind_phone
}

type Tag struct {
	Id           int64  `json:"id"`            // 标签ID
	TagName      string `json:"tag_name"`      // 标签名
	ArticleCount int64  `json:"article_count"` // 文章数量
	CreatedAt    int64  `json:"created_at"`    // 创建时间
	UpdatedAt    int64  `json:"updated_at"`    // 更新时间
}

type TagQueryReq struct {
	PageQuery
	TagName string `json:"tag_name,optional"` // 标签名
}

type Talk struct {
	Id           int64       `json:"id"`            // 说说ID
	UserId       string      `json:"user_id"`       // 用户ID
	Content      string      `json:"content"`       // 评论内容
	ImgList      []string    `json:"img_list"`      // 图片URL列表
	IsTop        int64       `json:"is_top"`        // 是否置顶
	Status       int64       `json:"status"`        // 状态 1.公开 2.私密
	LikeCount    int64       `json:"like_count"`    // 点赞量
	CommentCount int64       `json:"comment_count"` // 评论量
	CreatedAt    int64       `json:"created_at"`    // 创建时间
	UpdatedAt    int64       `json:"updated_at"`    // 更新时间
	User         *UserInfoVO `json:"user"`          // 用户信息
}

type TalkQueryReq struct {
	PageQuery
}

type ThirdLoginReq struct {
	Platform string `json:"platform"`      // 平台
	Code     string `json:"code,optional"` // 授权码
}

type Token struct {
	UserId           string `json:"user_id"`            // 用户id
	TokenType        string `json:"token_type"`         // token类型,Bearer
	AccessToken      string `json:"access_token"`       // 访问token,过期时间较短。2h
	ExpiresIn        int64  `json:"expires_in"`         // 访问token过期时间
	RefreshToken     string `json:"refresh_token"`      // 刷新token,过期时间较长。30d
	RefreshExpiresIn int64  `json:"refresh_expires_in"` // 刷新token过期时间
	Scope            string `json:"scope"`              // 作用域
}

type UpdateCommentReq struct {
	Id             int64  `json:"id"`                     // 主键
	ReplyUserId    string `json:"reply_user_id,optional"` // 回复用户id
	CommentContent string `json:"comment_content"`        // 评论内容
	Status         int64  `json:"status,optional"`        // 状态 0.正常 1.已编辑 2.已删除
}

type UpdateUserAvatarReq struct {
	Avatar string `json:"avatar"` // 头像
}

type UpdateUserBindEmailReq struct {
	Email      string `json:"email"`       // 邮箱
	VerifyCode string `json:"verify_code"` // 验证码
}

type UpdateUserBindPhoneReq struct {
	Phone      string `json:"phone"`       // 手机号
	VerifyCode string `json:"verify_code"` // 验证码
}

type UpdateUserBindThirdPartyReq struct {
	Platform string `json:"platform"`       // 平台
	Code     string `json:"code"`           // 授权码
	State    string `json:"state,optional"` // 状态
}

type UpdateUserInfoReq struct {
	Nickname string `json:"nickname"` // 昵称
	UserInfoExt
}

type UpdateUserPasswordReq struct {
	OldPassword     string `json:"old_password"`     // 旧密码
	NewPassword     string `json:"new_password"`     // 新密码
	ConfirmPassword string `json:"confirm_password"` // 确认密码
}

type UploadFileReq struct {
	File     interface{} `form:"file,optional"`      // 文件
	FilePath string      `form:"file_path,optional"` // 文件路径
}

type UserInfoExt struct {
	Gender  int64  `json:"gender,optional"`  // 性别 0未知 1男 2女
	Intro   string `json:"intro,optional"`   // 简介
	Website string `json:"website,optional"` // 网站
}

type UserInfoResp struct {
	UserId    string `json:"user_id"`    // 用户id
	Username  string `json:"username"`   // 用户名
	Nickname  string `json:"nickname"`   // 用户昵称
	Avatar    string `json:"avatar"`     // 用户头像
	Email     string `json:"email"`      // 用户邮箱
	Phone     string `json:"phone"`      // 用户手机号
	CreatedAt int64  `json:"created_at"` // 创建时间
	UserInfoExt
	ThirdParty []*UserThirdPartyInfo `json:"third_party"`
}

type UserInfoVO struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	UserInfoExt
}

type UserLikeResp struct {
	ArticleLikeSet []int64 `json:"article_like_set"`
	CommentLikeSet []int64 `json:"comment_like_set"`
	TalkLikeSet    []int64 `json:"talk_like_set"`
}

type UserThirdPartyInfo struct {
	Platform  string `json:"platform"`   // 平台
	OpenId    string `json:"open_id"`    // 平台用户id
	Nickname  string `json:"nickname"`   // 昵称
	Avatar    string `json:"avatar"`     // 头像
	CreatedAt int64  `json:"created_at"` // 创建时间
}

type WebsiteConfigVO struct {
	AdminUrl          string   `json:"admin_url"`           // 后台地址
	AlipayQrCode      string   `json:"alipay_qr_code"`      // 支付宝二维码
	Gitee             string   `json:"gitee"`               // Gitee
	Github            string   `json:"github"`              // Github
	IsChatRoom        int64    `json:"is_chat_room"`        // 是否开启聊天室
	IsCommentReview   int64    `json:"is_comment_review"`   // 是否开启评论审核
	IsEmailNotice     int64    `json:"is_email_notice"`     // 是否开启邮件通知
	IsMessageReview   int64    `json:"is_message_review"`   // 是否开启留言审核
	IsMusicPlayer     int64    `json:"is_music_player"`     // 是否开启音乐播放器
	IsReward          int64    `json:"is_reward"`           // 是否开启打赏
	Qq                string   `json:"qq"`                  // QQ
	SocialLoginList   []string `json:"social_login_list"`   // 社交登录列表
	SocialUrlList     []string `json:"social_url_list"`     // 社交地址列表
	TouristAvatar     string   `json:"tourist_avatar"`      // 游客头像
	UserAvatar        string   `json:"user_avatar"`         // 用户头像
	WebsiteAuthor     string   `json:"website_author"`      // 网站作者
	WebsiteAvatar     string   `json:"website_avatar"`      // 网站头像
	WebsiteCreateTime string   `json:"website_create_time"` // 网站创建时间
	WebsiteIntro      string   `json:"website_intro"`       // 网站介绍
	WebsiteName       string   `json:"website_name"`        // 网站名称
	WebsiteNotice     string   `json:"website_notice"`      // 网站公告
	WebsiteRecordNo   string   `json:"website_record_no"`   // 网站备案号
	WebsocketUrl      string   `json:"websocket_url"`       // websocket地址
	WeixinQrCode      string   `json:"weixin_qr_code"`      // 微信二维码
}
