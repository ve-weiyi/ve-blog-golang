package notice

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/noticerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新通知
func NewUpdateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeLogic {
	return &UpdateNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateNoticeLogic) UpdateNotice(req *types.UpdateNoticeReq) (resp *types.NoticeBackVO, err error) {
	in := &noticerpc.UpdateNoticeReq{
		Id:      req.Id,
		Title:   req.Title,
		Content: req.Content,
		Type:    req.Type,
		Level:   req.Level,
		AppName: req.AppName,
	}

	out, err := l.svcCtx.NoticeRpc.UpdateNotice(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = convertNoticeOut(out.Notice)
	return resp, nil
}
