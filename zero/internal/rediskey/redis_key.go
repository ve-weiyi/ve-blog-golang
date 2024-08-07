package rediskey

import (
	"fmt"
)

func GetUserLogoutKey(uid string) string {
	return fmt.Sprintf("user:logout:%s", uid)
}

func GetUserLikeArticleKey(uid string) string {
	return fmt.Sprintf("user:like:article:%s", uid)
}

func GetUserLikeCommentKey(uid string) string {
	return fmt.Sprintf("user:like:comment:%s", uid)
}

func GetUserLikeTalkKey(uid string) string {
	return fmt.Sprintf("user:like:talk:%s", uid)
}

func GetArticleLikeCountKey(cid string) string {
	return fmt.Sprintf("article:like:%v", cid)
}

func GetCommentLikeCountKey(cid string) string {
	return fmt.Sprintf("comment:like:%v", cid)
}

func GetTalkLikeCountKey(cid string) string {
	return fmt.Sprintf("talk:like:%v", cid)
}

func GetTotalVisitCountKey() string {
	return "visit:count"
}
