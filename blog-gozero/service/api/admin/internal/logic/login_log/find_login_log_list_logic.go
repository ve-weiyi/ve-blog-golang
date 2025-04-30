package login_log

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/common/apiutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLoginLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询用户登录历史
func NewFindLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLoginLogListLogic {
	return &FindLoginLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindLoginLogListLogic) FindLoginLogList(req *types.LoginLogQuery) (resp *types.PageResp, err error) {
	in := &syslogrpc.FindLoginLogListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sorts:    req.Sorts,
	}

	out, err := l.svcCtx.SyslogRpc.FindLoginLogList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var uids []string
	for _, v := range out.List {
		uids = append(uids, v.UserId)
	}

	// 查询用户信息
	usm, err := apiutils.GetUserInfos(l.ctx, l.svcCtx, uids)
	if err != nil {
		return nil, err
	}

	var list []*types.LoginLogBackVO
	for _, v := range out.List {
		m := ConvertLoginLogTypes(v, usm)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}

func ConvertLoginLogTypes(in *syslogrpc.LoginLogDetails, usm map[string]*accountrpc.User) (out *types.LoginLogBackVO) {
	out = &types.LoginLogBackVO{
		Id:        in.Id,
		UserId:    in.UserId,
		LoginType: in.LoginType,
		Agent:     in.Agent,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		LoginAt:   in.LoginAt,
		LogoutAt:  in.LogoutAt,
	}

	// 用户信息
	if in.UserId != "" {
		user, ok := usm[in.UserId]
		if ok && user != nil {
			out.User = &types.UserInfo{
				UserId:   user.UserId,
				Username: user.Username,
				Avatar:   user.Avatar,
				Nickname: user.Nickname,
			}
		}
	}

	return out
}
