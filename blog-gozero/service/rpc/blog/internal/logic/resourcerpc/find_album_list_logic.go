package resourcerpclogic

import (
	"context"
	"strings"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/resourcerpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindAlbumListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindAlbumListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindAlbumListLogic {
	return &FindAlbumListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询相册列表
func (l *FindAlbumListLogic) FindAlbumList(in *resourcerpc.FindAlbumListReq) (*resourcerpc.FindAlbumListResp, error) {
	page, size, sorts, conditions, params := convertAlbumQuery(in)

	records, total, err := l.svcCtx.TAlbumModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	cm, err := findPhotoCountGroupAlbum(l.ctx, l.svcCtx, records, in.IsDelete)
	if err != nil {
		return nil, err
	}

	var list []*resourcerpc.AlbumDetails
	for _, v := range records {
		list = append(list, convertAlbumOut(v, cm))
	}

	return &resourcerpc.FindAlbumListResp{
		List:  list,
		Total: total,
	}, nil
}

func convertAlbumQuery(in *resourcerpc.FindAlbumListReq) (page int, size int, sorts string, conditions string, params []any) {
	page = int(in.Page)
	size = int(in.PageSize)
	sorts = strings.Join(in.Sorts, ",")
	if sorts == "" {
		sorts = "id desc"
	}

	if in.AlbumName != "" {
		conditions += " album_name like ?"
		params = append(params, "%"+in.AlbumName+"%")
	}

	if conditions != "" {
		conditions += " and "
	}
	conditions += "is_delete = ?"
	params = append(params, in.IsDelete)

	return
}

func findPhotoCountGroupAlbum(ctx context.Context, svcCtx *svc.ServiceContext, list []*model.TAlbum, isDelete int64) (acm map[int64]int, err error) {
	var ids []int64
	for _, v := range list {
		ids = append(ids, v.Id)
	}

	// 查询每个 category_id 的文章数量
	var results []struct {
		AlbumID    int64 `gorm:"column:album_id"`
		PhotoCount int   `gorm:"column:photo_count"`
	}

	err = svcCtx.Gorm.Model(&model.TPhoto{}).
		Select("album_id, COUNT(*) as photo_count").
		Where("album_id IN ? and is_delete = ?", ids, isDelete).
		Group("album_id").
		Order("album_id").
		Scan(&results).Error
	if err != nil {
		return nil, err
	}

	acm = make(map[int64]int)
	for _, result := range results {
		acm[result.AlbumID] = result.PhotoCount
	}

	return acm, nil
}
