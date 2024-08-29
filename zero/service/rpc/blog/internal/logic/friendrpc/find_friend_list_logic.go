package friendrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/friendrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindFriendListLogic {
	return &FindFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询友链列表
func (l *FindFriendListLogic) FindFriendList(in *friendrpc.FindFriendListReq) (*friendrpc.FindFriendListResp, error) {
	var (
		page       int
		size       int
		sorts      string
		conditions string
		params     []interface{}
	)

	page = int(in.Page)
	size = int(in.PageSize)
	sorts = in.Sorts

	result, err := l.svcCtx.FriendModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*friendrpc.FriendDetails
	for _, v := range result {
		list = append(list, ConvertFriendOut(v))
	}

	return &friendrpc.FindFriendListResp{
		List: list,
	}, nil
}
