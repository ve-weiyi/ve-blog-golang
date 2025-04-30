package user

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/restx"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
)

type GetUserLoginHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询用户登录历史
func NewGetUserLoginHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLoginHistoryListLogic {
	return &GetUserLoginHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLoginHistoryListLogic) GetUserLoginHistoryList(req *types.UserLoginHistoryQuery) (resp *types.PageResp, err error) {
	in := &syslogrpc.FindLoginLogListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		UserId:   cast.ToString(l.ctx.Value(restx.HeaderUid)),
	}

	out, err := l.svcCtx.SyslogRpc.FindLoginLogList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserLoginHistory
	for _, v := range out.List {
		m := ConvertUserLoginHistoryTypes(v)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}

func ConvertUserLoginHistoryTypes(in *syslogrpc.LoginLogDetails) *types.UserLoginHistory {
	return &types.UserLoginHistory{
		Id:        in.Id,
		LoginType: in.LoginType,
		Os:        in.Os,
		Browser:   in.Browser,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		LoginAt:   in.LoginAt,
		LogoutAt:  in.LogoutAt,
	}
}
