package category

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除文章分类
func NewDeleteCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryListLogic {
	return &DeleteCategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCategoryListLogic) DeleteCategoryList(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := convert.ConvertIdsReq(req)

	out, err := l.svcCtx.CategoryRpc.DeleteCategoryList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}, nil
}
