// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.3
// Source: resource.proto

package resourcerpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/resourcerpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchResp              = resourcerpc.BatchResp
	EmptyReq               = resourcerpc.EmptyReq
	EmptyResp              = resourcerpc.EmptyResp
	FileFolderDetails      = resourcerpc.FileFolderDetails
	FileFolderNewReq       = resourcerpc.FileFolderNewReq
	FileUploadDetails      = resourcerpc.FileUploadDetails
	FileUploadNewReq       = resourcerpc.FileUploadNewReq
	FindFileFolderListReq  = resourcerpc.FindFileFolderListReq
	FindFileFolderListResp = resourcerpc.FindFileFolderListResp
	FindFileUploadListReq  = resourcerpc.FindFileUploadListReq
	FindFileUploadListResp = resourcerpc.FindFileUploadListResp
	IdReq                  = resourcerpc.IdReq
	IdsReq                 = resourcerpc.IdsReq
	UserIdReq              = resourcerpc.UserIdReq

	ResourceRpc interface {
		// 创建文件夹
		AddFileFolder(ctx context.Context, in *FileFolderNewReq, opts ...grpc.CallOption) (*FileFolderDetails, error)
		// 更新文件夹
		UpdateFileFolder(ctx context.Context, in *FileFolderNewReq, opts ...grpc.CallOption) (*FileFolderDetails, error)
		// 删除文件夹
		DeleteFileFolder(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询文件夹列表
		FindFileFolderList(ctx context.Context, in *FindFileFolderListReq, opts ...grpc.CallOption) (*FindFileFolderListResp, error)
		// 创建文件上传
		AddFileUpload(ctx context.Context, in *FileUploadNewReq, opts ...grpc.CallOption) (*FileUploadDetails, error)
		// 更新文件上传
		UpdateFileUpload(ctx context.Context, in *FileUploadNewReq, opts ...grpc.CallOption) (*FileUploadDetails, error)
		// 删除文件上传
		DeleteFileUpload(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询文件上传列表
		FindFileUploadList(ctx context.Context, in *FindFileUploadListReq, opts ...grpc.CallOption) (*FindFileUploadListResp, error)
	}

	defaultResourceRpc struct {
		cli zrpc.Client
	}
)

func NewResourceRpc(cli zrpc.Client) ResourceRpc {
	return &defaultResourceRpc{
		cli: cli,
	}
}

// 创建文件夹
func (m *defaultResourceRpc) AddFileFolder(ctx context.Context, in *FileFolderNewReq, opts ...grpc.CallOption) (*FileFolderDetails, error) {
	client := resourcerpc.NewResourceRpcClient(m.cli.Conn())
	return client.AddFileFolder(ctx, in, opts...)
}

// 更新文件夹
func (m *defaultResourceRpc) UpdateFileFolder(ctx context.Context, in *FileFolderNewReq, opts ...grpc.CallOption) (*FileFolderDetails, error) {
	client := resourcerpc.NewResourceRpcClient(m.cli.Conn())
	return client.UpdateFileFolder(ctx, in, opts...)
}

// 删除文件夹
func (m *defaultResourceRpc) DeleteFileFolder(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := resourcerpc.NewResourceRpcClient(m.cli.Conn())
	return client.DeleteFileFolder(ctx, in, opts...)
}

// 查询文件夹列表
func (m *defaultResourceRpc) FindFileFolderList(ctx context.Context, in *FindFileFolderListReq, opts ...grpc.CallOption) (*FindFileFolderListResp, error) {
	client := resourcerpc.NewResourceRpcClient(m.cli.Conn())
	return client.FindFileFolderList(ctx, in, opts...)
}

// 创建文件上传
func (m *defaultResourceRpc) AddFileUpload(ctx context.Context, in *FileUploadNewReq, opts ...grpc.CallOption) (*FileUploadDetails, error) {
	client := resourcerpc.NewResourceRpcClient(m.cli.Conn())
	return client.AddFileUpload(ctx, in, opts...)
}

// 更新文件上传
func (m *defaultResourceRpc) UpdateFileUpload(ctx context.Context, in *FileUploadNewReq, opts ...grpc.CallOption) (*FileUploadDetails, error) {
	client := resourcerpc.NewResourceRpcClient(m.cli.Conn())
	return client.UpdateFileUpload(ctx, in, opts...)
}

// 删除文件上传
func (m *defaultResourceRpc) DeleteFileUpload(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := resourcerpc.NewResourceRpcClient(m.cli.Conn())
	return client.DeleteFileUpload(ctx, in, opts...)
}

// 查询文件上传列表
func (m *defaultResourceRpc) FindFileUploadList(ctx context.Context, in *FindFileUploadListReq, opts ...grpc.CallOption) (*FindFileUploadListResp, error) {
	client := resourcerpc.NewResourceRpcClient(m.cli.Conn())
	return client.FindFileUploadList(ctx, in, opts...)
}
