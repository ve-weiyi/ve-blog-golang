package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserMenusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserMenusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserMenusLogic {
	return &GetUserMenusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserMenusLogic) GetUserMenus() (resp []types.MenuDetailsDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
