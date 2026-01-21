package notice

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/noticerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletesNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除通知
func NewDeletesNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletesNoticeLogic {
	return &DeletesNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeletesNoticeLogic) DeletesNotice(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &noticerpc.DeletesNoticeReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.NoticeRpc.DeletesNotice(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, nil
}
