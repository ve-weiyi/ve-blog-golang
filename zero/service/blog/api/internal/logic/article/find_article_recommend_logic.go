package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 文章相关推荐
func NewFindArticleRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleRecommendLogic {
	return &FindArticleRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleRecommendLogic) FindArticleRecommend(req *types.IdReq) (resp *types.ArticleRecommendResp, err error) {
	// todo: add your logic here and delete this line

	return
}
