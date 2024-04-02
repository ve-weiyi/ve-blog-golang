package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindApiLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindApiLogic {
	return &FindApiLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindApiLogic) FindApi(req *types.IdReq) (resp *types.Api, err error) {
	// todo: add your logic here and delete this line

	return
}
