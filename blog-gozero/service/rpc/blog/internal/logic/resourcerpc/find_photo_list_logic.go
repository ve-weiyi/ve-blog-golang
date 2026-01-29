package resourcerpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/infra/queryx"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPhotoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindPhotoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPhotoListLogic {
	return &FindPhotoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页获取照片列表
func (l *FindPhotoListLogic) FindPhotoList(in *resourcerpc.FindPhotoListReq) (*resourcerpc.FindPhotoListResp, error) {
	page, size, sorts, conditions, params := convertPhotoQuery(in)

	records, total, err := l.svcCtx.TPhotoModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*resourcerpc.Photo
	for _, v := range records {
		list = append(list, convertPhotoOut(v))
	}

	return &resourcerpc.FindPhotoListResp{
		List: list,
		Pagination: &resourcerpc.PageResp{
			Page:     int64(page),
			PageSize: int64(size),
			Total:    total,
		},
	}, nil
}

func convertPhotoQuery(in *resourcerpc.FindPhotoListReq) (page int, size int, sorts string, conditions string, params []any) {
	var opts []queryx.Option
	if in.Paginate != nil {
		opts = append(opts, queryx.WithPage(int(in.Paginate.Page)))
		opts = append(opts, queryx.WithSize(int(in.Paginate.PageSize)))
		opts = append(opts, queryx.WithSorts(in.Paginate.Sorts...))
	}

	if in.IsDelete >= 0 {
		opts = append(opts, queryx.WithCondition("is_delete = ?", in.IsDelete))
	}

	if in.AlbumId != 0 {
		opts = append(opts, queryx.WithCondition("album_id = ?", in.AlbumId))
	}

	return queryx.NewQueryBuilder(opts...).Build()
}
