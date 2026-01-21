package notice

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/noticerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoticeStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新通知状态
func NewUpdateNoticeStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeStatusLogic {
	return &UpdateNoticeStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNoticeStatusLogic) UpdateNoticeStatus(req *types.UpdateNoticeStatusReq) (resp *types.NoticeBackVO, err error) {
	in := &noticerpc.UpdateNoticeStatusReq{
		Id:            req.Id,
		PublishStatus: req.PublishStatus,
	}

	out, err := l.svcCtx.NoticeRpc.UpdateNoticeStatus(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = convertNoticeOut(out.Notice)
	return resp, nil
}
