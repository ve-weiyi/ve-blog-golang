package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type CategoryService struct {
	svcCtx *svc.ServiceContext
}

func NewCategoryService(svcCtx *svc.ServiceContext) *CategoryService {
	return &CategoryService{
		svcCtx: svcCtx,
	}
}

// 创建Category记录
func (s *CategoryService) CreateCategory(reqCtx *request.Context, category *entity.Category) (data *entity.Category, err error) {
	return s.svcCtx.CategoryRepository.CreateCategory(reqCtx, category)
}

// 更新Category记录
func (s *CategoryService) UpdateCategory(reqCtx *request.Context, category *entity.Category) (data *entity.Category, err error) {
	return s.svcCtx.CategoryRepository.UpdateCategory(reqCtx, category)
}

// 删除Category记录
func (s *CategoryService) DeleteCategory(reqCtx *request.Context, id int) (rows int64, err error) {
	return s.svcCtx.CategoryRepository.DeleteCategory(reqCtx, id)
}

// 查询Category记录
func (s *CategoryService) FindCategory(reqCtx *request.Context, id int) (data *entity.Category, err error) {
	return s.svcCtx.CategoryRepository.FindCategory(reqCtx, id)
}

// 批量删除Category记录
func (s *CategoryService) DeleteCategoryByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.CategoryRepository.DeleteCategoryByIds(reqCtx, ids)
}

// 分页获取Category记录
func (s *CategoryService) FindCategoryList(reqCtx *request.Context, page *request.PageInfo) (list []*entity.Category, total int64, err error) {
	return s.svcCtx.CategoryRepository.FindCategoryList(reqCtx, page)
}
