package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/internal/global"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/articlerpc"
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
func (l *FindArticlesByTagLogic) FindArticlesByTag(in *articlerpc.FindArticlesByTagReq) (*articlerpc.FindArticleListResp, error) {
	tag, err := l.svcCtx.TagModel.FindOneByTagName(l.ctx, in.TagName)
	if err != nil {
		return nil, err
	}

	ats, err := l.svcCtx.ArticleTagModel.FindALL(l.ctx, "tag_id = ?", tag.Id)
	if err != nil {
		return nil, err
	}

	var articlerpcIds []int64
	for _, v := range ats {
		articlerpcIds = append(articlerpcIds, v.ArticleId)
	}

	// 查询文章信息
	records, err := l.svcCtx.ArticleModel.FindALL(l.ctx, "id in (?) and status = ?", articlerpcIds, global.ArticleStatusPublic)
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
		List: list,
	}, nil
}
