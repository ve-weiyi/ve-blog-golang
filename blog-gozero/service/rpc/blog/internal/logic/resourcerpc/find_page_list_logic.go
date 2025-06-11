package resourcerpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPageListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPageListLogic {
	return &FindPageListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询页面列表
func (l *FindPageListLogic) FindPageList(in *resourcerpc.FindPageListReq) (*resourcerpc.FindPageListResp, error) {
	page, size, sorts, conditions, params := convertPageQuery(in)

	records, total, err := l.svcCtx.TPageModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*resourcerpc.PageDetails
	for _, v := range records {
		list = append(list, convertPageOut(v))
	}

	return &resourcerpc.FindPageListResp{
		List:  list,
		Total: total,
	}, nil
}

func convertPageQuery(in *resourcerpc.FindPageListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")

	if in.PageName != "" {
		conditions += "page_name like ?"
		params = append(params, "%"+in.PageName+"%")
	}

	return
}
