package resourceservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/app/model"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/pb/resourcerpc"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/internal/svc"
)

type CreateAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAlbumLogic {
	return &CreateAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建相册
func (l *CreateAlbumLogic) CreateAlbum(in *resourcerpc.CreateAlbumRequest) (*resourcerpc.CreateAlbumResponse, error) {
	entity := &model.TAlbum{
		AlbumName:  in.AlbumName,
		AlbumDesc:  in.AlbumDesc,
		AlbumCover: in.AlbumCover,
		Status:     in.Status,
	}

	_, err := l.svcCtx.TAlbumModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return &resourcerpc.CreateAlbumResponse{Id: entity.Id}, nil
}
