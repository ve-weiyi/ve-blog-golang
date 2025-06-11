package syslogrpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/syslogrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

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

// 查询上传记录列表
func (l *FindUploadLogListLogic) FindUploadLogList(in *syslogrpc.FindUploadLogListReq) (*syslogrpc.FindUploadLogListResp, error) {
	page, size, sorts, conditions, params := convertUploadLogQuery(in)

	records, total, err := l.svcCtx.TUploadLogModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*syslogrpc.UploadLogDetails
	for _, v := range records {
		list = append(list, convertUploadLogOut(v))
	}

	return &syslogrpc.FindUploadLogListResp{
		List:  list,
		Total: total,
	}, nil
}

func convertUploadLogQuery(in *syslogrpc.FindUploadLogListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")
	if sorts == "" {
		sorts = "id desc"
	}

	if in.FilePath != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions += " file_path like ?"
		params = append(params, in.FilePath+"%")
	}

	if in.FileName != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions += " file_name like ?"
		params = append(params, in.FileName+"%")
	}

	if in.FileType != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions += " file_type = ?"
		params = append(params, in.FileType)
	}

	return
}

func convertUploadLogOut(in *model.TUploadLog) (out *syslogrpc.UploadLogDetails) {
	out = &syslogrpc.UploadLogDetails{
		Id:        in.Id,
		UserId:    in.UserId,
		FilePath:  in.FilePath,
		FileName:  in.FileName,
		FileType:  in.FileType,
		FileSize:  in.FileSize,
		FileMd5:   in.FileMd5,
		FileUrl:   in.FileUrl,
		CreatedAt: in.CreatedAt.Unix(),
		UpdatedAt: in.UpdatedAt.Unix(),
	}

	return out
}
