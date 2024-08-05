// Code generated by goctl. DO NOT EDIT.
// Source: blog.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/logic/blogrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
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

// 上传文件
func (s *BlogRpcServer) GetUserVisitList(ctx context.Context, in *blog.EmptyReq) (*blog.UserVisitPageRsp, error) {
	l := blogrpclogic.NewGetUserVisitListLogic(ctx, s.svcCtx)
	return l.GetUserVisitList(in)
}
