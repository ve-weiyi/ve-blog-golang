package api

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
)

type GetApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取API详情
func NewGetApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetApiLogic {
	return &GetApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetApiLogic) GetApi(req *types.GetApiReq) (resp *types.ApiVO, err error) {
	return nil, nil
}
