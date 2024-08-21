package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/global"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticlesByTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindArticlesByTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticlesByTagLogic {
	return &FindArticlesByTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章列表
func (l *FindArticlesByTagLogic) FindArticlesByTag(in *blog.FindArticlesByTagReq) (*blog.FindArticleListResp, error) {
	tag, err := l.svcCtx.TagModel.FindOneByTagName(l.ctx, in.TagName)
	if err != nil {
		return nil, err
	}

	ats, err := l.svcCtx.ArticleTagModel.FindALL(l.ctx, "tag_id = ?", tag.Id)
	if err != nil {
		return nil, err
	}

	var articleIds []int64
	for _, v := range ats {
		articleIds = append(articleIds, v.ArticleId)
	}

	// 查询文章信息
	records, err := l.svcCtx.ArticleModel.FindALL(l.ctx, "id in (?) and status = ?", articleIds, global.ArticleStatusPublic)
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
		List: list,
	}, nil
}
