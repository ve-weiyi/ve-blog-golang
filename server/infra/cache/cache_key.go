package cache

import "fmt"

const (
	ArticleUserLike  = "article_user_like"
	ArticleLikeCount = "article_like_count"
)

const (
	CommentUserLike  = "comment_user_like"
	CommentLikeCount = "comment_like_count"
)

const (
	TalkUserLike  = "talk_user_like"
	TalkLikeCount = "talk_like_count"
)

func WrapCacheKey(prefix string, key interface{}) string {
	return fmt.Sprintf("%s:%v", prefix, key)
}
