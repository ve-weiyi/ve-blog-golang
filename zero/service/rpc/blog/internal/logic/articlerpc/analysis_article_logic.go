package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AnalysisArticleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAnalysisArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AnalysisArticleLogic {
	return &AnalysisArticleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章数量
func (l *AnalysisArticleLogic) AnalysisArticle(in *blog.EmptyReq) (*blog.AnalysisArticleResp, error) {
	ac, err := l.svcCtx.ArticleModel.FindCount(l.ctx, "")
	if err != nil {
		return nil, err
	}

	cc, err := l.svcCtx.CategoryModel.FindCount(l.ctx, "")
	if err != nil {
		return nil, err
	}

	tc, err := l.svcCtx.TagModel.FindCount(l.ctx, "")
	if err != nil {
		return nil, err
	}

	out := &blog.AnalysisArticleResp{
		ArticleCount:  ac,
		CategoryCount: cc,
		TagCount:      tc,
	}

	return out, nil
}
