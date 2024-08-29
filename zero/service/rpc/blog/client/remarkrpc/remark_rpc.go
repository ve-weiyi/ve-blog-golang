// Code generated by goctl. DO NOT EDIT.
// Source: remark.proto

package remarkrpc

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/remarkrpc"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	BatchResp          = remarkrpc.BatchResp
	CountResp          = remarkrpc.CountResp
	EmptyReq           = remarkrpc.EmptyReq
	EmptyResp          = remarkrpc.EmptyResp
	FindRemarkListReq  = remarkrpc.FindRemarkListReq
	FindRemarkListResp = remarkrpc.FindRemarkListResp
	IdReq              = remarkrpc.IdReq
	IdsReq             = remarkrpc.IdsReq
	RemarkDetails      = remarkrpc.RemarkDetails
	RemarkNew          = remarkrpc.RemarkNew
	UserIdReq          = remarkrpc.UserIdReq

	RemarkRpc interface {
		// 创建留言
		AddRemark(ctx context.Context, in *RemarkNew, opts ...grpc.CallOption) (*RemarkDetails, error)
		// 更新留言
		UpdateRemark(ctx context.Context, in *RemarkNew, opts ...grpc.CallOption) (*RemarkDetails, error)
		// 删除留言
		DeleteRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 批量删除留言
		DeleteRemarkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error)
		// 查询留言
		GetRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RemarkDetails, error)
		// 查询留言列表
		FindRemarkList(ctx context.Context, in *FindRemarkListReq, opts ...grpc.CallOption) (*FindRemarkListResp, error)
	}

	defaultRemarkRpc struct {
		cli zrpc.Client
	}
)

func NewRemarkRpc(cli zrpc.Client) RemarkRpc {
	return &defaultRemarkRpc{
		cli: cli,
	}
}

// 创建留言
func (m *defaultRemarkRpc) AddRemark(ctx context.Context, in *RemarkNew, opts ...grpc.CallOption) (*RemarkDetails, error) {
	client := remarkrpc.NewRemarkRpcClient(m.cli.Conn())
	return client.AddRemark(ctx, in, opts...)
}

// 更新留言
func (m *defaultRemarkRpc) UpdateRemark(ctx context.Context, in *RemarkNew, opts ...grpc.CallOption) (*RemarkDetails, error) {
	client := remarkrpc.NewRemarkRpcClient(m.cli.Conn())
	return client.UpdateRemark(ctx, in, opts...)
}

// 删除留言
func (m *defaultRemarkRpc) DeleteRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := remarkrpc.NewRemarkRpcClient(m.cli.Conn())
	return client.DeleteRemark(ctx, in, opts...)
}

// 批量删除留言
func (m *defaultRemarkRpc) DeleteRemarkList(ctx context.Context, in *IdsReq, opts ...grpc.CallOption) (*BatchResp, error) {
	client := remarkrpc.NewRemarkRpcClient(m.cli.Conn())
	return client.DeleteRemarkList(ctx, in, opts...)
}

// 查询留言
func (m *defaultRemarkRpc) GetRemark(ctx context.Context, in *IdReq, opts ...grpc.CallOption) (*RemarkDetails, error) {
	client := remarkrpc.NewRemarkRpcClient(m.cli.Conn())
	return client.GetRemark(ctx, in, opts...)
}

// 查询留言列表
func (m *defaultRemarkRpc) FindRemarkList(ctx context.Context, in *FindRemarkListReq, opts ...grpc.CallOption) (*FindRemarkListResp, error) {
	client := remarkrpc.NewRemarkRpcClient(m.cli.Conn())
	return client.FindRemarkList(ctx, in, opts...)
}
