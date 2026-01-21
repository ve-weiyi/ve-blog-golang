package newsrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/enums"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rpcutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/newsrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMessageLogic {
	return &AddMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建留言
func (l *AddMessageLogic) AddMessage(in *newsrpc.AddMessageReq) (*newsrpc.AddMessageResp, error) {
	uid, _ := rpcutils.GetUserIdFromCtx(l.ctx)
	tid, _ := rpcutils.GetTerminalIdFromCtx(l.ctx)

	entity := &model.TMessage{
		Id:             0,
		UserId:         uid,
		TerminalId:     tid,
		MessageContent: in.MessageContent,
		Status:         enums.MessageStatusNormal,
		//CreatedAt:      time.Time{},
		//UpdatedAt:      time.Time{},
	}

	_, err := l.svcCtx.TMessageModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &newsrpc.AddMessageResp{
		Message: convertMessageOut(entity),
	}, nil
}
