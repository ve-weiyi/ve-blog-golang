package rediskey

import (
	"fmt"
)

// 命名规则：服务:模块:操作:参数

// 验证码 redis类型：string
func GetCaptchaKey(module string, username string) string {
	return fmt.Sprintf("blog:captcha:%s:%s", module, username)
}

// 用户点赞的文章集合 redis类型：set
func GetUserLikeArticleKey(uid string) string {
	return fmt.Sprintf("blog:user:like:article:%s", uid)
}

// 用户点赞的评论集合 redis类型：set
func GetUserLikeCommentKey(uid string) string {
	return fmt.Sprintf("blog:user:like:comment:%s", uid)
}

// 用户点赞的说说集合 redis类型：set
func GetUserLikeTalkKey(uid string) string {
	return fmt.Sprintf("blog:user:like:talk:%s", uid)
}

// 文章点赞数排行
func GetArticleLikeCountKey() string {
	return fmt.Sprintf("blog:article:like_count")
}

// 评论点赞数排行
func GetCommentLikeCountKey() string {
	return fmt.Sprintf("blog:comment:like_count")
}

// 说说点赞数排行
func GetTalkLikeCountKey() string {
	return fmt.Sprintf("blog:talk:like_count")
}

// 文章访问量排行
func GetArticleViewCountKey() string {
	return fmt.Sprintf("blog:article:view_count")
}

// 网站日访问量排行
func GetBlogViewCountKey() string {
	return fmt.Sprintf("blog:visit:view_count")
}

// 网站总访问量
func GetBlogTotalViewCountKey() string {
	return fmt.Sprintf("blog:visit:total_view_count")
}

// 网站每天访客集合
func GetBlogDailyVisitorKey(day string) string {
	return fmt.Sprintf("blog:visit:visitor:%v", day)
}
