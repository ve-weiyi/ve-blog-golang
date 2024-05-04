package mine

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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

func (l *GetUserApisLogic) GetUserApis(reqCtx *types.RestHeader, req *types.EmptyReq) (resp *types.UserApisResp, err error) {
	in := &blog.UserReq{
		UserId: cast.ToInt64(reqCtx.HeaderXUserId),
	}

	out, err := l.svcCtx.UserRpc.GetUserApis(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserApi
	jsonconv.ObjectToObject(out.List, &list)

	resp = &types.UserApisResp{}
	resp.List = list
	return
}
