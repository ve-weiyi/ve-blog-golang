package menu

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMenuDetailsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMenuDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMenuDetailsListLogic {
	return &FindMenuDetailsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMenuDetailsListLogic) FindMenuDetailsList(req *types.PageQuery) (resp []types.MenuDetailsDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
