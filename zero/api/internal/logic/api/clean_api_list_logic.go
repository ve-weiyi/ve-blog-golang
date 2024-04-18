package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CleanApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCleanApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CleanApiListLogic {
	return &CleanApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CleanApiListLogic) CleanApiList(req *types.EmptyReq) (resp *types.BatchResult, err error) {
	// todo: add your logic here and delete this line

	return
}
