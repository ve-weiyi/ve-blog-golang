package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeletePhotoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeletePhotoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeletePhotoLogic {
	return &DeletePhotoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除照片
func (l *DeletePhotoLogic) DeletePhoto(in *photorpc.IdsReq) (*photorpc.BatchResp, error) {
	rows, err := l.svcCtx.TPhotoModel.Deletes(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &photorpc.BatchResp{
		SuccessCount: rows,
	}, nil
}
