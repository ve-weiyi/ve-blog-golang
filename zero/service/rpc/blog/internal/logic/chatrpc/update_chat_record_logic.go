package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateChatRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateChatRecordLogic {
	return &UpdateChatRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新聊天记录
func (l *UpdateChatRecordLogic) UpdateChatRecord(in *chatrpc.ChatRecordNewReq) (*chatrpc.ChatRecordDetails, error) {
	// todo: add your logic here and delete this line

	return &chatrpc.ChatRecordDetails{}, nil
}
