package messagerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/messagerpc"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateChatLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateChatLogic {
	return &UpdateChatLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新聊天记录
func (l *UpdateChatLogic) UpdateChat(in *messagerpc.ChatNewReq) (*messagerpc.ChatDetails, error) {
	entity := convertChatIn(in)

	_, err := l.svcCtx.TChatModel.Save(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertChatOut(entity), nil
}
