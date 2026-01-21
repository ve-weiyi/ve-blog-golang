package noticerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/noticerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateNoticeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateNoticeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateNoticeLogic {
	return &UpdateNoticeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新通知
func (l *UpdateNoticeLogic) UpdateNotice(in *noticerpc.UpdateNoticeReq) (*noticerpc.UpdateNoticeResp, error) {
	entity, err := l.svcCtx.TSystemNoticeModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	entity.Title = in.Title
	entity.Content = in.Content
	entity.Type = in.Type
	entity.Level = in.Level
	entity.AppName = in.AppName

	_, err = l.svcCtx.TSystemNoticeModel.Save(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &noticerpc.UpdateNoticeResp{
		Notice: convertNoticeOut(entity),
	}, nil
}
