package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleLogic {
	return &UpdateRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

//
func (l *UpdateRoleLogic) UpdateRole(in *account.Role) (*account.Role, error) {
	entity := convert.ConvertRolePbToModel(in)

	result, err := l.svcCtx.RoleModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertRoleModelToPb(result), nil
}
