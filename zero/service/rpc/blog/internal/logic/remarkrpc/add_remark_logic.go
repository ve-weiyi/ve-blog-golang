package remarkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/remarkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRemarkLogic {
	return &AddRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建留言
func (l *AddRemarkLogic) AddRemark(in *remarkrpc.RemarkNew) (*remarkrpc.RemarkDetails, error) {
	// todo: add your logic here and delete this line

	return &remarkrpc.RemarkDetails{}, nil
}
