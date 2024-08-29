// 声明 proto 语法版本，固定值

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.19.4
// source: photo.proto

// proto 包名

package photorpc

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PhotoRpc_AddPhoto_FullMethodName         = "/photorpc.PhotoRpc/AddPhoto"
	PhotoRpc_UpdatePhoto_FullMethodName      = "/photorpc.PhotoRpc/UpdatePhoto"
	PhotoRpc_DeletePhoto_FullMethodName      = "/photorpc.PhotoRpc/DeletePhoto"
	PhotoRpc_DeletePhotoList_FullMethodName  = "/photorpc.PhotoRpc/DeletePhotoList"
	PhotoRpc_FindPhotoList_FullMethodName    = "/photorpc.PhotoRpc/FindPhotoList"
	PhotoRpc_AddAlbum_FullMethodName         = "/photorpc.PhotoRpc/AddAlbum"
	PhotoRpc_UpdateAlbum_FullMethodName      = "/photorpc.PhotoRpc/UpdateAlbum"
	PhotoRpc_GetAlbum_FullMethodName         = "/photorpc.PhotoRpc/GetAlbum"
	PhotoRpc_DeleteAlbum_FullMethodName      = "/photorpc.PhotoRpc/DeleteAlbum"
	PhotoRpc_DeleteAlbumList_FullMethodName  = "/photorpc.PhotoRpc/DeleteAlbumList"
	PhotoRpc_FindAlbumList_FullMethodName    = "/photorpc.PhotoRpc/FindAlbumList"
	PhotoRpc_AddBanner_FullMethodName        = "/photorpc.PhotoRpc/AddBanner"
	PhotoRpc_UpdateBanner_FullMethodName     = "/photorpc.PhotoRpc/UpdateBanner"
	PhotoRpc_DeleteBanner_FullMethodName     = "/photorpc.PhotoRpc/DeleteBanner"
	PhotoRpc_DeleteBannerList_FullMethodName = "/photorpc.PhotoRpc/DeleteBannerList"
	PhotoRpc_FindBannerList_FullMethodName   = "/photorpc.PhotoRpc/FindBannerList"
)

// PhotoRpcClient is the client API for PhotoRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PhotoRpcClient interface {
	// 创建照片
	AddPhoto(ctx context.Context, in *PhotoNew, opts ...grpc.CallOption) (*PhotoDetails, error)
	// 更新照片
	UpdatePhoto(ctx context.Context, in *PhotoNew, opts ...grpc.CallOption) (*PhotoDetails, error)
	// 删除照片
	DeletePhoto(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 批量删除照片
	DeletePhotoList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询照片列表
	FindPhotoList(ctx context.Context, in *FindPhotoListReq, opts ...grpc.CallOption) (*FindPhotoListResp, error)
	// 创建相册
	AddAlbum(ctx context.Context, in *AlbumNew, opts ...grpc.CallOption) (*AlbumDetails, error)
	// 更新相册
	UpdateAlbum(ctx context.Context, in *AlbumNew, opts ...grpc.CallOption) (*AlbumDetails, error)
	// 获取相册
	GetAlbum(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*AlbumDetails, error)
	// 删除相册
	DeleteAlbum(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 批量删除相册
	DeleteAlbumList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询相册列表
	FindAlbumList(ctx context.Context, in *FindAlbumListReq, opts ...grpc.CallOption) (*FindAlbumListResp, error)
	// 创建页面
	AddBanner(ctx context.Context, in *BannerNew, opts ...grpc.CallOption) (*BannerDetails, error)
	// 更新页面
	UpdateBanner(ctx context.Context, in *BannerNew, opts ...grpc.CallOption) (*BannerDetails, error)
	// 删除页面
	DeleteBanner(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 批量删除页面
	DeleteBannerList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
	// 查询页面列表
	FindBannerList(ctx context.Context, in *FindBannerListReq, opts ...grpc.CallOption) (*FindBannerListResp, error)
}

type photoRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewPhotoRpcClient(cc grpc.ClientConnInterface) PhotoRpcClient {
	return &photoRpcClient{cc}
}

func (c *photoRpcClient) AddPhoto(ctx context.Context, in *PhotoNew, opts ...grpc.CallOption) (*PhotoDetails, error) {
	out := new(PhotoDetails)
	err := c.cc.Invoke(ctx, PhotoRpc_AddPhoto_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) UpdatePhoto(ctx context.Context, in *PhotoNew, opts ...grpc.CallOption) (*PhotoDetails, error) {
	out := new(PhotoDetails)
	err := c.cc.Invoke(ctx, PhotoRpc_UpdatePhoto_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) DeletePhoto(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, PhotoRpc_DeletePhoto_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) DeletePhotoList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, PhotoRpc_DeletePhotoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) FindPhotoList(ctx context.Context, in *FindPhotoListReq, opts ...grpc.CallOption) (*FindPhotoListResp, error) {
	out := new(FindPhotoListResp)
	err := c.cc.Invoke(ctx, PhotoRpc_FindPhotoList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) AddAlbum(ctx context.Context, in *AlbumNew, opts ...grpc.CallOption) (*AlbumDetails, error) {
	out := new(AlbumDetails)
	err := c.cc.Invoke(ctx, PhotoRpc_AddAlbum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) UpdateAlbum(ctx context.Context, in *AlbumNew, opts ...grpc.CallOption) (*AlbumDetails, error) {
	out := new(AlbumDetails)
	err := c.cc.Invoke(ctx, PhotoRpc_UpdateAlbum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) GetAlbum(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*AlbumDetails, error) {
	out := new(AlbumDetails)
	err := c.cc.Invoke(ctx, PhotoRpc_GetAlbum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) DeleteAlbum(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, PhotoRpc_DeleteAlbum_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) DeleteAlbumList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, PhotoRpc_DeleteAlbumList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) FindAlbumList(ctx context.Context, in *FindAlbumListReq, opts ...grpc.CallOption) (*FindAlbumListResp, error) {
	out := new(FindAlbumListResp)
	err := c.cc.Invoke(ctx, PhotoRpc_FindAlbumList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) AddBanner(ctx context.Context, in *BannerNew, opts ...grpc.CallOption) (*BannerDetails, error) {
	out := new(BannerDetails)
	err := c.cc.Invoke(ctx, PhotoRpc_AddBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) UpdateBanner(ctx context.Context, in *BannerNew, opts ...grpc.CallOption) (*BannerDetails, error) {
	out := new(BannerDetails)
	err := c.cc.Invoke(ctx, PhotoRpc_UpdateBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) DeleteBanner(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, PhotoRpc_DeleteBanner_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) DeleteBannerList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	out := new(BatchResp)
	err := c.cc.Invoke(ctx, PhotoRpc_DeleteBannerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *photoRpcClient) FindBannerList(ctx context.Context, in *FindBannerListReq, opts ...grpc.CallOption) (*FindBannerListResp, error) {
	out := new(FindBannerListResp)
	err := c.cc.Invoke(ctx, PhotoRpc_FindBannerList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PhotoRpcServer is the server API for PhotoRpc service.
// All implementations must embed UnimplementedPhotoRpcServer
// for forward compatibility
type PhotoRpcServer interface {
	// 创建照片
	AddPhoto(context.Context, *PhotoNew) (*PhotoDetails, error)
	// 更新照片
	UpdatePhoto(context.Context, *PhotoNew) (*PhotoDetails, error)
	// 删除照片
	DeletePhoto(context.Context, *IdReq) (*BatchResp, error)
	// 批量删除照片
	DeletePhotoList(context.Context, *IdsReq) (*BatchResp, error)
	// 查询照片列表
	FindPhotoList(context.Context, *FindPhotoListReq) (*FindPhotoListResp, error)
	// 创建相册
	AddAlbum(context.Context, *AlbumNew) (*AlbumDetails, error)
	// 更新相册
	UpdateAlbum(context.Context, *AlbumNew) (*AlbumDetails, error)
	// 获取相册
	GetAlbum(context.Context, *IdReq) (*AlbumDetails, error)
	// 删除相册
	DeleteAlbum(context.Context, *IdReq) (*BatchResp, error)
	// 批量删除相册
	DeleteAlbumList(context.Context, *IdsReq) (*BatchResp, error)
	// 查询相册列表
	FindAlbumList(context.Context, *FindAlbumListReq) (*FindAlbumListResp, error)
	// 创建页面
	AddBanner(context.Context, *BannerNew) (*BannerDetails, error)
	// 更新页面
	UpdateBanner(context.Context, *BannerNew) (*BannerDetails, error)
	// 删除页面
	DeleteBanner(context.Context, *IdReq) (*BatchResp, error)
	// 批量删除页面
	DeleteBannerList(context.Context, *IdsReq) (*BatchResp, error)
	// 查询页面列表
	FindBannerList(context.Context, *FindBannerListReq) (*FindBannerListResp, error)
	mustEmbedUnimplementedPhotoRpcServer()
}

// UnimplementedPhotoRpcServer must be embedded to have forward compatible implementations.
type UnimplementedPhotoRpcServer struct {
}

func (UnimplementedPhotoRpcServer) AddPhoto(context.Context, *PhotoNew) (*PhotoDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPhoto not implemented")
}
func (UnimplementedPhotoRpcServer) UpdatePhoto(context.Context, *PhotoNew) (*PhotoDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePhoto not implemented")
}
func (UnimplementedPhotoRpcServer) DeletePhoto(context.Context, *IdReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePhoto not implemented")
}
func (UnimplementedPhotoRpcServer) DeletePhotoList(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePhotoList not implemented")
}
func (UnimplementedPhotoRpcServer) FindPhotoList(context.Context, *FindPhotoListReq) (*FindPhotoListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindPhotoList not implemented")
}
func (UnimplementedPhotoRpcServer) AddAlbum(context.Context, *AlbumNew) (*AlbumDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAlbum not implemented")
}
func (UnimplementedPhotoRpcServer) UpdateAlbum(context.Context, *AlbumNew) (*AlbumDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAlbum not implemented")
}
func (UnimplementedPhotoRpcServer) GetAlbum(context.Context, *IdReq) (*AlbumDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAlbum not implemented")
}
func (UnimplementedPhotoRpcServer) DeleteAlbum(context.Context, *IdReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlbum not implemented")
}
func (UnimplementedPhotoRpcServer) DeleteAlbumList(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAlbumList not implemented")
}
func (UnimplementedPhotoRpcServer) FindAlbumList(context.Context, *FindAlbumListReq) (*FindAlbumListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindAlbumList not implemented")
}
func (UnimplementedPhotoRpcServer) AddBanner(context.Context, *BannerNew) (*BannerDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddBanner not implemented")
}
func (UnimplementedPhotoRpcServer) UpdateBanner(context.Context, *BannerNew) (*BannerDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBanner not implemented")
}
func (UnimplementedPhotoRpcServer) DeleteBanner(context.Context, *IdReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBanner not implemented")
}
func (UnimplementedPhotoRpcServer) DeleteBannerList(context.Context, *IdsReq) (*BatchResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBannerList not implemented")
}
func (UnimplementedPhotoRpcServer) FindBannerList(context.Context, *FindBannerListReq) (*FindBannerListResp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindBannerList not implemented")
}
func (UnimplementedPhotoRpcServer) mustEmbedUnimplementedPhotoRpcServer() {}

// UnsafePhotoRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PhotoRpcServer will
// result in compilation errors.
type UnsafePhotoRpcServer interface {
	mustEmbedUnimplementedPhotoRpcServer()
}

func RegisterPhotoRpcServer(s grpc.ServiceRegistrar, srv PhotoRpcServer) {
	s.RegisterService(&PhotoRpc_ServiceDesc, srv)
}

func _PhotoRpc_AddPhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhotoNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).AddPhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_AddPhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).AddPhoto(ctx, req.(*PhotoNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_UpdatePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PhotoNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).UpdatePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_UpdatePhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).UpdatePhoto(ctx, req.(*PhotoNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_DeletePhoto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).DeletePhoto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_DeletePhoto_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).DeletePhoto(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_DeletePhotoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).DeletePhotoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_DeletePhotoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).DeletePhotoList(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_FindPhotoList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindPhotoListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).FindPhotoList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_FindPhotoList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).FindPhotoList(ctx, req.(*FindPhotoListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_AddAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).AddAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_AddAlbum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).AddAlbum(ctx, req.(*AlbumNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_UpdateAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AlbumNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).UpdateAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_UpdateAlbum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).UpdateAlbum(ctx, req.(*AlbumNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_GetAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).GetAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_GetAlbum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).GetAlbum(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_DeleteAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).DeleteAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_DeleteAlbum_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).DeleteAlbum(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_DeleteAlbumList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).DeleteAlbumList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_DeleteAlbumList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).DeleteAlbumList(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_FindAlbumList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindAlbumListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).FindAlbumList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_FindAlbumList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).FindAlbumList(ctx, req.(*FindAlbumListReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_AddBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).AddBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_AddBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).AddBanner(ctx, req.(*BannerNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_UpdateBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BannerNew)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).UpdateBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_UpdateBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).UpdateBanner(ctx, req.(*BannerNew))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_DeleteBanner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).DeleteBanner(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_DeleteBanner_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).DeleteBanner(ctx, req.(*IdReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_DeleteBannerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdsReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).DeleteBannerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_DeleteBannerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).DeleteBannerList(ctx, req.(*IdsReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _PhotoRpc_FindBannerList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindBannerListReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PhotoRpcServer).FindBannerList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PhotoRpc_FindBannerList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PhotoRpcServer).FindBannerList(ctx, req.(*FindBannerListReq))
	}
	return interceptor(ctx, in, info, handler)
}

// PhotoRpc_ServiceDesc is the grpc.ServiceDesc for PhotoRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PhotoRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "photorpc.PhotoRpc",
	HandlerType: (*PhotoRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddPhoto",
			Handler:    _PhotoRpc_AddPhoto_Handler,
		},
		{
			MethodName: "UpdatePhoto",
			Handler:    _PhotoRpc_UpdatePhoto_Handler,
		},
		{
			MethodName: "DeletePhoto",
			Handler:    _PhotoRpc_DeletePhoto_Handler,
		},
		{
			MethodName: "DeletePhotoList",
			Handler:    _PhotoRpc_DeletePhotoList_Handler,
		},
		{
			MethodName: "FindPhotoList",
			Handler:    _PhotoRpc_FindPhotoList_Handler,
		},
		{
			MethodName: "AddAlbum",
			Handler:    _PhotoRpc_AddAlbum_Handler,
		},
		{
			MethodName: "UpdateAlbum",
			Handler:    _PhotoRpc_UpdateAlbum_Handler,
		},
		{
			MethodName: "GetAlbum",
			Handler:    _PhotoRpc_GetAlbum_Handler,
		},
		{
			MethodName: "DeleteAlbum",
			Handler:    _PhotoRpc_DeleteAlbum_Handler,
		},
		{
			MethodName: "DeleteAlbumList",
			Handler:    _PhotoRpc_DeleteAlbumList_Handler,
		},
		{
			MethodName: "FindAlbumList",
			Handler:    _PhotoRpc_FindAlbumList_Handler,
		},
		{
			MethodName: "AddBanner",
			Handler:    _PhotoRpc_AddBanner_Handler,
		},
		{
			MethodName: "UpdateBanner",
			Handler:    _PhotoRpc_UpdateBanner_Handler,
		},
		{
			MethodName: "DeleteBanner",
			Handler:    _PhotoRpc_DeleteBanner_Handler,
		},
		{
			MethodName: "DeleteBannerList",
			Handler:    _PhotoRpc_DeleteBannerList_Handler,
		},
		{
			MethodName: "FindBannerList",
			Handler:    _PhotoRpc_FindBannerList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "photo.proto",
}
