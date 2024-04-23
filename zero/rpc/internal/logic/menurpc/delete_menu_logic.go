package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuLogic {
	return &DeleteMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除菜单
func (l *DeleteMenuLogic) DeleteMenu(in *account.IdReq) (*account.BatchResp, error) {
	result, err := l.svcCtx.MenuModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	result2, err := l.svcCtx.MenuModel.DeleteBatch(l.ctx, "parent_id = ? ", in.Id)
	if err != nil {
		return nil, err
	}

	return &account.BatchResp{
		SuccessCount: result + result2,
	}, nil
}
