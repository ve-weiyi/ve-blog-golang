package api

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/permissionservice"
)

type UpdateApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新API
func NewUpdateApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateApiLogic {
	return &UpdateApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateApiLogic) UpdateApi(req *types.UpdateApiReq) (resp *types.EmptyResp, err error) {
	_, err = l.svcCtx.PermissionService.UpdateApi(l.ctx, &permissionservice.UpdateApiRequest{
		Id:        req.Id,
		ParentId:  req.ParentId,
		Name:      req.Name,
		Path:      req.Path,
		Method:    req.Method,
		Traceable: req.Traceable,
		Status:    req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
