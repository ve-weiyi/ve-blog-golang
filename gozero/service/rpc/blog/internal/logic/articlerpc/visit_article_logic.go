package articlerpclogic

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/common/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type VisitArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewVisitArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VisitArticleLogic {
	return &VisitArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 访问文章
func (l *VisitArticleLogic) VisitArticle(in *articlerpc.IdReq) (*articlerpc.CountResp, error) {
	id := cast.ToString(in.Id)
	key := rediskey.GetArticleVisitCountKey(id)

	// 浏览量+1
	_, err := l.svcCtx.Redis.ZIncrBy(l.ctx, key, 1, id).Result()
	if err != nil {
		return nil, err
	}

	return &articlerpc.CountResp{}, nil
}
