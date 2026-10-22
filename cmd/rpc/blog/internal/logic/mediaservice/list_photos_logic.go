package mediaservicelogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/pb/mediarpc"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/internal/svc"
	"github.com/ve-weiyi/blog-cloud/infra/domain/enums"
	"github.com/ve-weiyi/blog-cloud/infra/store/queryx"
)

type ListPhotosLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPhotosLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPhotosLogic {
	return &ListPhotosLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListPhotosLogic) ListPhotos(in *mediarpc.ListPhotosRequest) (*mediarpc.ListPhotosResponse, error) {
	if in.PublicOnly {
		// 公开照片列表必须先确认所属相册可见，否则私密相册的照片会泄漏
		if _, err := l.svcCtx.TAlbumModel.FindOne(l.ctx, "",
			"`id` = ? AND `status` = ? AND `is_delete` = ?",
			in.AlbumId, enums.AlbumStatusPublic, 0); err != nil {
			return nil, err
		}
		zero := int64(0)
		in.IsDelete = &zero
	}

	var opts []queryx.Option
	if in.ListQuery != nil {
		opts = append(opts, queryx.WithPage(int(in.ListQuery.Page)))
		opts = append(opts, queryx.WithSize(int(in.ListQuery.PageSize)))
		opts = append(opts, queryx.WithSorts(in.ListQuery.Sorts...))
	}
	if in.AlbumId != 0 {
		opts = append(opts, queryx.WithCondition("album_id = ?", in.AlbumId))
	}
	if in.IsDelete != nil {
		opts = append(opts, queryx.WithCondition("is_delete = ?", *in.IsDelete))
	}

	page, size, sorts, conditions, params := queryx.NewQueryBuilder(opts...).Build()
	records, total, err := l.svcCtx.TPhotoModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*mediarpc.Photo
	for _, v := range records {
		list = append(list, convertPhotoOut(v))
	}

	return &mediarpc.ListPhotosResponse{
		ListResult: &mediarpc.ListResult{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
		List: list,
	}, nil
}
