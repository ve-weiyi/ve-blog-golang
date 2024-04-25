package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleHomeListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取文章列表
func NewFindArticleHomeListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleHomeListLogic {
	return &FindArticleHomeListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleHomeListLogic) FindArticleHomeList(reqCtx *types.RestHeader, req *types.PageQuery) (resp *types.PageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
