package mediaservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/mediarpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizcode"
	"github.com/ve-weiyi/blog-cloud/infra/errors/bizerr"
	"github.com/ve-weiyi/blog-cloud/model"
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
func (l *UpdateAlbumLogic) UpdateAlbum(in *mediarpc.UpdateAlbumRequest) (*mediarpc.UpdateAlbumResponse, error) {
	if !enums.IsValidAlbumStatus(in.Status) {
		return nil, bizerr.NewBizError(bizcode.CodeParamValueInvalid, "相册状态不合法")
	}

	entity, err := l.svcCtx.TAlbumModel.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	entity.AlbumName = in.AlbumName
	entity.AlbumDesc = in.AlbumDesc
	entity.AlbumCover = in.AlbumCover
	entity.Status = in.Status

	_, err = l.svcCtx.TAlbumModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	cm := findPhotoCountGroupAlbum(l.ctx, l.svcCtx.GormDB, []*model.TAlbum{entity})

	return &mediarpc.UpdateAlbumResponse{
		Album: convertAlbumOut(entity, cm),
	}, nil
}
