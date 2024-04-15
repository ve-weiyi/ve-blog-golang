package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserRoleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserRoleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserRoleLogic {
	return &UpdateUserRoleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 修改用户角色
func (l *UpdateUserRoleLogic) UpdateUserRole(in *account.UpdateUserRoleReq) (*account.EmptyResp, error) {
	ua, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", in.UserId)
	if err != nil {
		return nil, err
	}

	// 删除用户角色
	_, err = l.svcCtx.UserRoleModel.BatchDelete(l.ctx, "user_id = ?", ua.Id)
	if err != nil {
		return nil, err
	}

	var userRoles []*model.UserRole
	for _, roleId := range in.RoleIds {
		m := &model.UserRole{
			UserId: in.UserId,
			RoleId: roleId,
		}
		userRoles = append(userRoles, m)
	}

	// 添加用户角色
	_, err = l.svcCtx.UserRoleModel.BatchCreate(l.ctx, userRoles...)
	if err != nil {
		return nil, err
	}

	return &account.EmptyResp{}, nil
}
