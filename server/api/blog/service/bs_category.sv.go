package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/svc"
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
func (l *CategoryService) CreateCategory(reqCtx *request.Context, category *entity.Category) (data *entity.Category, err error) {
	return l.svcCtx.CategoryRepository.Create(reqCtx, category)
}

// 更新Category记录
func (l *CategoryService) UpdateCategory(reqCtx *request.Context, category *entity.Category) (data *entity.Category, err error) {
	return l.svcCtx.CategoryRepository.Update(reqCtx, category)
}

// 删除Category记录
func (l *CategoryService) DeleteCategory(reqCtx *request.Context, req *request.IdReq) (rows int64, err error) {
	return l.svcCtx.CategoryRepository.Delete(reqCtx, "id = ?", req.Id)
}

// 查询Category记录
func (l *CategoryService) FindCategory(reqCtx *request.Context, req *request.IdReq) (data *entity.Category, err error) {
	return l.svcCtx.CategoryRepository.First(reqCtx, "id = ?", req.Id)
}

// 批量删除Category记录
func (l *CategoryService) DeleteCategoryList(reqCtx *request.Context, req *request.IdsReq) (rows int64, err error) {
	return l.svcCtx.CategoryRepository.Delete(reqCtx, "id in (?)", req.Ids)
}

// 分页获取Category记录
func (l *CategoryService) FindCategoryList(reqCtx *request.Context, page *request.PageQuery) (list []*entity.Category, total int64, err error) {
	p, s := page.PageClause()
	cond, args := page.ConditionClause()
	order := page.OrderClause()

	list, err = l.svcCtx.CategoryRepository.FindList(reqCtx, p, s, order, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	total, err = l.svcCtx.CategoryRepository.Count(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}
