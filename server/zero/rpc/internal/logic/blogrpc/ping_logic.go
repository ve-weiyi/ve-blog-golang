package blogrpclogic

import (
	"context"
	"fmt"

	"google.golang.org/grpc/metadata"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 连接检测
func (l *PingLogic) Ping(in *blog.PingReq) (*blog.PingResp, error) {

	md, ok := metadata.FromOutgoingContext(l.ctx)
	//md, ok := metadata.FromIncomingContext(l.ctx)
	//if !ok {
	//	return nil, fmt.Errorf("metadata error")
	//}

	fmt.Println("ping", ok, md)

	return &blog.PingResp{
		Pong: "pong",
	}, nil
}
