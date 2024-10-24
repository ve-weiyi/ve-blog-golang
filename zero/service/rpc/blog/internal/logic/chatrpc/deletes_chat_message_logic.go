package chatrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/chatrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesChatMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesChatMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesChatMessageLogic {
	return &DeletesChatMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除聊天记录
func (l *DeletesChatMessageLogic) DeletesChatMessage(in *chatrpc.IdsReq) (*chatrpc.BatchResp, error) {
	rows, err := l.svcCtx.TChatMessageModel.Deletes(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &chatrpc.BatchResp{
		SuccessCount: rows,
	}, nil
}
