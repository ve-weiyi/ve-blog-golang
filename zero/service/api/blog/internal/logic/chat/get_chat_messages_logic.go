package chat

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/chatrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetChatMessagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询聊天记录
func NewGetChatMessagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetChatMessagesLogic {
	return &GetChatMessagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetChatMessagesLogic) GetChatMessages(req *types.ChatMessageQueryReq) (resp *types.PageResp, err error) {
	in := &chatrpc.FindChatMessageListReq{
		After:       req.After,
		Before:      req.Before,
		Limit:       req.Limit,
		UserId:      req.UserId,
		ChatId:      req.ChatId,
		ChatContent: req.Keyword,
		Type:        req.Type,
	}
	out, err := l.svcCtx.ChatRpc.FindChatMessageList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.ChatMessage, 0)
	for _, v := range out.List {
		list = append(list, ConvertChatMessageTypes(v))
	}

	resp = &types.PageResp{}
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}

func ConvertChatMessageTypes(in *chatrpc.ChatMessageDetails) *types.ChatMessage {
	return &types.ChatMessage{
		Id:        in.Id,
		UserId:    cast.ToInt64(in.UserId),
		Content:   in.ChatContent,
		IpAddress: in.IpAddress,
		IpSource:  in.IpSource,
		Type:      in.Type,
		CreatedAt: in.CreatedAt,
		UpdatedAt: in.UpdatedAt,
	}
}
