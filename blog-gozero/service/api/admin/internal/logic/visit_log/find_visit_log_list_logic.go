package visit_log

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/common/apiutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/syslogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVisitLogListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取操作记录列表
func NewFindVisitLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVisitLogListLogic {
	return &FindVisitLogListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindVisitLogListLogic) FindVisitLogList(req *types.VisitLogQuery) (resp *types.PageResp, err error) {
	in := &syslogrpc.FindVisitLogListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
		Sorts:    req.Sorts,
		Keywords: req.Keywords,
	}

	out, err := l.svcCtx.SyslogRpc.FindVisitLogList(l.ctx, in)
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

	var list []*types.VisitLogBackVO
	for _, v := range out.List {
		m := ConvertVisitLogTypes(v, usm)
		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}

func ConvertVisitLogTypes(in *syslogrpc.VisitLogDetails, usm map[string]*accountrpc.User) (out *types.VisitLogBackVO) {

	out = &types.VisitLogBackVO{
		Id:        in.Id,
		UserId:    in.UserId,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		Os:        in.Os,
		Browser:   in.Browser,
		Page:      in.Page,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
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
