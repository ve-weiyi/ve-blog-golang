package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type RemarkService struct {
	svcCtx *svc.ServiceContext
}

func NewRemarkService(svcCtx *svc.ServiceContext) *RemarkService {
	return &RemarkService{
		svcCtx: svcCtx,
	}
}

// 创建Remark记录
func (s *RemarkService) CreateRemark(reqCtx *request.Context, remark *entity.Remark) (data *entity.Remark, err error) {
	return s.svcCtx.RemarkRepository.CreateRemark(reqCtx, remark)
}

// 更新Remark记录
func (s *RemarkService) UpdateRemark(reqCtx *request.Context, remark *entity.Remark) (data *entity.Remark, err error) {
	return s.svcCtx.RemarkRepository.UpdateRemark(reqCtx, remark)
}

// 删除Remark记录
func (s *RemarkService) DeleteRemark(reqCtx *request.Context, id int) (rows int64, err error) {
	return s.svcCtx.RemarkRepository.DeleteRemark(reqCtx, id)
}

// 查询Remark记录
func (s *RemarkService) FindRemark(reqCtx *request.Context, id int) (data *entity.Remark, err error) {
	return s.svcCtx.RemarkRepository.FindRemark(reqCtx, id)
}

// 批量删除Remark记录
func (s *RemarkService) DeleteRemarkByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.RemarkRepository.DeleteRemarkByIds(reqCtx, ids)
}

// 分页获取Remark记录
func (s *RemarkService) FindRemarkList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.Remark, total int64, err error) {
	return s.svcCtx.RemarkRepository.FindRemarkList(reqCtx, page)
}
