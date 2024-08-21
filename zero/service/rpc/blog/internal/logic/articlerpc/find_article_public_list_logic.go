package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/global"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticlePublicListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindArticlePublicListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticlePublicListLogic {
	return &FindArticlePublicListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章列表
func (l *FindArticlePublicListLogic) FindArticlePublicList(in *blog.FindArticleListReq) (*blog.FindArticleListResp, error) {
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
	conditions = "status = ?"
	params = append(params, global.ArticleStatusPublic)

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

	var list []*blog.ArticleDetails
	for _, v := range records {
		list = append(list, convertArticleOut(v, acm, atm))
	}

	return &blog.FindArticleListResp{
		List:  list,
		Total: count,
	}, nil
}
