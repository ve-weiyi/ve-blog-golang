package album

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/resourceservice"
)

type CreateAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建相册
func NewCreateAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateAlbumLogic {
	return &CreateAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateAlbumLogic) CreateAlbum(req *types.CreateAlbumReq) (resp *types.AlbumVO, err error) {
	out, err := l.svcCtx.ResourceService.CreateAlbum(l.ctx, &resourceservice.CreateAlbumRequest{
		AlbumName:  req.AlbumName,
		AlbumDesc:  req.AlbumDesc,
		AlbumCover: req.AlbumCover,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.AlbumVO{
		Id:         out.Id,
		AlbumName:  req.AlbumName,
		AlbumDesc:  req.AlbumDesc,
		AlbumCover: req.AlbumCover,
		Status:     req.Status,
		PhotoCount: 0,
	}, nil
}
