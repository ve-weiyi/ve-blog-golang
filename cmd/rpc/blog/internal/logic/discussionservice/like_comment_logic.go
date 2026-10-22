package discussionservicelogic

import (
	"context"

	"github.com/spf13/cast"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/discussionrpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/cachekey"
	"github.com/ve-weiyi/blog-cloud/infra/grpc/metax"
)

type LikeCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLikeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LikeCommentLogic {
	return &LikeCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 点赞评论
func (l *LikeCommentLogic) LikeComment(in *discussionrpc.LikeCommentRequest) (*discussionrpc.LikeCommentResponse, error) {
	uid, err := metax.GetUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}
	id := cast.ToString(in.Id)

	if _, err := l.svcCtx.TCommentModel.FindById(l.ctx, in.Id); err != nil {
		return nil, err
	}

	likeKey := cachekey.GetUserLikeCommentKey(uid)
	countKey := cachekey.CommentLikeCountKey

	// 点赞状态以 Redis 集合为准；读取失败时不能猜，否则会把取消点赞做成再次点赞
	liked, err := l.svcCtx.Redis.SIsMember(l.ctx, likeKey, id).Result()
	if err != nil {
		return nil, err
	}

	delta := int64(1)
	if liked {
		delta = -1
		if err := l.svcCtx.Redis.SRem(l.ctx, likeKey, id).Err(); err != nil {
			return nil, err
		}
	} else {
		if err := l.svcCtx.Redis.SAdd(l.ctx, likeKey, id).Err(); err != nil {
			return nil, err
		}
	}

	// 展示用的计数缓存失败不影响点赞结果，数据库计数是真相
	if err := l.svcCtx.Redis.ZIncrBy(l.ctx, countKey, float64(delta), id).Err(); err != nil {
		l.Errorf("LikeComment ZIncrBy error: %v", err)
	}

	// 字段级原子增减：整行更新会覆盖并发写入的其它字段
	if _, err := l.svcCtx.TCommentModel.UpdateFields(l.ctx,
		map[string]interface{}{"like_count": gorm.Expr("like_count + ?", delta)},
		"id = ?", in.Id); err != nil {
		// 数据库失败时回滚点赞状态，避免 Redis 与数据库分叉
		if liked {
			_ = l.svcCtx.Redis.SAdd(l.ctx, likeKey, id).Err()
		} else {
			_ = l.svcCtx.Redis.SRem(l.ctx, likeKey, id).Err()
		}
		_ = l.svcCtx.Redis.ZIncrBy(l.ctx, countKey, float64(-delta), id).Err()
		return nil, err
	}

	return &discussionrpc.LikeCommentResponse{Success: true}, nil
}
