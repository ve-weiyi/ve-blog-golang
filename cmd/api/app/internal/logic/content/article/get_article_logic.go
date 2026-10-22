package article

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/enrich"
	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/contentservice"
)

type GetArticleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文章详情
func NewGetArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleLogic {
	return &GetArticleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleLogic) GetArticle(req *types.GetArticleReq) (resp *types.ArticleDetails, err error) {
	articleOut, err := l.svcCtx.ContentService.GetArticle(l.ctx, &contentservice.GetArticleRequest{
		Id:         req.ArticleId,
		PublicOnly: true,
	})
	if err != nil {
		return nil, err
	}

	// 浏览量是详情页的副作用，且只在文章可见时计数
	if _, err = l.svcCtx.ContentService.IncrementArticleView(l.ctx, &contentservice.IncrementArticleViewRequest{
		Id: req.ArticleId,
	}); err != nil {
		l.Errorf("增加文章浏览量失败: %v", err)
	}
	if err != nil {
		return nil, err
	}

	relationOut, err := l.svcCtx.ContentService.GetArticleRelation(l.ctx, &contentservice.GetArticleRelationRequest{
		Id: req.ArticleId,
	})
	if err != nil {
		return nil, err
	}

	usm, err := enrich.GetUserInfos(l.ctx, l.svcCtx, []string{articleOut.Article.UserId})
	if err != nil {
		return nil, err
	}

	resp = &types.ArticleDetails{
		ArticleHome:          *convertArticleHomeTypes(articleOut.Article),
		Author:               usm[articleOut.Article.UserId],
		LastArticle:          convertArticlePreviewTypes(relationOut.Last),
		NextArticle:          convertArticlePreviewTypes(relationOut.Next),
		RecommendArticleList: make([]*types.ArticlePreview, 0),
		NewestArticleList:    make([]*types.ArticlePreview, 0),
	}

	for _, v := range relationOut.Recommends {
		resp.RecommendArticleList = append(resp.RecommendArticleList, convertArticlePreviewTypes(v))
	}

	for _, v := range relationOut.Latest {
		resp.NewestArticleList = append(resp.NewestArticleList, convertArticlePreviewTypes(v))
	}

	return
}
