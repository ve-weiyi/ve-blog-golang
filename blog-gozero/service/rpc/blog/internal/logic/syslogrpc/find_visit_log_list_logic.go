package syslogrpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindVisitLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindVisitLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindVisitLogListLogic {
	return &FindVisitLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询操作访问列表
func (l *FindVisitLogListLogic) FindVisitLogList(in *syslogrpc.FindVisitLogListReq) (*syslogrpc.FindVisitLogListResp, error) {
	page, size, sorts, conditions, params := convertVisitLogQuery(in)

	records, total, err := l.svcCtx.TVisitLogModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*syslogrpc.VisitLogDetails
	for _, v := range records {
		list = append(list, convertVisitLogOut(v))
	}

	return &syslogrpc.FindVisitLogListResp{
		List:  list,
		Total: total,
	}, nil
}

func convertVisitLogQuery(in *syslogrpc.FindVisitLogListReq) (page int, size int, sorts string, conditions string, params []any) {
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

	if in.TerminalId != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions += " terminal_id = ?"
		params = append(params, in.TerminalId)
	}

	if in.PageName != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions += " page like ?"
		params = append(params, "%"+in.PageName+"%")
	}

	return
}

func convertVisitLogOut(in *model.TVisitLog) (out *syslogrpc.VisitLogDetails) {
	out = &syslogrpc.VisitLogDetails{
		Id:         in.Id,
		UserId:     in.UserId,
		TerminalId: in.TerminalId,
		PageName:   in.PageName,
		IpAddress:  in.IpAddress,
		IpSource:   in.IpSource,
		Os:         in.Os,
		Browser:    in.Browser,
		CreatedAt:  in.CreatedAt.Unix(),
		UpdatedAt:  in.UpdatedAt.Unix(),
	}

	return out
}
