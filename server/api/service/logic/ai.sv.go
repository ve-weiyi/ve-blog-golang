package logic

import (
	"time"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
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
func (l *AIService) ChatAI(reqCtx *request.Context, req *request.ChatMessage) (data *chatgpt.ChatResponse, err error) {
	// 查询用户消息历史记录
	// 查询历史记录
	list, err := l.svcCtx.ChatMessageRepository.FindList(reqCtx, 0, 8, "created_at desc", "chat_id = ?", req.ChatID)
	if err != nil {
		return nil, err
	}

	var msgs []*chatgpt.ChatMessage
	for _, v := range list {
		if v.UserID != -1 {
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

	// 保存用户历史记录
	msg := &entity.ChatMessage{
		ChatID:    req.ChatID,
		UserID:    reqCtx.UID,
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
			ChatID:     req.ChatID,
			UserID:     -1,
			ReplyMsgID: create.ReplyMsgID,
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
func (l *AIService) ChatCos(reqCtx *request.Context, req *request.ChatMessage) (data *chatgpt.ChatResponse, err error) {
	resp, err := chatgpt.NewAIChatGPT().CosRole(req.Content)
	if err != nil {
		return nil, err
	}

	// 保存用户历史记录
	msg := &entity.ChatMessage{
		ChatID:    req.ChatID,
		UserID:    reqCtx.UID,
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
			ChatID:     req.ChatID,
			UserID:     -1,
			ReplyMsgID: create.ReplyMsgID,
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
func (l *AIService) ChatStream(reqCtx *request.Context, req *request.ChatStream) (data *chatgpt.ChatResponse, err error) {

	return l.ChatAI(reqCtx, &request.ChatMessage{
		ChatID:  req.ChatID,
		Content: req.Content,
	})
}

// 和Chatgpt聊天
func (l *AIService) ChatAssistant(reqCtx *request.Context, req *request.ChatMessage) (data *chatgpt.ChatResponse, err error) {
	// 查询历史记录
	list, err := l.svcCtx.ChatMessageRepository.FindList(reqCtx, 1, 3, "created_at desc", "chat_id = ?", req.ChatID)
	if err != nil {
		return nil, err
	}

	var msgs []*chatgpt.ChatMessage
	for _, v := range list {
		if v.UserID == reqCtx.UID {
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
		ChatID:    req.ChatID,
		UserID:    reqCtx.UID,
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

	for _, v := range resp.Choices {
		m := &entity.ChatMessage{
			ChatID:     req.ChatID,
			UserID:     -1,
			ReplyMsgID: create.ReplyMsgID,
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

func (l *AIService) ChatAssistantHistory(reqCtx *request.Context, req *request.ChatHistory) (data []*entity.ChatMessage, err error) {
	if req.Before == 0 {
		req.Before = time.Now().Unix()
	}

	// 查询历史记录
	list, err := l.svcCtx.ChatMessageRepository.FindList(reqCtx, 0, 0, "created_at desc", "chat_id = ? and ? < created_at and created_at < ?", req.ChatID, time.Unix(req.After, 0), time.Unix(req.Before, 0))
	if err != nil {
		return nil, err
	}

	return list, nil
}
