package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOnlineUserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询在线用户列表
func NewFindOnlineUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOnlineUserListLogic {
	return &FindOnlineUserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindOnlineUserListLogic) FindOnlineUserList(req *types.UserQuery) (resp *types.PageResp, err error) {
	in := &accountrpc.FindUserListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Nickname: req.Nickname,
	}

	out, err := l.svcCtx.AccountRpc.FindUserOnlineList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserInfoResp
	for _, v := range out.List {
		m := ConvertUserTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
