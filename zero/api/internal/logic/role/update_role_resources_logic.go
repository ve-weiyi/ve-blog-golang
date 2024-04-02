package role

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleResourcesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleResourcesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleResourcesLogic {
	return &UpdateRoleResourcesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleResourcesLogic) UpdateRoleResources(req *types.UpdateRoleApisReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
