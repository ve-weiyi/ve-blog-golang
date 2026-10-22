package enums

// 相册相关枚举 (t_album / t_photo)
const (
	// 相册删除状态 (is_delete)
	AlbumIsDeleteNo  = 0 // 未删除
	AlbumIsDeleteYes = 1 // 已删除

	// 照片删除状态 (is_delete)
	PhotoIsDeleteNo  = 0 // 未删除
	PhotoIsDeleteYes = 1 // 已删除

	// 相册状态 (status)
	AlbumStatusPublic  = 1 // 公开
	AlbumStatusPrivate = 2 // 私密
)

// 文章相关枚举 (t_article)
const (
	// 文章类型 (article_type)
	ArticleTypeOriginal  = 1 // 原创
	ArticleTypeReprint   = 2 // 转载
	ArticleTypeTranslate = 3 // 翻译

	// 文章置顶状态 (is_top)
	ArticleIsTopALL = -1
	ArticleIsTopNo  = 0 // 未置顶
	ArticleIsTopYes = 1 // 已置顶

	// 文章删除状态 (is_delete)
	ArticleIsDeleteALL = -1
	ArticleIsDeleteNo  = 0 // 未删除
	ArticleIsDeleteYes = 1 // 已删除

	// 文章状态 (status)
	ArticleStatusPublic         = 1 // 公开
	ArticleStatusPrivate        = 2 // 私密
	ArticleStatusDraft          = 3 // 草稿
	ArticleStatusCommentVisible = 4 // 评论可见
)

// IsValidArticleStatus 判断文章状态是否合法。
func IsValidArticleStatus(v int64) bool {
	switch v {
	case ArticleStatusPublic, ArticleStatusPrivate, ArticleStatusDraft, ArticleStatusCommentVisible:
		return true
	default:
		return false
	}
}

// IsValidAlbumStatus 判断相册状态是否合法。
func IsValidAlbumStatus(v int64) bool {
	return v == AlbumStatusPublic || v == AlbumStatusPrivate
}

// 页面相关枚举 (t_page)
const (
	// 页面轮播状态 (is_carousel)
	PageIsCarouselNo  = 0 // 非轮播
	PageIsCarouselYes = 1 // 轮播
)
