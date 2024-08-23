package article

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/articlerpc"
)

type FindArticleListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询文章列表
func NewFindArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleListLogic {
	return &FindArticleListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleListLogic) FindArticleList(req *types.ArticleQuery) (resp *types.PageResp, err error) {
	in := &articlerpc.FindArticleListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sorts:    "id desc",
	}
	out, err := l.svcCtx.ArticleRpc.FindArticleList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.ArticleBackDTO
	// 转换数据
	for _, v := range out.List {
		m := ConvertArticleBackTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = out.Total
	resp.List = list
	return
}

func ConvertArticleBackTypes(in *articlerpc.ArticleDetails) (out *types.ArticleBackDTO) {
	var category string
	if in.Category != nil {
		category = in.Category.CategoryName
	}

	var tags []string
	if len(in.TagList) > 0 {
		for _, t := range in.TagList {
			tags = append(tags, t.TagName)
		}
	}

	return &types.ArticleBackDTO{
		Id:             in.Id,
		ArticleCover:   in.ArticleCover,
		ArticleTitle:   in.ArticleTitle,
		ArticleContent: in.ArticleContent,
		ArticleType:    in.ArticleType,
		OriginalUrl:    in.OriginalUrl,
		IsTop:          in.IsTop,
		Status:         in.Status,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
		CategoryName:   category,
		TagNameList:    tags,
		LikeCount:      in.LikeCount,
		ViewsCount:     0,
	}
}
