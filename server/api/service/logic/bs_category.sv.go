package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
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
	return s.svcCtx.CategoryRepository.Create(reqCtx, category)
}

// 更新Category记录
func (s *CategoryService) UpdateCategory(reqCtx *request.Context, category *entity.Category) (data *entity.Category, err error) {
	return s.svcCtx.CategoryRepository.Update(reqCtx, category)
}

// 删除Category记录
func (s *CategoryService) DeleteCategory(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return s.svcCtx.CategoryRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Category记录
func (s *CategoryService) FindCategory(reqCtx *request.Context, req *request.IdReq) (data *entity.Category, err error) {
	return s.svcCtx.CategoryRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Category记录
func (s *CategoryService) DeleteCategoryList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return s.svcCtx.CategoryRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Category记录
func (s *CategoryService) FindCategoryList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Category, total int64, err error) {
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = s.svcCtx.CategoryRepository.FindList(reqCtx, page.Limit.Page, page.Limit.PageSize, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = s.svcCtx.CategoryRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
