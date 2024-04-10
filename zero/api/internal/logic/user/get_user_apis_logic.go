package user

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserApisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserApisLogic {
	return &GetUserApisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserApisLogic) GetUserApis(req *types.EmptyReq) (resp []types.UserApiDTO, err error) {
	in := convert.EmptyReq()
	out, err := l.svcCtx.UserRpc.GetUserApis(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = make([]types.UserApiDTO, 0)
	jsonconv.ObjectMarshal(out.List, &resp)

	return
}
