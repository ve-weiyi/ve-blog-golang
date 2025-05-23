package resourcerpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFileFolderListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindFileFolderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFileFolderListLogic {
	return &FindFileFolderListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询文件夹列表
func (l *FindFileFolderListLogic) FindFileFolderList(in *resourcerpc.FindFileFolderListReq) (*resourcerpc.FindFileFolderListResp, error) {
	page, size, sorts, conditions, params := convertFileFolderQuery(in)

	records, total, err := l.svcCtx.TFileFolderModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*resourcerpc.FileFolderDetails
	for _, v := range records {
		list = append(list, convertFileFolderOut(v))
	}

	return &resourcerpc.FindFileFolderListResp{
		List:  list,
		Total: total,
	}, nil
}

func convertFileFolderQuery(in *resourcerpc.FindFileFolderListReq) (page int, size int, sorts string, conditions string, params []any) {
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
		conditions += " file_path = ?"
		params = append(params, in.FilePath)
	}

	return
}
