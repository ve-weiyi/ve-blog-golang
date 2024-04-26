package articlerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleByTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindArticleByTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleByTagLogic {
	return &FindArticleByTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文章列表
func (l *FindArticleByTagLogic) FindArticleByTag(in *blog.FindArticleByTagReq) (*blog.ArticlePageResp, error) {
	ts, err := l.svcCtx.ArticleTagModel.FindALL(l.ctx, "tag_id in (?)", in.TagIds)
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, v := range ts {
		ids = append(ids, v.ArticleId)
	}

	result, err := l.svcCtx.ArticleModel.FindALL(l.ctx, "id in (?)", ids)
	if err != nil {
		return nil, err
	}

	var list []*blog.Article
	for _, v := range result {
		list = append(list, convert.ConvertArticleModelToPb(v))
	}

	return &blog.ArticlePageResp{
		List: list,
	}, nil
}
