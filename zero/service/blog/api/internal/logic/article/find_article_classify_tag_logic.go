package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleClassifyTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过标签或者id获取文章列表
func NewFindArticleClassifyTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleClassifyTagLogic {
	return &FindArticleClassifyTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleClassifyTagLogic) FindArticleClassifyTag(reqCtx *types.RestHeader, req *types.ArticleClassifyTagReq) (resp *types.ArticleClassifyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
