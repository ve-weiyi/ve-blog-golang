package logic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/constant"
)

// 查看评论id集合下的回复评论
func (s *CommentRepository) FindCommentReplyList(id int, page *request.PageInfo) (out []*entity.Comment, total int64, err error) {

	page.Conditions = append(page.Conditions, &request.Condition{
		Flag:  "AND",
		Field: "parent_id",
		Rule:  "=",
		Value: id,
	})

	return s.FindCommentList(nil, page)
}

// 点赞评论
func (s *CommentRepository) LikeComment(ctx context.Context, uid int, commentId int) (data interface{}, err error) {
	// 用户点赞的评论列表
	commentUserLikeKey := constant.RedisWrapKey("comment_user_like", uid)
	// 当前评论的点赞量
	commentLikeCountKey := constant.RedisWrapKey("comment_like_count", commentId)

	// 判断是否已经点赞
	if s.Cache.SIsMember(ctx, commentUserLikeKey, commentId).Val() {
		// 点过赞则删除评论id
		s.Cache.SRem(ctx, commentUserLikeKey, commentId)
		// 评论点赞量-1
		s.Cache.Decr(ctx, commentLikeCountKey)
	} else {
		// 未点赞则增加评论id
		s.Cache.SAdd(ctx, commentUserLikeKey, commentId)
		// 评论点赞量+1
		s.Cache.Incr(ctx, commentLikeCountKey)
	}

	return data, nil
}

// 获取用户点赞记录
func (s *CommentRepository) FindUserLikeComment(ctx context.Context, uid int) (data []string, err error) {
	// 用户点赞的评论列表
	commentUserLikeKey := constant.RedisWrapKey("comment_user_like", uid)
	return global.REDIS.SMembers(ctx, commentUserLikeKey).Result()
}
