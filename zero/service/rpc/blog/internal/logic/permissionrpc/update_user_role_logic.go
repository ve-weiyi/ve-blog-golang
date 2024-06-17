package permissionrpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/permissionrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
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
func (l *UpdateUserRoleLogic) UpdateUserRole(in *permissionrpc.UpdateUserRoleReq) (*permissionrpc.EmptyResp, error) {
	ua, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", in.UserId)
	if err != nil {
		return nil, err
	}

	// 删除用户角色
	_, err = l.svcCtx.UserRoleModel.DeleteBatch(l.ctx, "user_id = ?", ua.Id)
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
	_, err = l.svcCtx.UserRoleModel.InsertBatch(l.ctx, userRoles...)
	if err != nil {
		return nil, err
	}

	return &permissionrpc.EmptyResp{}, nil
}
