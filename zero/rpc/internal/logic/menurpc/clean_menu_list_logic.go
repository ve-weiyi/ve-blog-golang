package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanMenuListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCleanMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanMenuListLogic {
	return &CleanMenuListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 清空菜单列表
func (l *CleanMenuListLogic) CleanMenuList(in *account.EmptyReq) (*account.BatchResp, error) {
	row, err := l.svcCtx.MenuModel.DeleteBatch(l.ctx, "1 = 1")
	if err != nil {
		return nil, err
	}

	return &account.BatchResp{
		SuccessCount: row,
	}, nil
}
