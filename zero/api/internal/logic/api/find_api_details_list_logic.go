package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindApiDetailsListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindApiDetailsListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindApiDetailsListLogic {
	return &FindApiDetailsListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindApiDetailsListLogic) FindApiDetailsList(req *types.PageQuery) (resp []types.ApiDetailsDTO, err error) {
	// todo: add your logic here and delete this line

	return
}
