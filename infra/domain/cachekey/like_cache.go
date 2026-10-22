package cachekey

import "fmt"

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

const (
	// 文章点赞数排行
	ArticleLikeCountKey = "blog:article:like_count"
	// 评论点赞数排行
	CommentLikeCountKey = "blog:comment:like_count"
	// 说说点赞数排行
	TalkLikeCountKey = "blog:talk:like_count"
)
