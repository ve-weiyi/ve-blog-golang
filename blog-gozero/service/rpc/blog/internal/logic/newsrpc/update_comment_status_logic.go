package newsrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/newsrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCommentStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCommentStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCommentStatusLogic {
	return &UpdateCommentStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新评论状态
func (l *UpdateCommentStatusLogic) UpdateCommentStatus(in *newsrpc.UpdateCommentStatusReq) (*newsrpc.UpdateCommentStatusResp, error) {
	rows, err := l.svcCtx.TCommentModel.Updates(l.ctx, map[string]interface{}{
		"status": in.Status,
	}, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &newsrpc.UpdateCommentStatusResp{
		SuccessCount: rows,
	}, nil
}
