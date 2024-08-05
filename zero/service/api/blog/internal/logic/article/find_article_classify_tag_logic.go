package article

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleClassifyTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 通过标签获取文章列表
func NewFindArticleClassifyTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleClassifyTagLogic {
	return &FindArticleClassifyTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleClassifyTagLogic) FindArticleClassifyTag(req *types.ArticleClassifyReq) (resp *types.PageResp, err error) {
	cs, err := l.svcCtx.TagRpc.FindTagList(l.ctx, &blogrpc.PageQuery{
		Page:       1,
		PageSize:   1,
		Sorts:      "id desc",
		Conditions: "tag_name = ?",
		Args:       []string{cast.ToString(req.ClassifyName)},
	})
	if err != nil {
		return nil, err
	}

	ids := make([]int64, 0)
	for _, v := range cs.List {
		ids = append(ids, v.Id)
	}

	as, err := l.svcCtx.ArticleRpc.FindArticleByTag(l.ctx, &blogrpc.FindArticleByTagReq{
		TagIds: ids,
	})
	if err != nil {
		return nil, err
	}

	var list []*types.ArticlePreview
	for _, v := range as.List {
		m := convert.ConvertArticlePreviewTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = int64(len(list))
	resp.List = list
	return
}
