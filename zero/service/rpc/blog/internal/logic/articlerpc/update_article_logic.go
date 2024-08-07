package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleLogic {
	return &UpdateArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新文章
func (l *UpdateArticleLogic) UpdateArticle(in *articlerpc.ArticleNew) (*articlerpc.ArticleDetails, error) {
	entity := convertArticleIn(in)
	_, err := l.svcCtx.ArticleModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &articlerpc.ArticleDetails{}, nil
}
