package rediskey

import (
	"fmt"
)

func GetCaptchaKey(module string, username string) string {
	return fmt.Sprintf("blog:captcha:%s:%s", module, username)
}

// 用户点赞的文章
func GetUserLikeArticleKey(uid string) string {
	return fmt.Sprintf("blog:user:like:article:%s", uid)
}

// 用户点赞的评论
func GetUserLikeCommentKey(uid string) string {
	return fmt.Sprintf("blog:user:like:comment:%s", uid)
}

// 用户点赞的说说
func GetUserLikeTalkKey(uid string) string {
	return fmt.Sprintf("blog:user:like:talk:%s", uid)
}

// 文章点赞数
func GetArticleLikeCountKey(id string) string {
	return fmt.Sprintf("blog:article:like:%v", id)
}

// 评论点赞数
func GetCommentLikeCountKey(id string) string {
	return fmt.Sprintf("blog:comment:like:%v", id)
}

// 说说点赞数
func GetTalkLikeCountKey(id string) string {
	return fmt.Sprintf("blog:talk:like:%v", id)
}

// 网站总访问量
func GetBlogViewCountKey() string {
	return fmt.Sprintf("blog:view:count")
}

// 网站访客
func GetBlogVisitorKey(day string) string {
	return fmt.Sprintf("blog:view:visitor:%v", day)
}

// 文章访问量
func GetArticleVisitCountKey(id string) string {
	return fmt.Sprintf("blog:article:visit:%v", id)
}
