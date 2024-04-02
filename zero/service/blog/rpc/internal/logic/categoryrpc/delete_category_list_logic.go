package categoryrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteCategoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCategoryListLogic {
	return &DeleteCategoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除文章分类
func (l *DeleteCategoryListLogic) DeleteCategoryList(in *blog.IdsReq) (*blog.BatchResp, error) {
	result, err := l.svcCtx.CategoryModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: result,
	}, nil
}
