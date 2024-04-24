package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

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

func (l *CleanApiListLogic) CleanApiList(req *types.EmptyReq) (resp *types.BatchResp, err error) {
	// todo: add your logic here and delete this line

	return
}
