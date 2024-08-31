// Code generated by goctl. DO NOT EDIT.
// Source: photo.proto

package photorpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/photorpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AlbumDetails       = photorpc.AlbumDetails
	AlbumNewReq        = photorpc.AlbumNewReq
	BannerDetails      = photorpc.BannerDetails
	BannerNewReq       = photorpc.BannerNewReq
	BatchResp          = photorpc.BatchResp
	CountResp          = photorpc.CountResp
	EmptyReq           = photorpc.EmptyReq
	EmptyResp          = photorpc.EmptyResp
	FindAlbumListReq   = photorpc.FindAlbumListReq
	FindAlbumListResp  = photorpc.FindAlbumListResp
	FindBannerListReq  = photorpc.FindBannerListReq
	FindBannerListResp = photorpc.FindBannerListResp
	FindPhotoListReq   = photorpc.FindPhotoListReq
	FindPhotoListResp  = photorpc.FindPhotoListResp
	IdReq              = photorpc.IdReq
	IdsReq             = photorpc.IdsReq
	PhotoDetails       = photorpc.PhotoDetails
	PhotoNewReq        = photorpc.PhotoNewReq
	UserIdReq          = photorpc.UserIdReq

	PhotoRpc interface {
		// 创建照片
		AddPhoto(ctx context.Context, in *PhotoNewReq, opts ...grpc.CallOption) (*PhotoDetails, error)
		// 更新照片
		UpdatePhoto(ctx context.Context, in *PhotoNewReq, opts ...grpc.CallOption) (*PhotoDetails, error)
		// 删除照片
		DeletePhoto(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询照片列表
		FindPhotoList(ctx context.Context, in *FindPhotoListReq, opts ...grpc.CallOption) (*FindPhotoListResp, error)
		// 创建相册
		AddAlbum(ctx context.Context, in *AlbumNewReq, opts ...grpc.CallOption) (*AlbumDetails, error)
		// 更新相册
		UpdateAlbum(ctx context.Context, in *AlbumNewReq, opts ...grpc.CallOption) (*AlbumDetails, error)
		// 获取相册
		GetAlbum(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*AlbumDetails, error)
		// 删除相册
		DeleteAlbum(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询相册列表
		FindAlbumList(ctx context.Context, in *FindAlbumListReq, opts ...grpc.CallOption) (*FindAlbumListResp, error)
		// 创建页面
		AddBanner(ctx context.Context, in *BannerNewReq, opts ...grpc.CallOption) (*BannerDetails, error)
		// 更新页面
		UpdateBanner(ctx context.Context, in *BannerNewReq, opts ...grpc.CallOption) (*BannerDetails, error)
		// 删除页面
		DeleteBanner(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询页面列表
		FindBannerList(ctx context.Context, in *FindBannerListReq, opts ...grpc.CallOption) (*FindBannerListResp, error)
	}

	defaultPhotoRpc struct {
		cli zrpc.Client
	}
)

func NewPhotoRpc(cli zrpc.Client) PhotoRpc {
	return &defaultPhotoRpc{
		cli: cli,
	}
}

// 创建照片
func (m *defaultPhotoRpc) AddPhoto(ctx context.Context, in *PhotoNewReq, opts ...grpc.CallOption) (*PhotoDetails, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.AddPhoto(ctx, in, opts...)
}

// 更新照片
func (m *defaultPhotoRpc) UpdatePhoto(ctx context.Context, in *PhotoNewReq, opts ...grpc.CallOption) (*PhotoDetails, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.UpdatePhoto(ctx, in, opts...)
}

// 删除照片
func (m *defaultPhotoRpc) DeletePhoto(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.DeletePhoto(ctx, in, opts...)
}

// 查询照片列表
func (m *defaultPhotoRpc) FindPhotoList(ctx context.Context, in *FindPhotoListReq, opts ...grpc.CallOption) (*FindPhotoListResp, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.FindPhotoList(ctx, in, opts...)
}

// 创建相册
func (m *defaultPhotoRpc) AddAlbum(ctx context.Context, in *AlbumNewReq, opts ...grpc.CallOption) (*AlbumDetails, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.AddAlbum(ctx, in, opts...)
}

// 更新相册
func (m *defaultPhotoRpc) UpdateAlbum(ctx context.Context, in *AlbumNewReq, opts ...grpc.CallOption) (*AlbumDetails, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.UpdateAlbum(ctx, in, opts...)
}

// 获取相册
func (m *defaultPhotoRpc) GetAlbum(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*AlbumDetails, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.GetAlbum(ctx, in, opts...)
}

// 删除相册
func (m *defaultPhotoRpc) DeleteAlbum(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.DeleteAlbum(ctx, in, opts...)
}

// 查询相册列表
func (m *defaultPhotoRpc) FindAlbumList(ctx context.Context, in *FindAlbumListReq, opts ...grpc.CallOption) (*FindAlbumListResp, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.FindAlbumList(ctx, in, opts...)
}

// 创建页面
func (m *defaultPhotoRpc) AddBanner(ctx context.Context, in *BannerNewReq, opts ...grpc.CallOption) (*BannerDetails, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.AddBanner(ctx, in, opts...)
}

// 更新页面
func (m *defaultPhotoRpc) UpdateBanner(ctx context.Context, in *BannerNewReq, opts ...grpc.CallOption) (*BannerDetails, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.UpdateBanner(ctx, in, opts...)
}

// 删除页面
func (m *defaultPhotoRpc) DeleteBanner(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.DeleteBanner(ctx, in, opts...)
}

// 查询页面列表
func (m *defaultPhotoRpc) FindBannerList(ctx context.Context, in *FindBannerListReq, opts ...grpc.CallOption) (*FindBannerListResp, error) {
	client := photorpc.NewPhotoRpcClient(m.cli.Conn())
	return client.FindBannerList(ctx, in, opts...)
}
