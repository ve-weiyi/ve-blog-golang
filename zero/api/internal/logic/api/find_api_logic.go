package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/client/apirpc"

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
	in := &apirpc.IdReq{
		Id: req.Id,
	}

	out, err := l.svcCtx.ApiRpc.FindApi(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = convertApi(out)

	return resp, nil
}

func convertApi(in *apirpc.Api) (out *types.Api) {
	return &types.Api{
		ID:        in.Id,
		Name:      in.Name,
		Path:      in.Path,
		Method:    in.Method,
		ParentID:  in.ParentId,
		Traceable: in.Traceable,
		Status:    in.Status,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
