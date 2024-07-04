package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRolesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRolesLogic {
	return &UpdateUserRolesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserRolesLogic) UpdateUserRoles(req *types.UpdateUserRolesReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
