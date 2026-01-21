package noticerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/noticerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletesNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesNoticeLogic {
	return &DeletesNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除通知
func (l *DeletesNoticeLogic) DeletesNotice(in *noticerpc.DeletesNoticeReq) (*noticerpc.DeletesNoticeResp, error) {
	rows, err := l.svcCtx.TSystemNoticeModel.Deletes(l.ctx, "id in ?", in.Ids)
	if err != nil {
		return nil, err
	}

	return &noticerpc.DeletesNoticeResp{
		SuccessCount: rows,
	}, nil
}
