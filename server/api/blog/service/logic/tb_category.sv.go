package logic

import (
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/response"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/service/svc"
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
	return s.svcCtx.CategoryRepository.CreateCategory(category)
}

// 删除Category记录
func (s *CategoryService) DeleteCategory(reqCtx *request.Context, category *entity.Category) (rows int64, err error) {
	return s.svcCtx.CategoryRepository.DeleteCategory(category)
}

// 更新Category记录
func (s *CategoryService) UpdateCategory(reqCtx *request.Context, category *entity.Category) (data *entity.Category, err error) {
	return s.svcCtx.CategoryRepository.UpdateCategory(category)
}

// 根据id获取Category记录
func (s *CategoryService) FindCategory(reqCtx *request.Context, id int) (data *entity.Category, err error) {
	return s.svcCtx.CategoryRepository.FindCategory(id)
}

// 批量删除Category记录
func (s *CategoryService) DeleteCategoryByIds(reqCtx *request.Context, ids []int) (rows int64, err error) {
	return s.svcCtx.CategoryRepository.DeleteCategoryByIds(ids)
}

// 分页获取Category记录
func (s *CategoryService) GetCategoryList(reqCtx *request.Context, page *request.PageInfo) (list []*response.Category, total int64, err error) {
	categorys, total, err := s.svcCtx.CategoryRepository.GetCategoryList(page)
	if err != nil {
		return nil, 0, err
	}

	for _, item := range categorys {

		list = append(list, s.convertCategoryResponse(item))
	}

	return list, total, err
}

func (s *CategoryService) convertCategoryResponse(in *entity.Category) *response.Category {
	_, articleCount, err := s.svcCtx.ArticleRepository.GetArticleListByCategoryId(in.ID)
	if err != nil {
		return nil
	}
	out := &response.Category{
		ID:           in.ID,
		CategoryName: in.CategoryName,
		ArticleCount: articleCount,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}

	return out
}
