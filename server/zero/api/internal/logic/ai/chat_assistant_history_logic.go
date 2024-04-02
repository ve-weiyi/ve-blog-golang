package ai

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatAssistantHistoryLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatAssistantHistoryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatAssistantHistoryLogic {
	return &ChatAssistantHistoryLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatAssistantHistoryLogic) ChatAssistantHistory(req *types.ChatHistory) (resp []*types.ChatMessage, err error) {
	// todo: add your logic here and delete this line

	return
}
