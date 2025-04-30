package syslogrpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindLoginLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindLoginLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindLoginLogListLogic {
	return &FindLoginLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询登录记录列表
func (l *FindLoginLogListLogic) FindLoginLogList(in *syslogrpc.FindLoginLogListReq) (*syslogrpc.FindLoginLogListResp, error) {
	page, size, sorts, conditions, params := convertLoginLogQuery(in)

	records, total, err := l.svcCtx.TLoginLogModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*syslogrpc.LoginLogDetails
	for _, v := range records {
		list = append(list, convertLoginLogOut(v))
	}

	return &syslogrpc.FindLoginLogListResp{
		List:  list,
		Total: total,
	}, nil
}

func convertLoginLogQuery(in *syslogrpc.FindLoginLogListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")
	if sorts == "" {
		sorts = "id desc"
	}

	if in.UserId != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions += " user_id = ?"
		params = append(params, in.UserId)
	}
	return
}

func convertLoginLogOut(in *model.TLoginLog) (out *syslogrpc.LoginLogDetails) {
	out = &syslogrpc.LoginLogDetails{
		Id:        in.Id,
		UserId:    in.UserId,
		LoginType: in.LoginType,
		AppName:   in.AppName,
		Os:        in.Os,
		Browser:   in.Browser,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		LoginAt:   in.LoginAt.Unix(),
		LogoutAt:  0,
	}

	if in.LogoutAt.Valid {
		out.LogoutAt = in.LogoutAt.Time.Unix()
	}

	return out
}
