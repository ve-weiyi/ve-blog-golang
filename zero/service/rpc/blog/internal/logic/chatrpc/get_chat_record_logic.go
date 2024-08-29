package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatRecordLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetChatRecordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatRecordLogic {
	return &GetChatRecordLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询聊天记录
func (l *GetChatRecordLogic) GetChatRecord(in *chatrpc.IdReq) (*chatrpc.ChatRecordDetails, error) {
	// todo: add your logic here and delete this line

	return &chatrpc.ChatRecordDetails{}, nil
}
