package notice

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/client/noticerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddNoticeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建通知
func NewAddNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddNoticeLogic {
	return &AddNoticeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddNoticeLogic) AddNotice(req *types.AddNoticeReq) (resp *types.NoticeBackVO, err error) {
	in := &noticerpc.AddNoticeReq{
		Title:   req.Title,
		Content: req.Content,
		Type:    req.Type,
		Level:   req.Level,
		AppName: req.AppName,
	}

	out, err := l.svcCtx.NoticeRpc.AddNotice(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.NoticeBackVO{
		Id:            out.Notice.Id,
		Title:         out.Notice.Title,
		Content:       out.Notice.Content,
		Type:          out.Notice.Type,
		Level:         out.Notice.Level,
		PublishStatus: out.Notice.PublishStatus,
		AppName:       out.Notice.AppName,
		PublisherId:   out.Notice.PublisherId,
		PublishTime:   out.Notice.PublishTime,
		RevokeTime:    out.Notice.RevokeTime,
		CreatedAt:     out.Notice.CreatedAt,
		UpdatedAt:     out.Notice.UpdatedAt,
	}
	return resp, nil
}
