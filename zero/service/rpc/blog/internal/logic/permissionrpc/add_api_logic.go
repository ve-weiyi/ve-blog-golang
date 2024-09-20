package permissionrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddApiLogic {
	return &AddApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建接口
func (l *AddApiLogic) AddApi(in *permissionrpc.ApiNewReq) (*permissionrpc.ApiDetails, error) {
	entity := convertApiIn(in)
	_, err := l.svcCtx.TApiModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertApiOut(entity), nil
}
