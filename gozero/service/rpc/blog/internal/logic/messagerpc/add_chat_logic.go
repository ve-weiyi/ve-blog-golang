package messagerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/messagerpc"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddChatLogic {
	return &AddChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建聊天记录
func (l *AddChatLogic) AddChat(in *messagerpc.ChatNewReq) (*messagerpc.ChatDetails, error) {
	entity := convertChatIn(in)

	_, err := l.svcCtx.TChatModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertChatOut(entity), nil
}

func convertChatIn(in *messagerpc.ChatNewReq) (out *model.TChat) {
	out = &model.TChat{
		Id:         0,
		UserId:     in.UserId,
		TerminalId: in.TerminalId,
		IpAddress:  in.IpAddress,
		IpSource:   in.IpSource,
		Type:       in.Type,
		Content:    in.Content,
		Status:     0,
	}

	return out
}

func convertChatOut(in *model.TChat) (out *messagerpc.ChatDetails) {
	out = &messagerpc.ChatDetails{
		Id:         in.Id,
		UserId:     in.UserId,
		TerminalId: in.TerminalId,
		IpAddress:  in.IpAddress,
		IpSource:   in.IpSource,
		Type:       in.Type,
		Content:    in.Content,
		Status:     in.Status,
		CreatedAt:  in.CreatedAt.Unix(),
		UpdatedAt:  in.UpdatedAt.Unix(),
	}

	return out
}
