package permissionrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建菜单
func (l *AddMenuLogic) AddMenu(in *permissionrpc.MenuNew) (*permissionrpc.MenuDetails, error) {
	entity := ConvertMenuIn(in)

	_, err := l.svcCtx.MenuModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return ConvertMenuOut(entity), nil
}
