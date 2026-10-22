package mediaservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/mediarpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/model"
)

type GetAlbumLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAlbumLogic {
	return &GetAlbumLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询相册详情
func (l *GetAlbumLogic) GetAlbum(in *mediarpc.GetAlbumRequest) (*mediarpc.GetAlbumResponse, error) {
	conditions := "`id` = ?"
	args := []any{in.Id}
	if in.PublicOnly {
		conditions += " AND `status` = ? AND `is_delete` = ?"
		args = append(args, enums.AlbumStatusPublic, 0)
	}
	entity, err := l.svcCtx.TAlbumModel.FindOne(l.ctx, "", conditions, args...)
	if err != nil {
		return nil, err
	}

	cm := findPhotoCountGroupAlbum(l.ctx, l.svcCtx.GormDB, []*model.TAlbum{entity})

	return &mediarpc.GetAlbumResponse{
		Album: convertAlbumOut(entity, cm),
	}, nil
}
