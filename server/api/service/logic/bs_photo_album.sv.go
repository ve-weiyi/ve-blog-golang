package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type PhotoAlbumService struct {
	svcCtx *svc.ServiceContext
}

func NewPhotoAlbumService(svcCtx *svc.ServiceContext) *PhotoAlbumService {
	return &PhotoAlbumService{
		svcCtx: svcCtx,
	}
}

// 创建PhotoAlbum记录
func (l *PhotoAlbumService) CreatePhotoAlbum(reqCtx *request.Context, photoAlbum *entity.PhotoAlbum) (data *entity.PhotoAlbum, err error) {
	return l.svcCtx.PhotoAlbumRepository.Create(reqCtx, photoAlbum)
}

// 更新PhotoAlbum记录
func (l *PhotoAlbumService) UpdatePhotoAlbum(reqCtx *request.Context, photoAlbum *entity.PhotoAlbum) (data *entity.PhotoAlbum, err error) {
	return l.svcCtx.PhotoAlbumRepository.Update(reqCtx, photoAlbum)
}

// 删除PhotoAlbum记录
func (l *PhotoAlbumService) DeletePhotoAlbum(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.PhotoAlbumRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询PhotoAlbum记录
func (l *PhotoAlbumService) FindPhotoAlbum(reqCtx *request.Context, req *request.IdReq) (data *entity.PhotoAlbum, err error) {
	return l.svcCtx.PhotoAlbumRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除PhotoAlbum记录
func (l *PhotoAlbumService) DeletePhotoAlbumList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.PhotoAlbumRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取PhotoAlbum记录
func (l *PhotoAlbumService) FindPhotoAlbumList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.PhotoAlbum, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.PhotoAlbumRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.PhotoAlbumRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
