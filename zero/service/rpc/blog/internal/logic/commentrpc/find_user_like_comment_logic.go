package commentrpclogic

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLikeCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLikeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLikeCommentLogic {
	return &FindUserLikeCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户点赞的评论
func (l *FindUserLikeCommentLogic) FindUserLikeComment(in *blog.UserIdReq) (*blog.FindLikeCommentResp, error) {
	uid := cast.ToString(in.UserId)
	likeKey := rediskey.GetUserLikeCommentKey(uid)

	result, err := l.svcCtx.Redis.HGetAll(l.ctx, likeKey).Result()
	if err != nil {
		return nil, err
	}

	ids := make([]int64, 0)
	for k, _ := range result {
		ids = append(ids, cast.ToInt64(k))
	}

	return &blog.FindLikeCommentResp{
		LikeCommentList: ids,
	}, nil
}
