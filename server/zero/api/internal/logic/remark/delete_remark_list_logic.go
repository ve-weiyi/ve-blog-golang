package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRemarkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteRemarkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRemarkListLogic {
	return &DeleteRemarkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRemarkListLogic) DeleteRemarkList(req *types.IdsReq) (resp *types.BatchResult, err error) {
	// todo: add your logic here and delete this line

	return
}
