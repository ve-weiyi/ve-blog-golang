package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询文章
func NewFindArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleLogic {
	return &FindArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleLogic) FindArticle(req *types.IdReq) (resp *types.ArticleBackDTO, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.ArticleRpc.FindArticle(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertArticleBackTypes(out), nil
}
