// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/pagerpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
)

type PageRpcServer struct {
	svcCtx *svc.ServiceContext
	blog.UnimplementedPageRpcServer
}

func NewPageRpcServer(svcCtx *svc.ServiceContext) *PageRpcServer {
	return &PageRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建页面
func (s *PageRpcServer) AddPage(ctx context.Context, in *blog.Page) (*blog.Page, error) {
	l := pagerpclogic.NewAddPageLogic(ctx, s.svcCtx)
	return l.AddPage(in)
}

// 更新页面
func (s *PageRpcServer) UpdatePage(ctx context.Context, in *blog.Page) (*blog.Page, error) {
	l := pagerpclogic.NewUpdatePageLogic(ctx, s.svcCtx)
	return l.UpdatePage(in)
}

// 删除页面
func (s *PageRpcServer) DeletePage(ctx context.Context, in *blog.IdReq) (*blog.BatchResp, error) {
	l := pagerpclogic.NewDeletePageLogic(ctx, s.svcCtx)
	return l.DeletePage(in)
}

// 批量删除页面
func (s *PageRpcServer) DeletePageList(ctx context.Context, in *blog.IdsReq) (*blog.BatchResp, error) {
	l := pagerpclogic.NewDeletePageListLogic(ctx, s.svcCtx)
	return l.DeletePageList(in)
}

// 查询页面
func (s *PageRpcServer) FindPage(ctx context.Context, in *blog.IdReq) (*blog.Page, error) {
	l := pagerpclogic.NewFindPageLogic(ctx, s.svcCtx)
	return l.FindPage(in)
}

// 查询页面列表
func (s *PageRpcServer) FindPageList(ctx context.Context, in *blog.PageQuery) (*blog.PagePageResp, error) {
	l := pagerpclogic.NewFindPageListLogic(ctx, s.svcCtx)
	return l.FindPageList(in)
}

// 查询页面数量
func (s *PageRpcServer) FindPageCount(ctx context.Context, in *blog.PageQuery) (*blog.CountResp, error) {
	l := pagerpclogic.NewFindPageCountLogic(ctx, s.svcCtx)
	return l.FindPageCount(in)
}
