package message

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/newsrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建留言
func NewAddMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMessageLogic {
	return &AddMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddMessageLogic) AddMessage(req *types.NewMessageReq) (resp *types.EmptyResp, err error) {
	in := &newsrpc.AddMessageReq{
		MessageContent: req.MessageContent,
	}

	_, err = l.svcCtx.NewsRpc.AddMessage(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
