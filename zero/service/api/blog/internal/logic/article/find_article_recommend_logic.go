package article

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindArticleRecommendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取首页推荐文章列表
func NewFindArticleRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindArticleRecommendLogic {
	return &FindArticleRecommendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindArticleRecommendLogic) FindArticleRecommend(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := &blogrpc.FindArticleListReq{
		Page:       req.Page,
		PageSize:   req.PageSize,
		Conditions: "is_top = ?",
		Args:       []string{"1"},
	}
	out, err := l.svcCtx.ArticleRpc.FindArticlePublicList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.ArticleHome
	// 转换数据
	for _, v := range out.List {
		m := convert.ConvertArticleHomeTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = req.Page
	resp.PageSize = req.PageSize
	resp.Total = out.Total
	resp.List = list
	return
}
