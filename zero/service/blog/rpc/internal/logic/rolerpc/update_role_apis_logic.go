package rolerpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
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
func (l *UpdateRoleApisLogic) UpdateRoleApis(in *blog.UpdateRoleApisReq) (*blog.EmptyResp, error) {
	// 删除
	_, err := l.svcCtx.RoleApiModel.DeleteBatch(l.ctx, "role_id in (?)", in.RoleId)
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
	_, err = l.svcCtx.RoleApiModel.InsertBatch(l.ctx, roleApis...)
	if err != nil {
		return nil, err
	}

	return &blog.EmptyResp{}, nil
}
