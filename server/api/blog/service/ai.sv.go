package service

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/chatgpt"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
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
func (l *AIService) ChatAI(reqCtx *request.Context, req *dto.ChatMessage) (data *chatgpt.ChatResponse, err error) {
	// 查询历史记录
	list, err := l.svcCtx.ChatMessageRepository.FindList(reqCtx, 1, 3, "created_at desc", "chat_id = ?", req.ChatId)
	if err != nil {
		return nil, err
	}

	var msgs []*chatgpt.ChatMessage
	for _, v := range list {
		if v.UserId == reqCtx.Uid {
			msgs = append(msgs, &chatgpt.ChatMessage{
				Role:    chatgpt.RoleUser,
				Content: v.Content,
			})
		} else {
			msgs = append(msgs, &chatgpt.ChatMessage{
				Role:    chatgpt.RoleAI,
				Content: v.Content,
			})
		}
	}

	msgs = append(msgs, &chatgpt.ChatMessage{
		Role:    chatgpt.RoleUser,
		Content: req.Content,
	})

	resp, err := chatgpt.NewAIChatGPT().Chat(msgs)
	if err != nil {
		return nil, err
	}

	// 保存历史记录
	msg := &entity.ChatMessage{
		ChatId:    req.ChatId,
		UserId:    reqCtx.Uid,
		Content:   req.Content,
		IpAddress: reqCtx.IpAddress,
		//IpSource:  reqCtx.GetIpSource(),
		Type:   0,
		Status: 0,
	}

	create, err := l.svcCtx.ChatMessageRepository.Create(reqCtx, msg)
	if err != nil {
		return nil, err
	}

	// 保存ai历史记录
	for _, v := range resp.Choices {
		m := &entity.ChatMessage{
			ChatId:     req.ChatId,
			UserId:     -1,
			ReplyMsgId: create.ReplyMsgId,
			Content:    v.Message.Content,
			Type:       1,
			Status:     0,
		}

		_, err = l.svcCtx.ChatMessageRepository.Create(reqCtx, m)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// 和Chatgpt聊天
func (l *AIService) ChatCos(reqCtx *request.Context, req *dto.ChatMessage) (data *chatgpt.ChatResponse, err error) {
	resp, err := chatgpt.NewAIChatGPT().CosRole(req.Content)
	if err != nil {
		return nil, err
	}

	// 保存用户历史记录
	msg := &entity.ChatMessage{
		ChatId:    req.ChatId,
		UserId:    reqCtx.Uid,
		Content:   req.Content,
		IpAddress: reqCtx.IpAddress,
		//IpSource:  reqCtx.GetIpSource(),
		Type:   0,
		Status: 0,
	}

	create, err := l.svcCtx.ChatMessageRepository.Create(reqCtx, msg)
	if err != nil {
		return nil, err
	}

	// 保存ai历史记录
	for _, v := range resp.Choices {
		m := &entity.ChatMessage{
			ChatId:     req.ChatId,
			UserId:     -1,
			ReplyMsgId: create.ReplyMsgId,
			Content:    v.Message.Content,
			Type:       1,
			Status:     0,
		}

		_, err = l.svcCtx.ChatMessageRepository.Create(reqCtx, m)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// 和Chatgpt聊天
func (l *AIService) ChatStream(reqCtx *request.Context, req *dto.ChatStream) (data *chatgpt.ChatResponse, err error) {

	return l.ChatAI(reqCtx, &dto.ChatMessage{
		ChatId:  req.ChatId,
		Content: req.Content,
	})
}

func (l *AIService) ChatAssistantHistory(reqCtx *request.Context, req *dto.ChatHistory) (data []*entity.ChatMessage, err error) {
	if req.Before == 0 {
		req.Before = time.Now().Unix()
	}

	// 查询历史记录
	list, err := l.svcCtx.ChatMessageRepository.FindList(reqCtx, 0, 0, "created_at desc", "chat_id = ? and ? < created_at and created_at < ?", req.ChatId, time.Unix(req.After, 0), time.Unix(req.Before, 0))
	if err != nil {
		return nil, err
	}

	return list, nil
}
