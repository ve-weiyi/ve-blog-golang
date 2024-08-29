package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddChatRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddChatRecordLogic {
	return &AddChatRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建聊天记录
func (l *AddChatRecordLogic) AddChatRecord(in *chatrpc.ChatRecordNew) (*chatrpc.ChatRecordDetails, error) {
	// todo: add your logic here and delete this line

	return &chatrpc.ChatRecordDetails{}, nil
}
