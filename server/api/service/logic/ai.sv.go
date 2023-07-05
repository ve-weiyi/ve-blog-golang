package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/chatgpt"
)

type AIService struct {
	svcCtx *svc.ServiceContext
}

func NewAIService(svcCtx *svc.ServiceContext) *AIService {
	return &AIService{
		svcCtx: svcCtx,
	}
}

// 和Chatgpt聊天
func (s *AIService) ChatAI(reqCtx *request.Context, req []*chatgpt.ChatMessage) (data *chatgpt.ChatResponse, err error) {

	return chatgpt.NewAIChatGPT().Chat(req)
}

// 和Chatgpt聊天
func (s *AIService) ChatCos(reqCtx *request.Context, act string) (data *chatgpt.ChatResponse, err error) {

	return chatgpt.NewAIChatGPT().CosRole(act)
}
