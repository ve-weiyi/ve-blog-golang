package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章
func (l *GetArticleLogic) GetArticle(in *blog.IdReq) (*blog.ArticleDetails, error) {
	record, err := l.svcCtx.ArticleModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	acm, err := findCategoryGroupArticle(l.ctx, l.svcCtx, []*model.Article{record})
	if err != nil {
		return nil, err

	}

	atm, err := findTagGroupArticle(l.ctx, l.svcCtx, []*model.Article{record})
	if err != nil {
		return nil, err
	}

	return convertArticleOut(record, acm, atm), nil
}
