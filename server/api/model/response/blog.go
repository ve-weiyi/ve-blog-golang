package response

type BlogHomeInfoDTO struct {
	ArticleCount  int64            `json:"article_count"`  // 文章数量
	CategoryCount int64            `json:"category_count"` // 分类数量
	TagCount      int64            `json:"tag_count"`      // 标签数量
	ViewsCount    string           `json:"views_count"`    // 访问量
	WebsiteConfig *WebsiteConfigVO `json:"website_config"` // 网站配置
	PageList      []*PageVO        `json:"page_list"`      // 页面列表
}

type WebsiteConfigVO struct {
	// 网站配置字段...
}

type PageVO struct {
	// 页面字段...
}

type BlogBackInfoDTO struct {
	ViewsCount            int64                   `json:"views_count"`             // 访问量
	MessageCount          int64                   `json:"message_count"`           // 留言量
	UserCount             int64                   `json:"user_count"`              // 用户量
	ArticleCount          int64                   `json:"article_count"`           // 文章量
	CategoryDTOList       []*CategoryDTO          `json:"category_dto_list"`       // 分类统计
	TagDTOList            []*TagDTO               `json:"tag_dto_list"`            // 标签列表
	ArticleStatisticsList []*ArticleStatisticsDTO `json:"article_statistics_list"` // 文章统计列表
	UniqueViewDTOList     []*UniqueViewDTO        `json:"unique_view_dto_list"`    // 一周用户量集合
	ArticleRankDTOList    []*ArticleRankDTO       `json:"article_rank_dto_list"`   // 文章浏览量排行
}

// 文章统计字段
type ArticleStatisticsDTO struct {
	Day   string `json:"day"`   // 日期
	Count int64  `json:"count"` // 数量
}

// 一周用户量
type UniqueViewDTO struct {
	Day   string `json:"day"`   // 日期
	Count int64  `json:"count"` // 数量
}

// 文章浏览量排行
type ArticleRankDTO struct {
	ID           int    `json:"id"`            // 文章ID
	ArticleTitle string `json:"article_title"` // 文章标题
	Count        int64  `json:"count"`         // 数量
}
