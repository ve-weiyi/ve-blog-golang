package remarkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/remarkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRemarkLogic {
	return &FindRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询留言
func (l *FindRemarkLogic) FindRemark(in *remarkrpc.IdReq) (*remarkrpc.RemarkDetails, error) {
	// todo: add your logic here and delete this line

	return &remarkrpc.RemarkDetails{}, nil
}
