package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/articlerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindArticleListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleListLogic {
	return &FindArticleListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章数量
func (l *FindArticleListLogic) FindArticleList(in *articlerpc.FindArticleListReq) (*articlerpc.FindArticleListResp, error) {
	var (
		page       int
		size       int
		sorts      string
		conditions string
		params     []interface{}
	)

	page = int(in.Page)
	size = int(in.PageSize)
	sorts = in.Sorts

	// 查询文章信息
	records, err := l.svcCtx.ArticleModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.ArticleModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	acm, err := findCategoryGroupArticle(l.ctx, l.svcCtx, records)
	if err != nil {
		return nil, err

	}

	atm, err := findTagGroupArticle(l.ctx, l.svcCtx, records)
	if err != nil {
		return nil, err
	}

	var list []*articlerpc.ArticleDetails
	for _, v := range records {
		list = append(list, convertArticleOut(v, acm, atm))
	}

	return &articlerpc.FindArticleListResp{
		List:  list,
		Total: count,
	}, nil
}
