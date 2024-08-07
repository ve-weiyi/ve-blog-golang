package talkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/talkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTalkLogic {
	return &DeleteTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除说说
func (l *DeleteTalkLogic) DeleteTalk(in *talkrpc.IdReq) (*talkrpc.BatchResp, error) {
	rows, err := l.svcCtx.TalkModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &talkrpc.BatchResp{
		SuccessCount: rows,
	}, nil
}
