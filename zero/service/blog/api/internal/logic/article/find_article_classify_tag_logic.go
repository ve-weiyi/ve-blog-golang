package article

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleClassifyTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过标签或者id获取文章列表
func NewFindArticleClassifyTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleClassifyTagLogic {
	return &FindArticleClassifyTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleClassifyTagLogic) FindArticleClassifyTag(reqCtx *types.RestHeader, req *types.ArticleClassifyReq) (resp *types.ArticleClassifyResp, err error) {
	cs, err := l.svcCtx.TagRpc.FindTagList(l.ctx, &blog.PageQuery{
		Page:       1,
		PageSize:   1,
		Sorts:      "id desc",
		Conditions: "tag_name = ?",
		Args:       []string{cast.ToString(req.ClassifyName)},
	})
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, v := range cs.List {
		ids = append(ids, v.Id)
	}

	as, err := l.svcCtx.ArticleRpc.FindArticleByTag(l.ctx, &blog.FindArticleByTagReq{
		TagIds: ids,
	})
	if err != nil {
		return nil, err
	}

	var list []*types.ArticleHome
	for _, v := range as.List {
		m := convert.ConvertArticleHomeTypes(v)
		list = append(list, m)
	}

	resp = &types.ArticleClassifyResp{}
	resp.ConditionName = req.ClassifyName
	resp.ArticleList = list
	return
}
