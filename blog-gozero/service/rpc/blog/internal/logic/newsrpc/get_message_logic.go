package newsrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/newsrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询留言
func (l *GetMessageLogic) GetMessage(in *newsrpc.GetMessageReq) (*newsrpc.GetMessageResp, error) {
	entity, err := l.svcCtx.TMessageModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &newsrpc.GetMessageResp{
		Message: convertMessageOut(entity),
	}, nil
}
