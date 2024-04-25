package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除页面
func NewDeletePageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePageListLogic {
	return &DeletePageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletePageListLogic) DeletePageList(reqCtx *types.RestHeader, req *types.IdsReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
