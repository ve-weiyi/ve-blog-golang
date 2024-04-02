package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateArticleTopLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateArticleTopLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateArticleTopLogic {
	return &UpdateArticleTopLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateArticleTopLogic) UpdateArticleTop(req *types.ArticleTopReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
