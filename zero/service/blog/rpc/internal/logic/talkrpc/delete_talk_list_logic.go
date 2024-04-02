package talkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTalkListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTalkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTalkListLogic {
	return &DeleteTalkListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除说说
func (l *DeleteTalkListLogic) DeleteTalkList(in *blog.IdsReq) (*blog.BatchResp, error) {
	result, err := l.svcCtx.TalkModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: result,
	}, nil
}
