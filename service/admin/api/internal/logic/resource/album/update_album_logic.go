package album

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/resourceservice"
)

type UpdateAlbumLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新相册
func NewUpdateAlbumLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAlbumLogic {
	return &UpdateAlbumLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateAlbumLogic) UpdateAlbum(req *types.UpdateAlbumReq) (resp *types.EmptyResp, err error) {
	_, err = l.svcCtx.ResourceService.UpdateAlbum(l.ctx, &resourceservice.UpdateAlbumRequest{
		Id:         req.Id,
		AlbumName:  req.AlbumName,
		AlbumDesc:  req.AlbumDesc,
		AlbumCover: req.AlbumCover,
		Status:     req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
