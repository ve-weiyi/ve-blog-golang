package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteArticleLogic {
	return &DeleteArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteArticleLogic) DeleteArticle(in *blog.IdReq) (*blog.BatchResp, error) {
	rows, err := l.svcCtx.ArticleModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: rows,
	}, nil
}
