package syslogrpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOperationLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOperationLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOperationLogListLogic {
	return &FindOperationLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取操作记录列表
func (l *FindOperationLogListLogic) FindOperationLogList(in *syslogrpc.FindOperationLogListReq) (*syslogrpc.FindOperationLogListResp, error) {
	page, size, sorts, conditions, params := convertOperationLogQuery(in)

	records, total, err := l.svcCtx.TOperationLogModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*syslogrpc.OperationLogDetails
	for _, v := range records {
		list = append(list, convertOperationLogOut(v))
	}

	return &syslogrpc.FindOperationLogListResp{
		List:  list,
		Total: total,
	}, nil
}

func convertOperationLogQuery(in *syslogrpc.FindOperationLogListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")
	if sorts == "" {
		sorts = "id desc"
	}

	if in.Keywords != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions += " opt_desc = ?"
		params = append(params, "%"+in.Keywords+"%")
	}
	return
}

func convertOperationLogOut(in *model.TOperationLog) (out *syslogrpc.OperationLogDetails) {
	out = &syslogrpc.OperationLogDetails{
		Id:             in.Id,
		UserId:         in.UserId,
		TerminalId:     in.TerminalId,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		OptModule:      in.OptModule,
		OptDesc:        in.OptDesc,
		RequestUri:     in.RequestUri,
		RequestMethod:  in.RequestMethod,
		RequestData:    in.RequestData,
		ResponseData:   in.ResponseData,
		ResponseStatus: in.ResponseStatus,
		Cost:           in.Cost,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
	}

	return out
}
