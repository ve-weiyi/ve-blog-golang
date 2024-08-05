// Code generated by goctl. DO NOT EDIT.
package types

type AboutMe struct {
	Content string `json:"content,optional"`
}

type AdminHomeInfo struct {
	ViewsCount            int64                   `json:"views_count,optional"`             // 访问量
	MessageCount          int64                   `json:"message_count,optional"`           // 留言量
	UserCount             int64                   `json:"user_count,optional"`              // 用户量
	ArticleCount          int64                   `json:"article_count,optional"`           // 文章量
	CategoryList          []*CategoryDTO          `json:"category_list,optional"`           // 分类列表
	TagList               []*TagDTO               `json:"tag_list,optional"`                // 标签列表
	ArticleViewRankList   []*ArticleViewRankDTO   `json:"article_view_rank_list,optional"`  // 文章浏览量排行
	ArticleStatisticsList []*ArticleStatisticsDTO `json:"article_statistics_list,optional"` // 每日文章生产量
	UniqueViewList        []*UniqueViewDTO        `json:"unique_view_list,optional"`        // 每日用户访问量
}

type Api struct {
	Id        int64  `json:"id,optional"`         // 主键id
	Name      string `json:"name,optional"`       // api名称
	Path      string `json:"path,optional"`       // api路径
	Method    string `json:"method,optional"`     // api请求方法
	ParentId  int64  `json:"parent_id,optional"`  // 分组id
	Traceable int64  `json:"traceable,optional"`  // 是否追溯操作记录 0需要，1是
	Status    int64  `json:"status,optional"`     // 状态 1开，2关
	CreatedAt int64  `json:"created_at,optional"` // 创建时间
	UpdatedAt int64  `json:"updated_at,optional"` // 更新时间
}

type ApiDetails struct {
	Id        int64         `json:"id,optional"`         // 主键id
	Name      string        `json:"name,optional"`       // api名称
	Path      string        `json:"path,optional"`       // api路径
	Method    string        `json:"method,optional"`     // api请求方法
	ParentId  int64         `json:"parent_id,optional"`  // 分组id
	Traceable int64         `json:"traceable,optional"`  // 是否追溯操作记录 0需要，1是
	Status    int64         `json:"status,optional"`     // 状态 1开，2关
	CreatedAt int64         `json:"created_at,optional"` // 创建时间
	UpdatedAt int64         `json:"updated_at,optional"` // 更新时间
	Children  []*ApiDetails `json:"children,optional"`
}

type ArticleBackDTO struct {
	Id             int64    `json:"id,optional"`              // 文章ID
	ArticleCover   string   `json:"article_cover,optional"`   // 文章缩略图
	ArticleTitle   string   `json:"article_title,optional"`   // 标题
	ArticleContent string   `json:"article_content,optional"` // 内容
	ArticleType    int64    `json:"article_type,optional"`    // 文章类型 1原创 2转载 3翻译
	OriginalUrl    string   `json:"original_url,optional"`    // 原文链接
	IsTop          int64    `json:"is_top,optional"`          // 是否置顶
	IsDelete       int64    `json:"is_delete,optional"`       // 是否删除
	Status         int64    `json:"status,optional"`          // 状态值 1 公开 2 私密 3 评论可见
	CreatedAt      int64    `json:"created_at,optional"`      // 发表时间
	UpdatedAt      int64    `json:"updated_at,optional"`      // 更新时间
	CategoryName   string   `json:"category_name,optional"`   // 文章分类名
	TagNameList    []string `json:"tag_name_list,optional"`   // 文章标签列表
	LikeCount      int64    `json:"like_count,optional"`      // 点赞量
	ViewsCount     int64    `json:"views_count,optional"`     // 浏览量
}

type ArticleNewReq struct {
	Id             int64    `json:"id,optional"`              // id
	UserId         int64    `json:"user_id,optional"`         // 作者
	CategoryId     int64    `json:"category_id,optional"`     // 文章分类
	ArticleCover   string   `json:"article_cover,optional"`   // 文章缩略图
	ArticleTitle   string   `json:"article_title,optional"`   // 标题
	ArticleContent string   `json:"article_content,optional"` // 内容
	ArticleType    int64    `json:"article_type,optional"`    // 文章类型 1原创 2转载 3翻译
	OriginalUrl    string   `json:"original_url,optional"`    // 原文链接
	IsTop          int64    `json:"is_top,optional"`          // 是否置顶 0否 1是
	IsDelete       int64    `json:"is_delete,optional"`       // 是否删除  0否 1是
	Status         int64    `json:"status,optional"`          // 状态值 1公开 2私密 3评论可见
	CreatedAt      int64    `json:"created_at,optional"`      // 发表时间
	UpdatedAt      int64    `json:"updated_at,optional"`      // 更新时间
	CategoryName   string   `json:"category_name,optional"`   // 文章分类名
	TagNameList    []string `json:"tag_name_list,optional"`   // 文章标签列表
}

type ArticleRecycleReq struct {
	Id       int64 `json:"id,optional"`        // 文章ID
	IsDelete int64 `json:"is_delete,optional"` // 是否删除
}

type ArticleStatisticsDTO struct {
	Date  string `json:"date,optional"`  // 日期
	Count int64  `json:"count,optional"` // 数量
}

type ArticleTopReq struct {
	Id    int64 `json:"id,optional"`     // 文章ID
	IsTop int64 `json:"is_top,optional"` // 是否置顶
}

type ArticleViewRankDTO struct {
	Id           int64  `json:"id,optional"`            // 文章ID
	ArticleTitle string `json:"article_title,optional"` // 文章标题
	Count        int64  `json:"count,optional"`         // 数量
}

type BatchResp struct {
	SuccessCount int64 `json:"success_count"`
}

type BlogHomeInfo struct {
	ArticleCount  int64         `json:"article_count,optional"`  // 文章数量
	CategoryCount int64         `json:"category_count,optional"` // 分类数量
	TagCount      int64         `json:"tag_count,optional"`      // 标签数量
	ViewsCount    string        `json:"views_count,optional"`    // 访问量
	WebsiteConfig WebsiteConfig `json:"website_config,optional"` // 网站配置
	PageList      []*PageDTO    `json:"page_list,optional"`      // 页面列表
}

type Category struct {
	Id           int64  `json:"id,optional"`            // id
	CategoryName string `json:"category_name,optional"` // 分类名
	CreatedAt    int64  `json:"created_at,optional"`    // 创建时间
	UpdatedAt    int64  `json:"updated_at,optional"`    // 更新时间
}

type CategoryDTO struct {
	Id           int64  `json:"id,optional"`
	CategoryName string `json:"category_name,optional"` // 分类名
}

type CategoryDetails struct {
	Id           int64  `json:"id,optional"`
	CategoryName string `json:"category_name,optional"` // 分类名
	ArticleCount int64  `json:"article_count,optional"`
	CreatedAt    int64  `json:"created_at,optional"` // 创建时间
	UpdatedAt    int64  `json:"updated_at,optional"` // 更新时间
}

type CommentBackDTO struct {
	Id             int64  `json:"id,optional"`
	Type           int64  `json:"type,optional"`
	TopicTitle     string `json:"topic_title,optional"`
	Nickname       string `json:"nickname,optional"` // 用户昵称
	Avatar         string `json:"avatar,optional"`   // 用户头像
	CommentContent string `json:"comment_content,optional"`
	IsReview       int64  `json:"is_review,optional"`
	CreatedAt      int64  `json:"created_at,optional"`
}

type CommentDTO struct {
	Id               int64           `json:"id,optional"`                 // 评论id
	TopicId          int64           `json:"topic_id,optional"`           // 主题id
	ParentId         int64           `json:"parent_id,optional"`          // 父评论id
	SessionId        int64           `json:"session_id,optional"`         // 会话id
	UserId           int64           `json:"user_id,optional"`            // 用户id
	Nickname         string          `json:"nickname,optional"`           // 用户昵称
	Avatar           string          `json:"avatar,optional"`             // 用户头像
	Website          string          `json:"website,optional"`            // 用户网站
	ReplyUserId      int64           `json:"reply_user_id,optional"`      // 被回复用户id
	ReplyNickname    string          `json:"reply_nickname,optional"`     // 被回复用户昵称
	ReplyAvatar      string          `json:"reply_avatar,optional"`       // 被回复用户头像
	ReplyWebsite     string          `json:"reply_website,optional"`      // 被回复用户网站
	CommentContent   string          `json:"comment_content,optional"`    // 评论内容
	Type             int64           `json:"type,optional"`               // 评论类型 1.文章 2.友链 3.说说
	CreatedAt        int64           `json:"created_at,optional"`         // 评论时间
	LikeCount        int64           `json:"like_count,optional"`         // 点赞数
	ReplyCount       int64           `json:"reply_count,optional"`        // 回复量
	CommentReplyList []*CommentReply `json:"comment_reply_list,optional"` // 评论回复列表
}

type CommentNewReq struct {
	TopicId        int64  `json:"topic_id,optional"`        // 主题id
	ParentId       int64  `json:"parent_id,optional"`       // 父评论id
	SessionId      int64  `json:"session_id,optional"`      // 会话id
	ReplyUserId    int64  `json:"reply_user_id,optional"`   // 回复用户id
	CommentContent string `json:"comment_content,optional"` // 评论内容
	Type           int64  `json:"type,optional"`            // 评论类型 1.文章 2.友链 3.说说
}

type CommentQueryReq struct {
	Page     int64  `json:"page,optional"`      // 页码
	PageSize int64  `json:"page_size,optional"` // 每页数量
	TopicId  int64  `json:"topic_id,optional"`  // 主题id
	ParentId int64  `json:"parent_id,optional"` // 父评论id
	Type     int64  `json:"type,optional"`      // 评论类型 1.文章 2.友链 3.说说
	OrderBy  string `json:"order_by,optional"`  // 排序字段 create_at|like_count
}

type CommentReply struct {
	Id             int64  `json:"id,optional"`              // 评论id
	TopicId        int64  `json:"topic_id,optional"`        // 主题id
	ParentId       int64  `json:"parent_id,optional"`       // 父评论id
	SessionId      int64  `json:"session_id,optional"`      // 会话id
	UserId         int64  `json:"user_id,optional"`         // 用户id
	Nickname       string `json:"nickname,optional"`        // 用户昵称
	Avatar         string `json:"avatar,optional"`          // 用户头像
	Website        string `json:"website,optional"`         // 用户网站
	ReplyUserId    int64  `json:"reply_user_id,optional"`   // 被回复用户id
	ReplyNickname  string `json:"reply_nickname,optional"`  // 被回复用户昵称
	ReplyAvatar    string `json:"reply_avatar,optional"`    // 被回复用户头像
	ReplyWebsite   string `json:"reply_website,optional"`   // 被回复用户网站
	CommentContent string `json:"comment_content,optional"` // 评论内容
	Type           int64  `json:"type,optional"`            // 评论类型 1.文章 2.友链 3.说说
	CreatedAt      int64  `json:"created_at,optional"`      // 评论时间
	LikeCount      int64  `json:"like_count,optional"`      // 点赞数
}

type CommentUserInfo struct {
	Id       int64  `json:"id,optional"`
	Nickname string `json:"nickname,optional"`
	Avatar   string `json:"avatar,optional"`
	Website  string `json:"website,optional"`
}

type EmptyReq struct {
}

type EmptyResp struct {
}

type FriendLink struct {
	Id          int64  `json:"id,optional"`           // id
	LinkName    string `json:"link_name,optional"`    // 链接名
	LinkAvatar  string `json:"link_avatar,optional"`  // 链接头像
	LinkAddress string `json:"link_address,optional"` // 链接地址
	LinkIntro   string `json:"link_intro,optional"`   // 链接介绍
	CreatedAt   int64  `json:"created_at,optional"`   // 创建时间
	UpdatedAt   int64  `json:"updated_at,optional"`   // 更新时间
}

type IdReq struct {
	Id int64 `json:"id"`
}

type IdsReq struct {
	Ids []int64 `json:"ids"`
}

type LoginHistory struct {
	Id        int64  `json:"id,optional"`
	LoginType string `json:"login_type,optional"` // 登录类型
	Agent     string `json:"agent,optional"`      // 代理
	IpAddress string `json:"ip_address,optional"` // ip host
	IpSource  string `json:"ip_source,optional"`  // ip 源
	LoginTime string `json:"login_time,optional"` // 创建时间
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Code     string `json:"code,optional"`
}

type LoginResp struct {
	Token *Token `json:"token,optional"`
}

type MenuDetails struct {
	Id        int64          `json:"id,optional"`        // 主键
	ParentId  int64          `json:"parent_id,optional"` // 父id
	Title     string         `json:"title,optional"`     // 菜单标题
	Type      int64          `json:"type,optional"`      // 菜单类型（0代表菜单、1代表iframe、2代表外链、3代表按钮）
	Path      string         `json:"path,optional"`      // 路由地址
	Name      string         `json:"name,optional"`      // 路由名字
	Component string         `json:"component,optional"` // Layout组件
	Redirect  string         `json:"redirect,optional"`  // 路由重定向
	Meta      Meta           `json:"meta,optional"`      // meta配置
	Children  []*MenuDetails `json:"children,optional"`
	CreatedAt int64          `json:"created_at,optional"` // 创建时间
	UpdatedAt int64          `json:"updated_at,optional"` // 更新时间
}

type MenuPageQuery struct {
}

type Meta struct {
	Title        string     `json:"title,optional"`         // 菜单名称
	Icon         string     `json:"icon,optional"`          // 菜单图标
	ShowLink     bool       `json:"show_link,optional"`     // 是否在菜单中显示
	Rank         int64      `json:"rank,optional"`          // 菜单升序排序
	ExtraIcon    string     `json:"extra_icon,optional"`    // 菜单名称右侧的额外图标
	ShowParent   bool       `json:"show_parent,optional"`   // 是否显示父级菜单
	Roles        []string   `json:"roles,optional"`         // 页面级别权限设置
	Auths        []string   `json:"auths,optional"`         // 按钮级别权限设置
	KeepAlive    bool       `json:"keep_alive,optional"`    // 路由组件缓存
	FrameSrc     string     `json:"frame_src,optional"`     // 内嵌的iframe链接
	FrameLoading bool       `json:"frame_loading,optional"` // iframe页是否开启首次加载动画
	Transition   Transition `json:"transition,optional"`    // 页面加载动画
	HiddenTag    bool       `json:"hidden_tag,optional"`    // 是否不添加信息到标签页
	DynamicLevel int64      `json:"dynamic_level,optional"` // 动态路由可打开的最大数量
	ActivePath   string     `json:"active_path,optional"`   // 将某个菜单激活
}

type OperationLog struct {
	Id             int64  `json:"id,optional"`              // 主键id
	UserId         int64  `json:"user_id,optional"`         // 用户id
	Nickname       string `json:"nickname,optional"`        // 用户昵称
	IpAddress      string `json:"ip_address,optional"`      // 操作ip
	IpSource       string `json:"ip_source,optional"`       // 操作地址
	OptModule      string `json:"opt_module,optional"`      // 操作模块
	OptDesc        string `json:"opt_desc,optional"`        // 操作描述
	RequestUrl     string `json:"request_url,optional"`     // 请求地址
	RequestMethod  string `json:"request_method,optional"`  // 请求方式
	RequestHeader  string `json:"request_header,optional"`  // 请求头参数
	RequestData    string `json:"request_data,optional"`    // 请求参数
	ResponseData   string `json:"response_data,optional"`   // 返回数据
	ResponseStatus int64  `json:"response_status,optional"` // 响应状态码
	Cost           string `json:"cost,optional"`            // 耗时（ms）
	CreatedAt      int64  `json:"created_at,optional"`      // 创建时间
	UpdatedAt      int64  `json:"updated_at,optional"`      // 更新时间
}

type Page struct {
	Id        int64  `json:"id,optional"`         // 页面id
	PageName  string `json:"page_name,optional"`  // 页面名
	PageLabel string `json:"page_label,optional"` // 页面标签
	PageCover string `json:"page_cover,optional"` // 页面封面
	CreatedAt int64  `json:"created_at,optional"` // 创建时间
	UpdatedAt int64  `json:"updated_at,optional"` // 更新时间
}

type PageCondition struct {
	Field    string `json:"field,optional"`    // 字段
	Value    string `json:"value,optional"`    // 值
	Logic    string `json:"logic,optional"`    // and | or
	Operator string `json:"operator,optional"` // = | >= | < | in | not in |....
}

type PageDTO struct {
	Id        int64  `json:"id,optional"`         // 页面ID
	PageName  string `json:"page_name,optional"`  // 页面名称
	PageLabel string `json:"page_label,optional"` // 页面标签
	PageCover string `json:"page_cover,optional"` // 页面封面
}

type PageLimit struct {
	Page     int64 `json:"page,optional"`
	PageSize int64 `json:"page_size,optional"`
}

type PageQuery struct {
	Page       int64            `json:"page,optional"`
	PageSize   int64            `json:"page_size,optional"`
	Sorts      []*PageSort      `json:"sorts,optional"`
	Conditions []*PageCondition `json:"conditions,optional"`
}

type PageResp struct {
	Page     int64       `json:"page"`
	PageSize int64       `json:"page_size"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}

type PageSort struct {
	Field string `json:"field,optional"`
	Order string `json:"order,optional"` // asc | desc
}

type Photo struct {
	Id        int64  `json:"id,optional"`         // 主键
	AlbumId   int64  `json:"album_id,optional"`   // 相册id
	PhotoName string `json:"photo_name,optional"` // 照片名
	PhotoDesc string `json:"photo_desc,optional"` // 照片描述
	PhotoSrc  string `json:"photo_src,optional"`  // 照片地址
	IsDelete  int64  `json:"is_delete,optional"`  // 是否删除
	CreatedAt int64  `json:"created_at,optional"` // 创建时间
	UpdatedAt int64  `json:"updated_at,optional"` // 更新时间
}

type PhotoAlbum struct {
	Id         int64  `json:"id,optional"`          // 主键
	AlbumName  string `json:"album_name,optional"`  // 相册名
	AlbumDesc  string `json:"album_desc,optional"`  // 相册描述
	AlbumCover string `json:"album_cover,optional"` // 相册封面
	IsDelete   int64  `json:"is_delete,optional"`   // 是否删除
	Status     int64  `json:"status,optional"`      // 状态值 1公开 2私密
	CreatedAt  int64  `json:"created_at,optional"`  // 创建时间
	UpdatedAt  int64  `json:"updated_at,optional"`  // 更新时间
	PhotoCount int64  `json:"photo_count,optional"` // 照片数量
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

type Remark struct {
	Id             int64  `json:"id,optional"`              // 主键id
	Nickname       string `json:"nickname,optional"`        // 昵称
	Avatar         string `json:"avatar,optional"`          // 头像
	MessageContent string `json:"message_content,optional"` // 留言内容
	IpAddress      string `json:"ip_address,optional"`      // 用户ip
	IpSource       string `json:"ip_source,optional"`       // 用户地址
	Time           int64  `json:"time,optional"`            // 弹幕速度
	IsReview       int64  `json:"is_review,optional"`       // 是否审核
	CreatedAt      int64  `json:"created_at,optional"`      // 发布时间
	UpdatedAt      int64  `json:"updated_at,optional"`      // 更新时间
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

type Role struct {
	Id          int64  `json:"id,optional"`           // 主键id
	RolePid     int64  `json:"role_pid,optional"`     // 父角色id
	RoleDomain  string `json:"role_domain,optional"`  // 角色域
	RoleName    string `json:"role_name,optional"`    // 角色名
	RoleComment string `json:"role_comment,optional"` // 角色备注
	IsDisable   int64  `json:"is_disable,optional"`   // 是否禁用  0否 1是
	IsDefault   int64  `json:"is_default,optional"`   // 是否默认角色 0否 1是
	CreatedAt   int64  `json:"created_at,optional"`   // 创建时间
	UpdatedAt   int64  `json:"updated_at,optional"`   // 更新时间
}

type RoleDetails struct {
	Id          int64   `json:"id,optional"`           // 主键id
	RolePid     int64   `json:"role_pid,optional"`     // 父角色id
	RoleDomain  string  `json:"role_domain,optional"`  // 角色域
	RoleName    string  `json:"role_name,optional"`    // 角色名
	RoleComment string  `json:"role_comment,optional"` // 角色备注
	IsDisable   int64   `json:"is_disable,optional"`   // 是否禁用  0否 1是
	IsDefault   int64   `json:"is_default,optional"`   // 是否默认角色 0否 1是
	CreatedAt   int64   `json:"created_at,optional"`   // 创建时间
	UpdatedAt   int64   `json:"updated_at,optional"`   // 更新时间
	MenuIdList  []int64 `json:"menu_id_list,optional"`
	ApiIdList   []int64 `json:"resource_id_list,optional"`
}

type RoleLabel struct {
	RoleId      int64  `json:"role_id,optional"`
	RoleName    string `json:"role_name,optional"`
	RoleComment string `json:"role_comment,optional"`
}

type RoleResourcesResp struct {
	RoleId  int64   `json:"role_id,optional"`
	ApiIds  []int64 `json:"api_ids,optional"`
	MenuIds []int64 `json:"menu_ids,optional"`
}

type RouteConfigsTable struct {
	Type      int64               `json:"type,optional"`      // 菜单类型（0代表菜单、1代表iframe、2代表外链、3代表按钮）
	Path      string              `json:"path,optional"`      // 路由地址
	Name      string              `json:"name,optional"`      // 路由名字
	Component string              `json:"component,optional"` // Layout组件
	Redirect  string              `json:"redirect,optional"`  // 路由重定向
	Meta      Meta                `json:"meta,optional"`      // meta配置
	Children  []RouteConfigsTable `json:"children,optional"`  // 子路由配置项
}

type Server struct {
	Os   interface{} `json:"os"`
	Cpu  interface{} `json:"cpu"`
	Ram  interface{} `json:"ram"`
	Disk interface{} `json:"disk"`
}

type SyncMenuReq struct {
	Menus []RouteConfigsTable `json:"menus,optional"`
}

type Tag struct {
	Id        int64  `json:"id,optional"`         // id
	TagName   string `json:"tag_name,optional"`   // 标签名
	CreatedAt int64  `json:"created_at,optional"` // 创建时间
	UpdatedAt int64  `json:"updated_at,optional"` // 更新时间
}

type TagDTO struct {
	Id      int64  `json:"id,optional"`       // 标签ID
	TagName string `json:"tag_name,optional"` // 标签名
}

type TagDetails struct {
	Id           int64  `json:"id,optional"`            // 标签ID
	TagName      string `json:"tag_name,optional"`      // 标签名
	ArticleCount int64  `json:"article_count,optional"` // 文章数量
	CreatedAt    int64  `json:"created_at,optional"`    // 创建时间
	UpdatedAt    int64  `json:"updated_at,optional"`    // 更新时间
}

type TalkDetails struct {
	Id           int64    `json:"id,optional"`            // 说说ID
	UserId       int64    `json:"user_id,optional"`       // 用户ID
	Nickname     string   `json:"nickname,optional"`      // 用户昵称
	Avatar       string   `json:"avatar,optional"`        // 用户头像
	Content      string   `json:"content,optional"`       // 评论内容
	ImgList      []string `json:"img_list,optional"`      // 图片URL列表
	IsTop        int64    `json:"is_top,optional"`        // 是否置顶
	Status       int64    `json:"status,optional"`        // 状态 1.公开 2.私密
	LikeCount    int64    `json:"like_count,optional"`    // 点赞量
	CommentCount int64    `json:"comment_count,optional"` // 评论量
	CreatedAt    int64    `json:"created_at,optional"`    // 创建时间
	UpdatedAt    int64    `json:"updated_at,optional"`    // 更新时间
}

type Token struct {
	UserId           int64  `json:"user_id,optional"`            // 用户id
	TokenType        string `json:"token_type,optional"`         // token类型,Bearer
	AccessToken      string `json:"access_token,optional"`       // 访问token,过期时间较短。2h
	ExpiresIn        int64  `json:"expires_in,optional"`         // 访问token过期时间
	RefreshToken     string `json:"refresh_token,optional"`      // 刷新token,过期时间较长。30d
	RefreshExpiresIn int64  `json:"refresh_expires_in,optional"` // 刷新token过期时间
	Scope            string `json:"scope,optional"`              // 作用域
}

type Transition struct {
	Name            string `json:"name,optional"`             // 当前路由动画效果
	EnterTransition string `json:"enter_transition,optional"` // 进场动画
	LeaveTransition string `json:"leave_transition,optional"` // 离场动画
}

type UniqueViewDTO struct {
	Date  string `json:"date,optional"`  // 日期
	Count int64  `json:"count,optional"` // 数量
}

type UpdateRoleApisReq struct {
	RoleId int64   `json:"role_id,optional"`
	ApiIds []int64 `json:"api_ids,optional"`
}

type UpdateRoleMenusReq struct {
	RoleId  int64   `json:"role_id,optional"`
	MenuIds []int64 `json:"menu_ids,optional"`
}

type UpdateUserRolesReq struct {
	UserId  int64   `json:"user_id,optional"`
	RoleIds []int64 `json:"role_ids,optional"`
}

type UpdateUserStatusReq struct {
	UserId int64 `json:"user_id,optional"`
	Status int64 `json:"status,optional"` // 状态: -1删除 0正常 1禁用
}

type UploadFileReq struct {
	Label    string      `form:"label,optional"`
	File     interface{} `form:"file,optional"`      // 文件
	FileSize int64       `form:"file_size,optional"` // 文件大小
	FileMd5  string      `form:"file_md5,optional"`  // 文件md5值
}

type UploadFileResp struct {
	Id        int64  `json:"id"`         // id
	UserId    int64  `json:"user_id"`    // 用户id
	Label     string `json:"label"`      // 标签
	FileName  string `json:"file_name"`  // 文件名称
	FileSize  int64  `json:"file_size"`  // 文件大小
	FileMd5   string `json:"file_md5"`   // 文件md5值
	FileUrl   string `json:"file_url"`   // 上传路径
	CreatedAt int64  `json:"created_at"` // 创建时间
	UpdatedAt int64  `json:"updated_at"` // 更新时间
}

type User struct {
	Id        int64        `json:"id,optional"`
	Username  string       `json:"username"`
	Nickname  string       `json:"nickname,optional"`
	Avatar    string       `json:"avatar,optional"`
	Status    int64        `json:"status,optional"`
	LoginType string       `json:"login_type,optional"`
	IpAddress string       `json:"ip_address,optional"` // ip host
	IpSource  string       `json:"ip_source,optional"`  // ip 源
	CreatedAt int64        `json:"created_at,optional"`
	UpdatedAt int64        `json:"updated_at,optional"`
	LoginAt   int64        `json:"login_at,optional"`
	LogoutAt  int64        `json:"logout_at,optional"`
	Roles     []*RoleLabel `json:"roles,optional"`
	UserInfoExt
}

type UserApi struct {
	Id        int64      `json:"id,optional"`         // 主键id
	Name      string     `json:"name,optional"`       // api名称
	Path      string     `json:"path,optional"`       // api路径
	Method    string     `json:"method,optional"`     // api请求方法
	ParentId  int64      `json:"parent_id,optional"`  // 分组id
	Traceable int64      `json:"traceable,optional"`  // 是否追溯操作记录 0需要，1是
	Status    int64      `json:"status,optional"`     // 状态 1开，2关
	CreatedAt int64      `json:"created_at,optional"` // 创建时间
	UpdatedAt int64      `json:"updated_at,optional"` // 更新时间
	Children  []*UserApi `json:"children,optional"`
}

type UserApisResp struct {
	List []*UserApi `json:"list,optional"`
}

type UserArea struct {
	Name  string `json:"name,optional"`
	Value int64  `json:"value,optional"`
}

type UserInfoExt struct {
	Email   string `json:"email,optional"`   // 用户邮箱
	Phone   string `json:"phone,optional"`   // 用户手机号
	Intro   string `json:"intro,optional"`   // 简介
	Website string `json:"website,optional"` // 网站
}

type UserInfoReq struct {
	Nickname string `json:"nickname,optional"` // 昵称
	Avatar   string `json:"avatar,optional"`   // 头像
	UserInfoExt
}

type UserInfoResp struct {
	UserId   int64  `json:"user_id,optional"`  // 用户id
	Username string `json:"username"`          // 用户名
	Nickname string `json:"nickname,optional"` // 用户昵称
	Avatar   string `json:"avatar,optional"`   // 用户头像
	UserInfoExt
}

type UserMenu struct {
	Id        int64        `json:"id,optional"`        // 主键
	ParentId  int64        `json:"parent_id,optional"` // 父id
	Title     string       `json:"title,optional"`     // 菜单标题
	Type      int64        `json:"type,optional"`      // 菜单类型（0代表菜单、1代表iframe、2代表外链、3代表按钮）
	Path      string       `json:"path,optional"`      // 路由地址
	Name      string       `json:"name,optional"`      // 路由名字
	Component string       `json:"component,optional"` // Layout组件
	Redirect  string       `json:"redirect,optional"`  // 路由重定向
	Meta      UserMenuMeta `json:"meta,optional"`      // meta配置
	Children  []*UserMenu  `json:"children,optional"`
}

type UserMenuMeta struct {
	Title        string      `json:"title,optional"`         // 菜单名称
	Icon         string      `json:"icon,optional"`          // 菜单图标
	ShowLink     bool        `json:"show_link,optional"`     // 是否在菜单中显示
	Rank         int64       `json:"rank,optional"`          // 菜单升序排序
	ExtraIcon    string      `json:"extra_icon,optional"`    // 菜单名称右侧的额外图标
	ShowParent   bool        `json:"show_parent,optional"`   // 是否显示父级菜单
	Roles        []string    `json:"roles,optional"`         // 页面级别权限设置
	Auths        []string    `json:"auths,optional"`         // 按钮级别权限设置
	KeepAlive    bool        `json:"keep_alive,optional"`    // 路由组件缓存
	FrameSrc     string      `json:"frame_src,optional"`     // 内嵌的iframe链接
	FrameLoading bool        `json:"frame_loading,optional"` // iframe页是否开启首次加载动画
	Transition   interface{} `json:"transition,optional"`    // 页面加载动画
	HiddenTag    bool        `json:"hidden_tag,optional"`    // 是否不添加信息到标签页
	DynamicLevel int64       `json:"dynamic_level,optional"` // 动态路由可打开的最大数量
	ActivePath   string      `json:"active_path,optional"`   // 将某个菜单激活
}

type UserMenusResp struct {
	List []*UserMenu `json:"list,optional"`
}

type UserRole struct {
	Id             int64   `json:"id,optional"`           // 主键id
	RolePId        int64   `json:"role_pid,optional"`     // 父角色id
	RoleDomain     string  `json:"role_domain,optional"`  // 角色域
	RoleName       string  `json:"role_name,optional"`    // 角色名
	RoleComment    string  `json:"role_comment,optional"` // 角色备注
	IsDisable      int64   `json:"is_disable,optional"`   // 是否禁用  0否 1是
	IsDefault      int64   `json:"is_default,optional"`   // 是否默认角色 0否 1是
	CreatedAt      int64   `json:"created_at,optional"`   // 创建时间
	UpdatedAt      int64   `json:"updated_at,optional"`   // 更新时间
	MenuIdList     []int64 `json:"menu_id_list,optional"`
	ResourceIdList []int64 `json:"resource_id_list,optional"`
}

type UserRolesResp struct {
	List []*UserRole `json:"list,optional"`
}

type WebsiteConfig struct {
	AdminUrl          string   `json:"admin_url,optional"`           // 后台地址
	AlipayQrCode      string   `json:"alipay_qr_code,optional"`      // 支付宝二维码
	Gitee             string   `json:"gitee,optional"`               // Gitee
	Github            string   `json:"github,optional"`              // Github
	IsChatRoom        int64    `json:"is_chat_room,optional"`        // 是否开启聊天室
	IsCommentReview   int64    `json:"is_comment_review,optional"`   // 是否开启评论审核
	IsEmailNotice     int64    `json:"is_email_notice,optional"`     // 是否开启邮件通知
	IsMessageReview   int64    `json:"is_message_review,optional"`   // 是否开启留言审核
	IsMusicPlayer     int64    `json:"is_music_player,optional"`     // 是否开启音乐播放器
	IsReward          int64    `json:"is_reward,optional"`           // 是否开启打赏
	Qq                string   `json:"qq,optional"`                  // QQ
	SocialLoginList   []string `json:"social_login_list,optional"`   // 社交登录列表
	SocialUrlList     []string `json:"social_url_list,optional"`     // 社交地址列表
	TouristAvatar     string   `json:"tourist_avatar,optional"`      // 游客头像
	UserAvatar        string   `json:"user_avatar,optional"`         // 用户头像
	WebsiteAuthor     string   `json:"website_author,optional"`      // 网站作者
	WebsiteAvatar     string   `json:"website_avatar,optional"`      // 网站头像
	WebsiteCreateTime string   `json:"website_create_time,optional"` // 网站创建时间
	WebsiteIntro      string   `json:"website_intro,optional"`       // 网站介绍
	WebsiteName       string   `json:"website_name,optional"`        // 网站名称
	WebsiteNotice     string   `json:"website_notice,optional"`      // 网站公告
	WebsiteRecordNo   string   `json:"website_record_no,optional"`   // 网站备案号
	WebsocketUrl      string   `json:"websocket_url,optional"`       // websocket地址
	WeixinQrCode      string   `json:"weixin_qr_code,optional"`      // 微信二维码
}
