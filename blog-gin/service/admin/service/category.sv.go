package service

import (
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/common/request"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/service/admin/dto"
	"github.com/ve-weiyi/ve-blog-golang/blog-gin/svctx"
)

type CategoryService struct {
	svcCtx *svctx.ServiceContext
}

func NewCategoryService(svcCtx *svctx.ServiceContext) *CategoryService {
	return &CategoryService{
		svcCtx: svcCtx,
	}
}

// 创建文章分类
func (s *CategoryService) AddCategory(reqCtx *request.Context, in *dto.CategoryNewReq) (out *dto.CategoryBackVO, err error) {
	// todo

	return
}

// 删除文章分类
func (s *CategoryService) DeletesCategory(reqCtx *request.Context, in *dto.IdsReq) (out *dto.BatchResp, err error) {
	// todo

	return
}

// 分页获取文章分类列表
func (s *CategoryService) FindCategoryList(reqCtx *request.Context, in *dto.CategoryQuery) (out *dto.PageResp, err error) {
	// todo

	return
}

// 更新文章分类
func (s *CategoryService) UpdateCategory(reqCtx *request.Context, in *dto.CategoryNewReq) (out *dto.CategoryBackVO, err error) {
	// todo

	return
}
