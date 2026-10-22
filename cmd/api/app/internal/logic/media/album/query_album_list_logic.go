package album

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/app/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/mediaservice"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
)

type QueryAlbumListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取相册列表
func NewQueryAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QueryAlbumListLogic {
	return &QueryAlbumListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QueryAlbumListLogic) QueryAlbumList(req *types.QueryAlbumListReq) (resp *types.ListResult, err error) {
	isDelete := int64(0)
	status := int64(enums.AlbumStatusPublic)

	in := &mediaservice.ListAlbumsRequest{
		ListQuery: &mediaservice.ListQuery{
			Page:     req.Page,
			PageSize: req.PageSize,
			Sorts:    req.Sorts,
		},
		IsDelete: &isDelete,
		Status:   &status,
	}

	out, err := l.svcCtx.MediaService.ListAlbums(l.ctx, in)
	if err != nil {
		return nil, err
	}

	list := make([]*types.Album, 0)
	for _, v := range out.List {
		list = append(list, &types.Album{
			Id:         v.Id,
			AlbumName:  v.AlbumName,
			AlbumDesc:  v.AlbumDesc,
			AlbumCover: v.AlbumCover,
		})
	}

	resp = &types.ListResult{
		Page:     out.ListResult.Page,
		PageSize: out.ListResult.PageSize,
		Total:    out.ListResult.Total,
		List:     list,
	}
	return
}
