// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/categoryrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
)

type CategoryRpcServer struct {
	svcCtx *svc.ServiceContext
	blog.UnimplementedCategoryRpcServer
}

func NewCategoryRpcServer(svcCtx *svc.ServiceContext) *CategoryRpcServer {
	return &CategoryRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建文章分类
func (s *CategoryRpcServer) AddCategory(ctx context.Context, in *blog.Category) (*blog.Category, error) {
	l := categoryrpclogic.NewAddCategoryLogic(ctx, s.svcCtx)
	return l.AddCategory(in)
}

// 更新文章分类
func (s *CategoryRpcServer) UpdateCategory(ctx context.Context, in *blog.Category) (*blog.Category, error) {
	l := categoryrpclogic.NewUpdateCategoryLogic(ctx, s.svcCtx)
	return l.UpdateCategory(in)
}

// 删除文章分类
func (s *CategoryRpcServer) DeleteCategory(ctx context.Context, in *blog.IdReq) (*blog.BatchResp, error) {
	l := categoryrpclogic.NewDeleteCategoryLogic(ctx, s.svcCtx)
	return l.DeleteCategory(in)
}

// 批量删除文章分类
func (s *CategoryRpcServer) DeleteCategoryList(ctx context.Context, in *blog.IdsReq) (*blog.BatchResp, error) {
	l := categoryrpclogic.NewDeleteCategoryListLogic(ctx, s.svcCtx)
	return l.DeleteCategoryList(in)
}

// 查询文章分类
func (s *CategoryRpcServer) FindCategory(ctx context.Context, in *blog.IdReq) (*blog.Category, error) {
	l := categoryrpclogic.NewFindCategoryLogic(ctx, s.svcCtx)
	return l.FindCategory(in)
}

// 查询文章分类数量
func (s *CategoryRpcServer) FindCategoryCount(ctx context.Context, in *blog.PageQuery) (*blog.CountResp, error) {
	l := categoryrpclogic.NewFindCategoryCountLogic(ctx, s.svcCtx)
	return l.FindCategoryCount(in)
}

// 查询文章分类列表
func (s *CategoryRpcServer) FindCategoryList(ctx context.Context, in *blog.PageQuery) (*blog.CategoryPageResp, error) {
	l := categoryrpclogic.NewFindCategoryListLogic(ctx, s.svcCtx)
	return l.FindCategoryList(in)
}

// 查询文章分类列表(通过ids)
func (s *CategoryRpcServer) FindCategoryListByIds(ctx context.Context, in *blog.IdsReq) (*blog.CategoryPageResp, error) {
	l := categoryrpclogic.NewFindCategoryListByIdsLogic(ctx, s.svcCtx)
	return l.FindCategoryListByIds(in)
}
