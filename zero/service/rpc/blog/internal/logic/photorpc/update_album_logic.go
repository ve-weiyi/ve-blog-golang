package photorpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAlbumLogic {
	return &UpdateAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新相册
func (l *UpdateAlbumLogic) UpdateAlbum(in *photorpc.AlbumNew) (*photorpc.AlbumDetails, error) {
	entity := ConvertAlbumIn(in)

	_, err := l.svcCtx.AlbumModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return ConvertAlbumOut(entity), nil
}
