package api

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/permissionservice"
)

type CleanApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 清空接口列表
func NewCleanApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanApiLogic {
	return &CleanApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CleanApiLogic) CleanApi(req *types.EmptyReq) (resp *types.CleanApiResp, err error) {
	out, err := l.svcCtx.PermissionService.CleanApis(l.ctx, &permissionservice.CleanApisRequest{})
	if err != nil {
		return nil, err
	}

	return &types.CleanApiResp{
		SuccessCount: out.SuccessCount,
	}, nil
}
