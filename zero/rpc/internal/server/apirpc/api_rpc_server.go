// Code generated by goctl. DO NOT EDIT.
// Source: account.proto

package server

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/logic/apirpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
)

type ApiRpcServer struct {
	svcCtx *svc.ServiceContext
	account.UnimplementedApiRpcServer
}

func NewApiRpcServer(svcCtx *svc.ServiceContext) *ApiRpcServer {
	return &ApiRpcServer{
		svcCtx: svcCtx,
	}
}

func (s *ApiRpcServer) CreateApi(ctx context.Context, in *account.Api) (*account.Api, error) {
	l := apirpclogic.NewCreateApiLogic(ctx, s.svcCtx)
	return l.CreateApi(in)
}

func (s *ApiRpcServer) UpdateApi(ctx context.Context, in *account.Api) (*account.Api, error) {
	l := apirpclogic.NewUpdateApiLogic(ctx, s.svcCtx)
	return l.UpdateApi(in)
}

func (s *ApiRpcServer) DeleteApi(ctx context.Context, in *account.IdReq) (*account.EmptyResp, error) {
	l := apirpclogic.NewDeleteApiLogic(ctx, s.svcCtx)
	return l.DeleteApi(in)
}

func (s *ApiRpcServer) FindApi(ctx context.Context, in *account.IdReq) (*account.Api, error) {
	l := apirpclogic.NewFindApiLogic(ctx, s.svcCtx)
	return l.FindApi(in)
}

func (s *ApiRpcServer) DeleteListApi(ctx context.Context, in *account.IdsReq) (*account.BatchResult, error) {
	l := apirpclogic.NewDeleteListApiLogic(ctx, s.svcCtx)
	return l.DeleteListApi(in)
}

func (s *ApiRpcServer) FindListApi(ctx context.Context, in *account.PageQuery) (*account.PageResult, error) {
	l := apirpclogic.NewFindListApiLogic(ctx, s.svcCtx)
	return l.FindListApi(in)
}
