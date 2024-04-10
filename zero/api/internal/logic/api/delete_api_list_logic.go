package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteApiListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiListLogic {
	return &DeleteApiListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteApiListLogic) DeleteApiList(req *types.IdsReq) (resp *types.BatchResult, err error) {
	in := convert.ConvertIdsReq(req)

	_, err = l.svcCtx.ApiRpc.DeleteApiList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResult{}, nil
}
