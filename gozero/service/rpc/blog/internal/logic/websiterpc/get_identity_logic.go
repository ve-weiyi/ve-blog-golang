package websiterpclogic

import (
	"context"

	"github.com/mssola/useragent"

	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/rpcutil"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetIdentityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetIdentityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetIdentityLogic {
	return &GetIdentityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取身份标识
func (l *GetIdentityLogic) GetIdentity(in *websiterpc.EmptyReq) (*websiterpc.GetIdentityResp, error) {
	ip, err := rpcutil.GetRPCClientIP(l.ctx)
	if err != nil {
		return nil, err
	}

	ua, err := rpcutil.GetRPCUserAgent(l.ctx)
	if err != nil {
		return nil, err
	}

	browser, _ := useragent.New(ua).Browser()
	os := useragent.New(ua).OS()

	terminal := crypto.Md5v(ip+browser+os, "")

	return &websiterpc.GetIdentityResp{
		TerminalId: terminal,
	}, nil
}
