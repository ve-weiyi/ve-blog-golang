package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuListLogic {
	return &DeleteMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除菜单
func (l *DeleteMenuListLogic) DeleteMenuList(in *account.IdsReq) (*account.BatchResult, error) {
	result, err := l.svcCtx.MenuModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &account.BatchResult{
		SuccessCount: result,
	}, nil
}
