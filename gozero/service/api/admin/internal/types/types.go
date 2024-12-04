// Code generated by goctl. DO NOT EDIT.
package types

type AboutMe struct {
	Content string `json:"content"`
}

type AccountArea struct {
	Name  string `json:"name"`
	Value int64  `json:"value"`
}

type AccountLoginHistory struct {
	Id        int64  `json:"id,optional"`
	Username  string `json:"username"`   // 用户名
	Nickname  string `json:"nickname"`   // 用户昵称
	Avatar    string `json:"avatar"`     // 用户头像
	LoginType string `json:"login_type"` // 登录类型
	Agent     string `json:"agent"`      // 代理
	IpAddress string `json:"ip_address"` // ip host
	IpSource  string `json:"ip_source"`  // ip 源
	LoginAt   int64  `json:"login_at"`   // 登录时间
	LogoutAt  int64  `json:"logout_out"` // 登出时间
}

type AccountQuery struct {
	PageQuery
	Username string `json:"username,optional"`
	Nickname string `json:"nickname,optional"`
}

type AdminHomeInfo struct {
	ViewCount         int64                   `json:"view_count"`         // 访问量
	UserCount         int64                   `json:"user_count"`         // 用户量
	ArticleCount      int64                   `json:"article_count"`      // 文章量
	RemarkCount       int64                   `json:"remark_count"`       // 留言量
	CategoryList      []*CategoryDTO          `json:"category_list"`      // 分类列表
	TagList           []*TagDTO               `json:"tag_list"`           // 标签列表
	ArticleViewRanks  []*ArticleViewDTO       `json:"article_view_ranks"` // 文章浏览量排行
	ArticleStatistics []*ArticleStatisticsDTO `json:"article_statistics"` // 文章提交统计
	UserVisitDaliy    []*UserVisitDTO         `json:"user_visit_daliy"`   // 用户每日访问量
}

type AlbumBackDTO struct {
	Id         int64  `json:"id,optional"` // 主键
	AlbumName  string `json:"album_name"`  // 相册名
	AlbumDesc  string `json:"album_desc"`  // 相册描述
	AlbumCover string `json:"album_cover"` // 相册封面
	IsDelete   int64  `json:"is_delete"`   // 是否删除
	Status     int64  `json:"status"`      // 状态值 1公开 2私密
	CreatedAt  int64  `json:"created_at"`  // 创建时间
	UpdatedAt  int64  `json:"updated_at"`  // 更新时间
	PhotoCount int64  `json:"photo_count"` // 照片数量
}

type AlbumNewReq struct {
	Id         int64  `json:"id,optional"` // 主键
	AlbumName  string `json:"album_name"`  // 相册名
	AlbumDesc  string `json:"album_desc"`  // 相册描述
	AlbumCover string `json:"album_cover"` // 相册封面
	IsDelete   int64  `json:"is_delete"`   // 是否删除
	Status     int64  `json:"status"`      // 状态值 1公开 2私密
}

type AlbumQuery struct {
	PageQuery
	AlbumName string `json:"album_name,optional"` // 相册名
}

type ApiBackDTO struct {
	Id        int64         `json:"id,optional"`         // 主键id
	ParentId  int64         `json:"parent_id"`           // 分组id
	Name      string        `json:"name"`                // api名称
	Path      string        `json:"path"`                // api路径
	Method    string        `json:"method"`              // api请求方法
	Traceable int64         `json:"traceable"`           // 是否追溯操作记录 0需要，1是
	IsDisable int64         `json:"is_disable,optional"` // 是否禁用 0否 1是
	CreatedAt int64         `json:"created_at"`          // 创建时间
	UpdatedAt int64         `json:"updated_at"`          // 更新时间
	Children  []*ApiBackDTO `json:"children"`
}

type ApiNewReq struct {
	Id        int64  `json:"id,optional"`         // 主键id
	ParentId  int64  `json:"parent_id"`           // 分组id
	Name      string `json:"name"`                // api名称
	Path      string `json:"path"`                // api路径
	Method    string `json:"method"`              // api请求方法
	Traceable int64  `json:"traceable"`           // 是否追溯操作记录 0需要，1是
	IsDisable int64  `json:"is_disable,optional"` // 是否禁用 0否 1是
}

type ApiQuery struct {
	PageQuery
	Name   string `json:"name,optional"`   // api名称
	Path   string `json:"path,optional"`   // api路径
	Method string `json:"method,optional"` // api请求方法
}

type ArticleBackDTO struct {
	Id             int64    `json:"id,optional"`     // 文章ID
	ArticleCover   string   `json:"article_cover"`   // 文章缩略图
	ArticleTitle   string   `json:"article_title"`   // 标题
	ArticleContent string   `json:"article_content"` // 内容
	ArticleType    int64    `json:"article_type"`    // 文章类型 1原创 2转载 3翻译
	OriginalUrl    string   `json:"original_url"`    // 原文链接
	IsTop          int64    `json:"is_top"`          // 是否置顶
	IsDelete       int64    `json:"is_delete"`       // 是否删除
	Status         int64    `json:"status"`          // 状态值 1 公开 2 私密 3 草稿 4 已删除
	CreatedAt      int64    `json:"created_at"`      // 发表时间
	UpdatedAt      int64    `json:"updated_at"`      // 更新时间
	CategoryName   string   `json:"category_name"`   // 文章分类名
	TagNameList    []string `json:"tag_name_list"`   // 文章标签列表
	LikeCount      int64    `json:"like_count"`      // 点赞量
	ViewsCount     int64    `json:"views_count"`     // 浏览量
}

type ArticleNewReq struct {
	Id             int64    `json:"id,optional"`            // id
	ArticleCover   string   `json:"article_cover"`          // 文章缩略图
	ArticleTitle   string   `json:"article_title"`          // 标题
	ArticleContent string   `json:"article_content"`        // 内容
	ArticleType    int64    `json:"article_type"`           // 文章类型 1原创 2转载 3翻译
	OriginalUrl    string   `json:"original_url"`           // 原文链接
	Status         int64    `json:"status"`                 // 状态值 1 公开 2 私密 3 草稿 4 已删除
	CategoryName   string   `json:"category_name,optional"` // 文章分类名
	TagNameList    []string `json:"tag_name_list,optional"` // 文章标签列表
}

type ArticleQuery struct {
	PageQuery
	ArticleTitle string `json:"article_title,optional"` // 标题
	ArticleType  int64  `json:"article_type,optional"`  // 文章类型 1原创 2转载 3翻译
	IsTop        int64  `json:"is_top,optional"`        // 是否置顶
	IsDelete     int64  `json:"is_delete,optional"`     // 是否删除
	Status       int64  `json:"status,optional"`        // 状态值 1 公开 2 私密 3 草稿 4 已删除
	CategoryName string `json:"category_name,optional"` // 文章分类名
	TagName      string `json:"tag_name,optional"`
}

type ArticleRecycleReq struct {
	Id       int64 `json:"id,optional"` // 文章ID
	IsDelete int64 `json:"is_delete"`   // 是否删除
}

type ArticleStatisticsDTO struct {
	Date  string `json:"date"`  // 日期
	Count int64  `json:"count"` // 数量
}

type ArticleTopReq struct {
	Id    int64 `json:"id,optional"` // 文章ID
	IsTop int64 `json:"is_top"`      // 是否置顶
}

type ArticleViewDTO struct {
	Id           int64  `json:"id,optional"`   // 文章ID
	ArticleTitle string `json:"article_title"` // 文章标题
	ViewCount    int64  `json:"view_count"`    // 浏览量
}

type BatchResp struct {
	SuccessCount int64 `json:"success_count"`
}

type CategoryBackDTO struct {
	Id           int64  `json:"id,optional"`
	CategoryName string `json:"category_name"` // 分类名
	ArticleCount int64  `json:"article_count"`
	CreatedAt    int64  `json:"created_at"` // 创建时间
	UpdatedAt    int64  `json:"updated_at"` // 更新时间
}

type CategoryDTO struct {
	Id           int64  `json:"id,optional"`
	CategoryName string `json:"category_name"` // 分类名
	ArticleCount int64  `json:"article_count"` // 文章数量
}

type CategoryNewReq struct {
	Id           int64  `json:"id,optional"`
	CategoryName string `json:"category_name"` // 分类名
}

type CategoryQuery struct {
	PageQuery
	CategoryName string `json:"category_name,optional"` // 分类名
}

type CommentBackDTO struct {
	Id             int64  `json:"id"`              // 评论ID
	Type           int64  `json:"type"`            // 评论类型 1.文章 2.友链 3.说说
	TopicTitle     string `json:"topic_title"`     // 评论主题
	Avatar         string `json:"avatar"`          // 用户头像
	Nickname       string `json:"nickname"`        // 用户昵称
	ToNickname     string `json:"to_nickname"`     // 被回复人昵称
	CommentContent string `json:"comment_content"` // 评论内容
	IsReview       int64  `json:"is_review"`       // 是否审核 0.未审核 1.已审核
	CreatedAt      int64  `json:"created_at"`      // 创建时间
}

type CommentQuery struct {
	PageQuery
	Avatar   string `json:"avatar,optional"` // 用户头像
	IsReview int64  `json:"is_review,optional"`
	Type     int64  `json:"type,optional"` // 评论类型 1.文章 2.友链 3.说说
}

type CommentReviewReq struct {
	Ids      []int64 `json:"ids,optional"`
	IsReview int64   `json:"is_review,optional"`
}

type EmptyReq struct {
}

type EmptyResp struct {
}

type FileBackDTO struct {
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

type FileFolderNewReq struct {
	FilePath string `json:"file_path"` // 文件路径
	FileName string `json:"file_name"` // 文件名称
}

type FileQuery struct {
	PageQuery
	FilePath string `json:"file_path,optional"` // 文件路径
	FileType string `json:"file_type,optional"` // 文件类型
}

type FriendBackDTO struct {
	Id          int64  `json:"id,optional"`  // id
	LinkName    string `json:"link_name"`    // 链接名
	LinkAvatar  string `json:"link_avatar"`  // 链接头像
	LinkAddress string `json:"link_address"` // 链接地址
	LinkIntro   string `json:"link_intro"`   // 链接介绍
	CreatedAt   int64  `json:"created_at"`   // 创建时间
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间
}

type FriendNewReq struct {
	Id          int64  `json:"id,optional"`  // id
	LinkName    string `json:"link_name"`    // 链接名
	LinkAvatar  string `json:"link_avatar"`  // 链接头像
	LinkAddress string `json:"link_address"` // 链接地址
	LinkIntro   string `json:"link_intro"`   // 链接介绍
}

type FriendQuery struct {
	PageQuery
	LinkName string `json:"link_name,optional"` // 链接名
}

type IdReq struct {
	Id int64 `json:"id"`
}

type IdsReq struct {
	Ids []int64 `json:"ids"`
}

type ListUploadFileReq struct {
	FilePath string `json:"file_path,optional"` // 文件路径
	Limit    int64  `json:"limit,optional"`     // 限制
}

type ListUploadFileResp struct {
	Urls []string `json:"urls"` // 文件路径
}

type LoginReq struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	VerifyCode string `json:"verify_code"` // 验证码
}

type LoginResp struct {
	Token *Token `json:"token"`
}

type MenuBackDTO struct {
	Id        int64  `json:"id,optional"`        // 主键
	ParentId  int64  `json:"parent_id,optional"` // 父id
	Path      string `json:"path,optional"`      // 路由地址
	Name      string `json:"name,optional"`      // 路由名字
	Component string `json:"component,optional"` // Layout组件
	Redirect  string `json:"redirect,optional"`  // 路由重定向
	MenuMeta
	Children  []*MenuBackDTO `json:"children,optional"`
	CreatedAt int64          `json:"created_at"` // 创建时间
	UpdatedAt int64          `json:"updated_at"` // 更新时间
}

type MenuMeta struct {
	Type       int64             `json:"type,optional"`        // 菜单类型（0代表目录、1代表菜单、2代表按钮、3代表外链）
	Title      string            `json:"title,optional"`       // 菜单标题
	Icon       string            `json:"icon,optional"`        // 菜单图标
	Rank       int64             `json:"rank,optional"`        // 排序
	Perm       string            `json:"perm,optional"`        // 权限标识
	Params     []*MenuMetaParams `json:"params,optional"`      // 参数
	KeepAlive  int64             `json:"keep_alive,optional"`  // 是否缓存
	AlwaysShow int64             `json:"always_show,optional"` // 是否一直显示菜单
	IsHidden   int64             `json:"is_hidden,optional"`   // 是否隐藏
	IsDisable  int64             `json:"is_disable,optional"`  // 是否禁用
}

type MenuMetaParams struct {
	Key   string `json:"key,optional"`
	Value string `json:"value,optional"`
}

type MenuNewReq struct {
	Id        int64  `json:"id,optional"`        // 主键
	ParentId  int64  `json:"parent_id,optional"` // 父id
	Path      string `json:"path,optional"`      // 路由地址
	Name      string `json:"name,optional"`      // 路由名字
	Component string `json:"component,optional"` // Layout组件
	Redirect  string `json:"redirect,optional"`  // 路由重定向
	MenuMeta
	Children []*MenuNewReq `json:"children,optional"`
}

type MenuQuery struct {
	PageQuery
	Name  string `json:"name,optional"`  // 路由名字
	Title string `json:"title,optional"` // 菜单标题
}

type MultiUploadFileReq struct {
	Files    []interface{} `form:"files,optional"`     // 文件列表
	FilePath string        `form:"file_path,optional"` // 文件路径
}

type OperationLogBackDTO struct {
	Id             int64  `json:"id,optional"`     // 主键id
	UserId         string `json:"user_id"`         // 用户id
	Nickname       string `json:"nickname"`        // 用户昵称
	IpAddress      string `json:"ip_address"`      // 操作ip
	IpSource       string `json:"ip_source"`       // 操作地址
	OptModule      string `json:"opt_module"`      // 操作模块
	OptDesc        string `json:"opt_desc"`        // 操作描述
	RequestUrl     string `json:"request_url"`     // 请求地址
	RequestMethod  string `json:"request_method"`  // 请求方式
	RequestHeader  string `json:"request_header"`  // 请求头参数
	RequestData    string `json:"request_data"`    // 请求参数
	ResponseData   string `json:"response_data"`   // 返回数据
	ResponseStatus int64  `json:"response_status"` // 响应状态码
	Cost           string `json:"cost"`            // 耗时（ms）
	CreatedAt      int64  `json:"created_at"`      // 创建时间
	UpdatedAt      int64  `json:"updated_at"`      // 更新时间
}

type OperationLogQuery struct {
	PageQuery
}

type PageBackDTO struct {
	Id             int64    `json:"id,optional"`     // 页面id
	PageName       string   `json:"page_name"`       // 页面名
	PageLabel      string   `json:"page_label"`      // 页面标签
	PageCover      string   `json:"page_cover"`      // 页面封面
	IsCarousel     int64    `json:"is_carousel"`     // 是否轮播
	CarouselCovers []string `json:"carousel_covers"` // 轮播封面
	CreatedAt      int64    `json:"created_at"`      // 创建时间
	UpdatedAt      int64    `json:"updated_at"`      // 更新时间
}

type PageNewReq struct {
	Id             int64    `json:"id,optional"`              // 页面id
	PageName       string   `json:"page_name"`                // 页面名
	PageLabel      string   `json:"page_label"`               // 页面标签
	PageCover      string   `json:"page_cover"`               // 页面封面
	IsCarousel     int64    `json:"is_carousel,optional"`     // 是否轮播
	CarouselCovers []string `json:"carousel_covers,optional"` // 轮播封面
}

type PageQuery struct {
	Page     int64    `json:"page,optional"`      // 当前页码
	PageSize int64    `json:"page_size,optional"` // 每页数量
	Sorts    []string `json:"sorts,optional"`     // 排序
}

type PageQueryReq struct {
	PageQuery
	PageName string `json:"page_name,optional"` // 页面名
}

type PageResp struct {
	Page     int64       `json:"page"`
	PageSize int64       `json:"page_size"`
	Total    int64       `json:"total"`
	List     interface{} `json:"list"`
}

type PhotoBackDTO struct {
	Id        int64  `json:"id,optional"` // 主键
	AlbumId   int64  `json:"album_id"`    // 相册id
	PhotoName string `json:"photo_name"`  // 照片名
	PhotoDesc string `json:"photo_desc"`  // 照片描述
	PhotoSrc  string `json:"photo_src"`   // 照片地址
	IsDelete  int64  `json:"is_delete"`   // 是否删除
	CreatedAt int64  `json:"created_at"`  // 创建时间
	UpdatedAt int64  `json:"updated_at"`  // 更新时间
}

type PhotoNewReq struct {
	Id        int64  `json:"id,optional"` // 主键
	AlbumId   int64  `json:"album_id"`    // 相册id
	PhotoName string `json:"photo_name"`  // 照片名
	PhotoDesc string `json:"photo_desc"`  // 照片描述
	PhotoSrc  string `json:"photo_src"`   // 照片地址
	IsDelete  int64  `json:"is_delete"`   // 是否删除
}

type PhotoQuery struct {
	PageQuery
	AlbumId int64 `json:"album_id,optional"` // 相册id
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

type RemarkBackDTO struct {
	Id             int64  `json:"id,optional"`     // 主键id
	Nickname       string `json:"nickname"`        // 昵称
	Avatar         string `json:"avatar"`          // 头像
	MessageContent string `json:"message_content"` // 留言内容
	IpAddress      string `json:"ip_address"`      // 用户ip
	IpSource       string `json:"ip_source"`       // 用户地址
	Time           int64  `json:"time"`            // 弹幕速度
	IsReview       int64  `json:"is_review"`       // 是否审核
	CreatedAt      int64  `json:"created_at"`      // 发布时间
	UpdatedAt      int64  `json:"updated_at"`      // 更新时间
}

type RemarkNewReq struct {
	Id       int64 `json:"id,optional"` // 主键id
	IsReview int64 `json:"is_review"`   // 是否审核
}

type RemarkQuery struct {
	PageQuery
	Nickname string `json:"nickname,optional"`  // 昵称
	IsReview int64  `json:"is_review,optional"` // 是否审核
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

type RoleBackDTO struct {
	Id          int64  `json:"id,optional"`  // 主键id
	ParentId    int64  `json:"parent_id"`    // 父角色id
	RoleName    string `json:"role_name"`    // 角色名
	RoleLabel   string `json:"role_label"`   // 角色标签
	RoleComment string `json:"role_comment"` // 角色备注
	IsDisable   int64  `json:"is_disable"`   // 是否禁用  0否 1是
	IsDefault   int64  `json:"is_default"`   // 是否默认角色 0否 1是
	CreatedAt   int64  `json:"created_at"`   // 创建时间
	UpdatedAt   int64  `json:"updated_at"`   // 更新时间
}

type RoleNewReq struct {
	Id          int64  `json:"id,optional"`  // 主键id
	ParentId    int64  `json:"parent_id"`    // 父角色id
	RoleName    string `json:"role_name"`    // 角色名
	RoleLabel   string `json:"role_label"`   // 角色标签
	RoleComment string `json:"role_comment"` // 角色备注
	IsDisable   int64  `json:"is_disable"`   // 是否禁用  0否 1是
	IsDefault   int64  `json:"is_default"`   // 是否默认角色 0否 1是
}

type RoleQuery struct {
	PageQuery
	RoleName  string `json:"role_name,optional"`  // 角色名
	RoleLabel string `json:"role_label,optional"` // 角色标签
	IsDisable int64  `json:"is_disable,optional"` // 是否禁用  0否 1是
}

type RoleResourcesResp struct {
	RoleId  int64   `json:"role_id"`
	ApiIds  []int64 `json:"api_ids"`
	MenuIds []int64 `json:"menu_ids"`
}

type Server struct {
	Os   interface{} `json:"os"`
	Cpu  interface{} `json:"cpu"`
	Ram  interface{} `json:"ram"`
	Disk interface{} `json:"disk"`
}

type SyncApiReq struct {
	ApiFilePath string `json:"api_file_path"` // api文件路径
}

type SyncMenuReq struct {
	Menus []*MenuNewReq `json:"menus"`
}

type TagBackDTO struct {
	Id           int64  `json:"id,optional"`   // 标签ID
	TagName      string `json:"tag_name"`      // 标签名
	ArticleCount int64  `json:"article_count"` // 文章数量
	CreatedAt    int64  `json:"created_at"`    // 创建时间
	UpdatedAt    int64  `json:"updated_at"`    // 更新时间
}

type TagDTO struct {
	Id           int64  `json:"id,optional"`   // 标签ID
	TagName      string `json:"tag_name"`      // 标签名
	ArticleCount int64  `json:"article_count"` // 文章数量
}

type TagNewReq struct {
	Id      int64  `json:"id,optional"`
	TagName string `json:"tag_name"` // 标签名
}

type TagQuery struct {
	PageQuery
	TagName string `json:"tag_name,optional"` // 标签名
}

type TalkBackDTO struct {
	Id           int64    `json:"id,optional"`   // 说说ID
	UserId       string   `json:"user_id"`       // 用户ID
	Nickname     string   `json:"nickname"`      // 用户昵称
	Avatar       string   `json:"avatar"`        // 用户头像
	Content      string   `json:"content"`       // 说说内容
	ImgList      []string `json:"img_list"`      // 图片URL列表
	IsTop        int64    `json:"is_top"`        // 是否置顶
	Status       int64    `json:"status"`        // 状态 1.公开 2.私密
	LikeCount    int64    `json:"like_count"`    // 点赞量
	CommentCount int64    `json:"comment_count"` // 评论量
	CreatedAt    int64    `json:"created_at"`    // 创建时间
	UpdatedAt    int64    `json:"updated_at"`    // 更新时间
}

type TalkNewReq struct {
	Id      int64    `json:"id,optional"` // 说说ID
	Content string   `json:"content"`     // 说说内容
	ImgList []string `json:"img_list"`    // 图片URL列表
	IsTop   int64    `json:"is_top"`      // 是否置顶
	Status  int64    `json:"status"`      // 状态 1.公开 2.私密
}

type TalkQuery struct {
	PageQuery
	Status int64 `json:"status,optional"` // 状态 1.公开 2.私密
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

type UpdateAccountRolesReq struct {
	UserId  string  `json:"user_id"`
	RoleIds []int64 `json:"role_ids"`
}

type UpdateAccountStatusReq struct {
	UserId string `json:"user_id"`
	Status int64  `json:"status"` // 状态: -1删除 0正常 1禁用
}

type UpdateRoleApisReq struct {
	RoleId int64   `json:"role_id"`
	ApiIds []int64 `json:"api_ids"`
}

type UpdateRoleMenusReq struct {
	RoleId  int64   `json:"role_id"`
	MenuIds []int64 `json:"menu_ids"`
}

type UploadFileReq struct {
	File     interface{} `form:"file,optional"`      // 文件
	FilePath string      `form:"file_path,optional"` // 文件路径
}

type UserApi struct {
	Id        int64      `json:"id,optional"` // 主键id
	ParentId  int64      `json:"parent_id"`   // 父id
	Name      string     `json:"name"`        // api名称
	Path      string     `json:"path"`        // api路径
	Method    string     `json:"method"`      // api请求方法
	CreatedAt int64      `json:"created_at"`  // 创建时间
	UpdatedAt int64      `json:"updated_at"`  // 更新时间
	Children  []*UserApi `json:"children"`
}

type UserApisResp struct {
	List []*UserApi `json:"list"`
}

type UserInfoExt struct {
	Intro   string `json:"intro"`   // 简介
	Website string `json:"website"` // 网站
}

type UserInfoReq struct {
	Nickname string `json:"nickname"` // 昵称
	Avatar   string `json:"avatar"`   // 头像
	UserInfoExt
}

type UserInfoResp struct {
	UserId    string           `json:"user_id"`    // 用户id
	Username  string           `json:"username"`   // 用户名
	Nickname  string           `json:"nickname"`   // 用户昵称
	Avatar    string           `json:"avatar"`     // 用户头像
	Email     string           `json:"email"`      // 用户邮箱
	Phone     string           `json:"phone"`      // 用户手机号
	Status    int64            `json:"status"`     // 状态
	LoginType string           `json:"login_type"` // 登录方式
	IpAddress string           `json:"ip_address"` // ip host
	IpSource  string           `json:"ip_source"`  // ip 源
	CreatedAt int64            `json:"created_at"`
	UpdatedAt int64            `json:"updated_at"`
	Roles     []*UserRoleLabel `json:"roles"`
	Perms     []*UserApi       `json:"perms"`
	UserInfoExt
}

type UserLoginHistory struct {
	Id        int64  `json:"id,optional"`
	LoginType string `json:"login_type"` // 登录类型
	Agent     string `json:"agent"`      // 代理
	IpAddress string `json:"ip_address"` // ip host
	IpSource  string `json:"ip_source"`  // ip 源
	LoginAt   int64  `json:"login_at"`   // 登录时间
	LogoutAt  int64  `json:"login_out"`  // 登出时间
}

type UserLoginHistoryQuery struct {
	PageQuery
}

type UserMenu struct {
	Id        int64        `json:"id,optional"` // 主键
	ParentId  int64        `json:"parent_id"`   // 父id
	Path      string       `json:"path"`        // 路由地址
	Name      string       `json:"name"`        // 路由名字
	Component string       `json:"component"`   // Layout组件
	Redirect  string       `json:"redirect"`    // 路由重定向
	Meta      UserMenuMeta `json:"meta"`        // meta配置
	Children  []*UserMenu  `json:"children,optional"`
	CreatedAt int64        `json:"created_at"` // 创建时间
	UpdatedAt int64        `json:"updated_at"` // 更新时间
}

type UserMenuMeta struct {
	Title      string `json:"title,optional"`
	Icon       string `json:"icon,optional"`
	Hidden     bool   `json:"hidden,optional"`
	AlwaysShow bool   `json:"alwaysShow,optional"`
	Affix      bool   `json:"affix,optional"`
	KeepAlive  bool   `json:"keepAlive,optional"`
	Breadcrumb bool   `json:"breadcrumb,optional"`
}

type UserMenusResp struct {
	List []*UserMenu `json:"list"`
}

type UserRole struct {
	Id          int64  `json:"id,optional"`  // 主键id
	ParentId    int64  `json:"parent_id"`    // 父id
	RoleName    string `json:"role_name"`    // 角色名
	RoleLabel   string `json:"role_label"`   // 角色标签
	RoleComment string `json:"role_comment"` // 角色备注
}

type UserRoleLabel struct {
	RoleId      int64  `json:"role_id"`
	RoleName    string `json:"role_name"`
	RoleComment string `json:"role_comment"`
}

type UserRolesResp struct {
	List []*UserRole `json:"list"`
}

type UserVisitDTO struct {
	Date  string `json:"date"`  // 日期
	Count int64  `json:"count"` // 数量
}

type WebsiteConfig struct {
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
