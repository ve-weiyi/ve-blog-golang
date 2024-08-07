package menu

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除菜单
func NewBatchDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteMenuLogic {
	return &BatchDeleteMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteMenuLogic) BatchDeleteMenu(req *types.IdsReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
