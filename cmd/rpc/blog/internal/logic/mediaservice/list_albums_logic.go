package mediaservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/mediarpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/store/queryx"
)

type ListAlbumsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAlbumsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAlbumsLogic {
	return &ListAlbumsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListAlbumsLogic) ListAlbums(in *mediarpc.ListAlbumsRequest) (*mediarpc.ListAlbumsResponse, error) {
	var opts []queryx.Option
	if in.ListQuery != nil {
		opts = append(opts, queryx.WithPage(int(in.ListQuery.Page)))
		opts = append(opts, queryx.WithSize(int(in.ListQuery.PageSize)))
		opts = append(opts, queryx.WithSorts(in.ListQuery.Sorts...))
	}
	if in.IsDelete != nil {
		opts = append(opts, queryx.WithCondition("is_delete = ?", *in.IsDelete))
	}
	if in.AlbumName != nil {
		opts = append(opts, queryx.WithCondition("album_name like ?", "%"+*in.AlbumName+"%"))
	}
	if in.Status != nil {
		opts = append(opts, queryx.WithCondition("status = ?", *in.Status))
	}

	page, size, sorts, conditions, params := queryx.NewQueryBuilder(opts...).Build()
	records, total, err := l.svcCtx.TAlbumModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	cm := findPhotoCountGroupAlbum(l.ctx, l.svcCtx.GormDB, records)

	var list []*mediarpc.Album
	for _, v := range records {
		list = append(list, convertAlbumOut(v, cm))
	}

	return &mediarpc.ListAlbumsResponse{
		ListResult: &mediarpc.ListResult{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
		List: list,
	}, nil
}
