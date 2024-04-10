package role

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleApisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRoleApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleApisLogic {
	return &UpdateRoleApisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRoleApisLogic) UpdateRoleApis(req *types.UpdateRoleApisReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
