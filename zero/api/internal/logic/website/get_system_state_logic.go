package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetSystemStateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetSystemStateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSystemStateLogic {
	return &GetSystemStateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetSystemStateLogic) GetSystemState(req *types.EmptyReq) (resp *types.EmptyResp, err error) {
	// todo: add your logic here and delete this line

	return
}
