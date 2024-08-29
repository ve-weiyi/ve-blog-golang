package account

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/permissionrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserApisLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户接口权限
func NewGetUserApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserApisLogic {
	return &GetUserApisLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserApisLogic) GetUserApis(req *types.EmptyReq) (resp *types.UserApisResp, err error) {
	in := &permissionrpc.UserIdReq{
		UserId: cast.ToInt64(l.ctx.Value("uid")),
	}

	out, err := l.svcCtx.PermissionRpc.FindUserApis(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserApi
	jsonconv.ObjectToObject(out.List, &list)

	resp = &types.UserApisResp{}
	resp.List = list
	return
}
