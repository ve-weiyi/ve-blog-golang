package friendlinkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteFriendLinkListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteFriendLinkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteFriendLinkListLogic {
	return &DeleteFriendLinkListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除友链
func (l *DeleteFriendLinkListLogic) DeleteFriendLinkList(in *blog.IdsReq) (*blog.BatchResp, error) {
	rows, err := l.svcCtx.FriendLinkModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: rows,
	}, nil
}
