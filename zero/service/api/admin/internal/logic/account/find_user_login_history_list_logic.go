package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLoginHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserLoginHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLoginHistoryListLogic {
	return &FindUserLoginHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserLoginHistoryListLogic) FindUserLoginHistoryList(req *types.PageQuery) (resp *types.PageResp, err error) {
	in := &accountrpc.FindLoginHistoryListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		UserId:   l.ctx.Value("uid").(int64),
	}
	out, err := l.svcCtx.AccountRpc.FindUserLoginHistoryList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserLoginHistory
	for _, role := range out.List {
		list = append(list, ConvertUserLoginHistoryTypes(role))
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
