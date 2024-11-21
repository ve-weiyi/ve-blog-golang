package pagerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/pagerpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdatePageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdatePageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdatePageLogic {
	return &UpdatePageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新页面
func (l *UpdatePageLogic) UpdatePage(in *pagerpc.PageNewReq) (*pagerpc.PageDetails, error) {
	entity := convertPageIn(in)

	_, err := l.svcCtx.TPageModel.Save(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertPageOut(entity), nil
}
