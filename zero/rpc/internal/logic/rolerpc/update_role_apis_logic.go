package rolerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRoleApisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRoleApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRoleApisLogic {
	return &UpdateRoleApisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新角色资源
func (l *UpdateRoleApisLogic) UpdateRoleApis(in *account.UpdateRoleApisReq) (*account.EmptyResp, error) {
	// 删除
	_, err := l.svcCtx.RoleApiModel.BatchDelete(l.ctx, "role_id in (?)", in.RoleId)
	if err != nil {
		return nil, err
	}

	var roleApis []*model.RoleApi
	for _, apiId := range in.ApiIds {
		m := &model.RoleApi{
			RoleId: in.RoleId,
			ApiId:  apiId,
		}

		roleApis = append(roleApis, m)
	}

	// 添加
	_, err = l.svcCtx.RoleApiModel.BatchCreate(l.ctx, roleApis...)
	if err != nil {
		return nil, err
	}

	return &account.EmptyResp{}, nil
}
