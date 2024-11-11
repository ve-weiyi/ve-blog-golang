package service

import (
	"github.com/ve-weiyi/ve-blog-golang/gin/api/blog/dto"
	"github.com/ve-weiyi/ve-blog-golang/gin/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/gin/svctx"
)

type ChatService struct {
	svcCtx *svctx.ServiceContext
}

func NewChatService(svcCtx *svctx.ServiceContext) *ChatService {
	return &ChatService{
		svcCtx: svcCtx,
	}
}

// 查询聊天记录
func (s *ChatService) GetChatMessages(reqCtx *request.Context, in *dto.ChatMessageQueryReq) (out *dto.PageResp, err error) {
	// todo

	return
}
