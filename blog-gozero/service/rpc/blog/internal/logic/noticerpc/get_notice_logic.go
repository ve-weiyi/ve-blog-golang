package noticerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/noticerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetNoticeLogic {
	return &GetNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询通知
func (l *GetNoticeLogic) GetNotice(in *noticerpc.GetNoticeReq) (*noticerpc.GetNoticeResp, error) {
	entity, err := l.svcCtx.TSystemNoticeModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &noticerpc.GetNoticeResp{
		Notice: convertNoticeOut(entity),
	}, nil
}
