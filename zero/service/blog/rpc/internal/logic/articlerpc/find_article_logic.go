package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleLogic {
	return &FindArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindArticleLogic) FindArticle(in *blog.IdReq) (*blog.Article, error) {
	entity, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return convert.ConvertArticleModelToPb(entity), nil
}
