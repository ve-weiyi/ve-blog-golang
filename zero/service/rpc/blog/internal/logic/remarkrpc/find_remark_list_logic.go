package remarkrpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/remarkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRemarkListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRemarkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRemarkListLogic {
	return &FindRemarkListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询留言列表
func (l *FindRemarkListLogic) FindRemarkList(in *remarkrpc.FindRemarkListReq) (*remarkrpc.FindRemarkListResp, error) {
	var (
		page       int
		size       int
		sorts      string
		conditions string
		params     []interface{}
	)

	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")

	result, err := l.svcCtx.RemarkModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	count, err := l.svcCtx.RemarkModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*remarkrpc.RemarkDetails
	for _, v := range result {
		list = append(list, convertRemarkOut(v))
	}

	return &remarkrpc.FindRemarkListResp{
		List:  list,
		Total: count,
	}, nil
}
