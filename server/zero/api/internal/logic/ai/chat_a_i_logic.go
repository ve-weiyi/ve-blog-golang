package ai

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatAILogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatAILogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatAILogic {
	return &ChatAILogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatAILogic) ChatAI(req *types.ChatMessage) (resp *types.ChatResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
