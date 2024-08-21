package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecycleArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRecycleArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecycleArticleLogic {
	return &RecycleArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 回收文章
func (l *RecycleArticleLogic) RecycleArticle(in *blog.RecycleArticleReq) (*blog.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &blog.EmptyResp{}, nil
}
