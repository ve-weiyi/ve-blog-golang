package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleClassifyCategoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过标签或者id获取文章列表
func NewFindArticleClassifyCategoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleClassifyCategoryLogic {
	return &FindArticleClassifyCategoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleClassifyCategoryLogic) FindArticleClassifyCategory(reqCtx *types.RestHeader, req *types.ArticleClassifyCategoryReq) (resp *types.ArticleClassifyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
