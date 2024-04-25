package api

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

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

func (l *FindApiLogic) FindApi(reqCtx *types.RestHeader, req *types.IdReq) (resp *types.Api, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.ApiRpc.FindApi(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertApiTypes(out), nil
}
