// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/logic/blogrpc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/pb/blog"
)

type BlogRpcServer struct {
	svcCtx *svc.ServiceContext
	blog.UnimplementedBlogRpcServer
}

func NewBlogRpcServer(svcCtx *svc.ServiceContext) *BlogRpcServer {
	return &BlogRpcServer{
		svcCtx: svcCtx,
	}
}

// 连接检测
func (s *BlogRpcServer) Ping(ctx context.Context, in *blog.PingReq) (*blog.PingResp, error) {
	l := blogrpclogic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
