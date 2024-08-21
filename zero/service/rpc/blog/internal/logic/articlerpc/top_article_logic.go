package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type TopArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTopArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TopArticleLogic {
	return &TopArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 置顶文章
func (l *TopArticleLogic) TopArticle(in *blog.TopArticleReq) (*blog.EmptyResp, error) {
	// todo: add your logic here and delete this line

	return &blog.EmptyResp{}, nil
}
