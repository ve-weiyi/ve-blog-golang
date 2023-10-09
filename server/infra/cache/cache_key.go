package cache

import "fmt"

const (
	CommentUserLike  = "comment_user_like"
	CommentLikeCount = "comment_like_count"
)

func WrapCacheKey(prefix string, key interface{}) string {
	return fmt.Sprintf("%s:%v", prefix, key)
}
