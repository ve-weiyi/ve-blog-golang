package websiterpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

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
func (l *FindPageListLogic) FindPageList(in *websiterpc.FindPageListReq) (*websiterpc.FindPageListResp, error) {
	page, size, sorts, conditions, params := convertPageQuery(in)

	result, err := l.svcCtx.TPageModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*websiterpc.PageDetails
	for _, v := range result {
		list = append(list, convertPageOut(v))
	}

	return &websiterpc.FindPageListResp{
		List: list,
	}, nil
}

func convertPageQuery(in *websiterpc.FindPageListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")
	if sorts == "" {
		sorts = "id desc"
	}

	return
}
