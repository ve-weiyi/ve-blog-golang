package newsrpclogic

import (
	"context"
	"fmt"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rpcutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/newsrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMessageLogic {
	return &UpdateMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新留言
func (l *UpdateMessageLogic) UpdateMessage(in *newsrpc.UpdateMessageReq) (*newsrpc.UpdateMessageResp, error) {
	uid, err := rpcutils.GetUserIdFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	message, err := l.svcCtx.TMessageModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	if message.UserId != uid {
		return nil, fmt.Errorf("无权限操作")
	}

	message.MessageContent = in.MessageContent
	message.Status = in.Status
	_, err = l.svcCtx.TMessageModel.Save(l.ctx, message)
	if err != nil {
		return nil, err
	}

	return &newsrpc.UpdateMessageResp{
		Message: convertMessageOut(message),
	}, nil
}
