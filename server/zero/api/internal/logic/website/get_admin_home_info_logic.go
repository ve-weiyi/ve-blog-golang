package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAdminHomeInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAdminHomeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAdminHomeInfoLogic {
	return &GetAdminHomeInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAdminHomeInfoLogic) GetAdminHomeInfo() (resp *types.AdminHomeInfo, err error) {
	// todo: add your logic here and delete this line

	return
}
