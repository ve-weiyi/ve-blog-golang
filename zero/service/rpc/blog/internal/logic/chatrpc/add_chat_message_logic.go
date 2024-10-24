package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddChatMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddChatMessageLogic {
	return &AddChatMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建聊天记录
func (l *AddChatMessageLogic) AddChatMessage(in *chatrpc.ChatMessageNewReq) (*chatrpc.ChatMessageDetails, error) {
	entity := convertChatMessageIn(in)

	_, err := l.svcCtx.TChatMessageModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertChatMessageOut(entity), nil
}

func convertChatMessageIn(in *chatrpc.ChatMessageNewReq) (out *model.TChatMessage) {
	out = &model.TChatMessage{
		Id:          in.Id,
		UserId:      in.UserId,
		DeviceId:    in.DeviceId,
		ChatId:      in.ChatId,
		ReplyMsgId:  in.ReplyMsgId,
		ReplyUsers:  in.ReplyUsers,
		IpAddress:   in.IpAddress,
		IpSource:    in.IpSource,
		ChatContent: in.ChatContent,
		Type:        in.Type,
		Status:      0,
	}

	return out
}

func convertChatMessageOut(in *model.TChatMessage) (out *chatrpc.ChatMessageDetails) {
	out = &chatrpc.ChatMessageDetails{
		Id:          in.Id,
		ChatId:      in.ChatId,
		DeviceId:    in.DeviceId,
		UserId:      in.UserId,
		ReplyMsgId:  in.ReplyMsgId,
		ReplyUsers:  in.ReplyUsers,
		IpAddress:   in.IpAddress,
		IpSource:    in.IpSource,
		ChatContent: in.ChatContent,
		Type:        in.Type,
		Status:      in.Status,
		CreatedAt:   in.CreatedAt.Unix(),
		UpdatedAt:   in.UpdatedAt.Unix(),
	}

	return out
}
