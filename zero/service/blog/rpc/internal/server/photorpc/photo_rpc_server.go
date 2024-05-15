// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/logic/photorpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"
)

type PhotoRpcServer struct {
	svcCtx *svc.ServiceContext
	blog.UnimplementedPhotoRpcServer
}

func NewPhotoRpcServer(svcCtx *svc.ServiceContext) *PhotoRpcServer {
	return &PhotoRpcServer{
		svcCtx: svcCtx,
	}
}

// 创建照片
func (s *PhotoRpcServer) CreatePhoto(ctx context.Context, in *blog.Photo) (*blog.Photo, error) {
	l := photorpclogic.NewCreatePhotoLogic(ctx, s.svcCtx)
	return l.CreatePhoto(in)
}

// 更新照片
func (s *PhotoRpcServer) UpdatePhoto(ctx context.Context, in *blog.Photo) (*blog.Photo, error) {
	l := photorpclogic.NewUpdatePhotoLogic(ctx, s.svcCtx)
	return l.UpdatePhoto(in)
}

// 删除照片
func (s *PhotoRpcServer) DeletePhoto(ctx context.Context, in *blog.IdReq) (*blog.BatchResp, error) {
	l := photorpclogic.NewDeletePhotoLogic(ctx, s.svcCtx)
	return l.DeletePhoto(in)
}

// 批量删除照片
func (s *PhotoRpcServer) DeletePhotoList(ctx context.Context, in *blog.IdsReq) (*blog.BatchResp, error) {
	l := photorpclogic.NewDeletePhotoListLogic(ctx, s.svcCtx)
	return l.DeletePhotoList(in)
}

// 查询照片
func (s *PhotoRpcServer) FindPhoto(ctx context.Context, in *blog.IdReq) (*blog.Photo, error) {
	l := photorpclogic.NewFindPhotoLogic(ctx, s.svcCtx)
	return l.FindPhoto(in)
}

// 分页获取照片列表
func (s *PhotoRpcServer) FindPhotoList(ctx context.Context, in *blog.PageQuery) (*blog.PhotoPageResp, error) {
	l := photorpclogic.NewFindPhotoListLogic(ctx, s.svcCtx)
	return l.FindPhotoList(in)
}

// 查询照片数量
func (s *PhotoRpcServer) FindPhotoCount(ctx context.Context, in *blog.PageQuery) (*blog.CountResp, error) {
	l := photorpclogic.NewFindPhotoCountLogic(ctx, s.svcCtx)
	return l.FindPhotoCount(in)
}

// 创建相册
func (s *PhotoRpcServer) CreatePhotoAlbum(ctx context.Context, in *blog.PhotoAlbum) (*blog.PhotoAlbum, error) {
	l := photorpclogic.NewCreatePhotoAlbumLogic(ctx, s.svcCtx)
	return l.CreatePhotoAlbum(in)
}

// 更新相册
func (s *PhotoRpcServer) UpdatePhotoAlbum(ctx context.Context, in *blog.PhotoAlbum) (*blog.PhotoAlbum, error) {
	l := photorpclogic.NewUpdatePhotoAlbumLogic(ctx, s.svcCtx)
	return l.UpdatePhotoAlbum(in)
}

// 删除相册
func (s *PhotoRpcServer) DeletePhotoAlbum(ctx context.Context, in *blog.IdReq) (*blog.BatchResp, error) {
	l := photorpclogic.NewDeletePhotoAlbumLogic(ctx, s.svcCtx)
	return l.DeletePhotoAlbum(in)
}

// 批量删除相册
func (s *PhotoRpcServer) DeletePhotoAlbumList(ctx context.Context, in *blog.IdsReq) (*blog.BatchResp, error) {
	l := photorpclogic.NewDeletePhotoAlbumListLogic(ctx, s.svcCtx)
	return l.DeletePhotoAlbumList(in)
}

// 查询相册
func (s *PhotoRpcServer) FindPhotoAlbum(ctx context.Context, in *blog.IdReq) (*blog.PhotoAlbum, error) {
	l := photorpclogic.NewFindPhotoAlbumLogic(ctx, s.svcCtx)
	return l.FindPhotoAlbum(in)
}

// 分页获取相册列表
func (s *PhotoRpcServer) FindPhotoAlbumList(ctx context.Context, in *blog.PageQuery) (*blog.PhotoAlbumPageResp, error) {
	l := photorpclogic.NewFindPhotoAlbumListLogic(ctx, s.svcCtx)
	return l.FindPhotoAlbumList(in)
}

// 查询相册数量
func (s *PhotoRpcServer) FindPhotoAlbumCount(ctx context.Context, in *blog.PageQuery) (*blog.CountResp, error) {
	l := photorpclogic.NewFindPhotoAlbumCountLogic(ctx, s.svcCtx)
	return l.FindPhotoAlbumCount(in)
}