package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteChatRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteChatRecordLogic {
	return &DeleteChatRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除聊天记录
func (l *DeleteChatRecordLogic) DeleteChatRecord(in *chatrpc.IdsReq) (*chatrpc.BatchResp, error) {
	// todo: add your logic here and delete this line

	return &chatrpc.BatchResp{}, nil
}
