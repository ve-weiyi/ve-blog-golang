package response

type BlogHomeInfoDTO struct {
	ArticleCount  int64            `json:"articleCount"`  // 文章数量
	CategoryCount int64            `json:"categoryCount"` // 分类数量
	TagCount      int64            `json:"tagCount"`      // 标签数量
	ViewsCount    string           `json:"viewsCount"`    // 访问量
	WebsiteConfig *WebsiteConfigVO `json:"websiteConfig"` // 网站配置
	PageList      []*PageVO        `json:"pageList"`      // 页面列表
}

type WebsiteConfigVO struct {
	// 网站配置字段...
}

type PageVO struct {
	// 页面字段...
}

type BlogBackInfoDTO struct {
	ViewsCount            int64                   `json:"viewsCount"`            // 访问量
	MessageCount          int64                   `json:"messageCount"`          // 留言量
	UserCount             int64                   `json:"userCount"`             // 用户量
	ArticleCount          int64                   `json:"articleCount"`          // 文章量
	CategoryDTOList       []*CategoryDTO          `json:"categoryDTOList"`       // 分类统计
	TagDTOList            []*TagDTO               `json:"tagDTOList"`            // 标签列表
	ArticleStatisticsList []*ArticleStatisticsDTO `json:"articleStatisticsList"` // 文章统计列表
	UniqueViewDTOList     []*UniqueViewDTO        `json:"uniqueViewDTOList"`     // 一周用户量集合
	ArticleRankDTOList    []*ArticleRankDTO       `json:"articleRankDTOList"`    // 文章浏览量排行
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
	ID           int    `json:"id"`           // 文章ID
	ArticleTitle string `json:"articleTitle"` // 文章标题
	Count        int64  `json:"count"`        // 数量
}
