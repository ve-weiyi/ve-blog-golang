package syslogrpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUploadLogListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUploadLogListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUploadLogListLogic {
	return &FindUploadLogListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文件列表
func (l *FindUploadLogListLogic) FindUploadLogList(in *syslogrpc.FindUploadLogListReq) (*syslogrpc.FindOperationLogListResp, error) {
	page, size, sorts, conditions, params := convertUploadLogQuery(in)

	result, err := l.svcCtx.TOperationLogModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.TOperationLogModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*syslogrpc.OperationLogDetails
	for _, v := range result {
		list = append(list, convertOperationLogOut(v))
	}

	return &syslogrpc.FindOperationLogListResp{
		List:  list,
		Total: count,
	}, nil
}

func convertUploadLogQuery(in *syslogrpc.FindUploadLogListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")
	if sorts == "" {
		sorts = "id desc"
	}

	return
}
