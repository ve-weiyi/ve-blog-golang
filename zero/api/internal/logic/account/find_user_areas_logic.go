package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserAreasLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserAreasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserAreasLogic {
	return &FindUserAreasLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserAreasLogic) FindUserAreas(req *types.PageQuery) (resp []types.UserAreaDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
