package article

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/articlerpc"
)

type GetArticleDetailsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取文章详情
func NewGetArticleDetailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetArticleDetailsLogic {
	return &GetArticleDetailsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetArticleDetailsLogic) GetArticleDetails(req *types.IdReq) (resp *types.ArticleDetails, err error) {
	in := &articlerpc.IdReq{
		Id: req.Id,
	}

	// 添加文章访问量
	_, err = l.svcCtx.ArticleRpc.VisitArticle(l.ctx, in)
	if err != nil {
		return nil, err
	}

	out, err := l.svcCtx.ArticleRpc.GetArticle(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 查询关联文章
	relation, err := l.svcCtx.ArticleRpc.GetArticleRelation(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.ArticleDetails{
		ArticleHome:          types.ArticleHome{},
		LastArticle:          nil,
		NextArticle:          nil,
		RecommendArticleList: nil,
		NewestArticleList:    nil,
	}

	resp.ArticleHome = *ConvertArticleHomeTypes(out)

	resp.LastArticle = ConvertArticlePreviewTypes(relation.Last)

	resp.NextArticle = ConvertArticlePreviewTypes(relation.Next)

	for _, v := range relation.Recommend {
		resp.RecommendArticleList = append(resp.RecommendArticleList, ConvertArticlePreviewTypes(v))
	}

	for _, v := range relation.Newest {
		resp.NewestArticleList = append(resp.NewestArticleList, ConvertArticlePreviewTypes(v))
	}

	return
}
