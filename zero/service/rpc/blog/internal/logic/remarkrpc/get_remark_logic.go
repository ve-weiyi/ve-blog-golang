package remarkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/remarkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRemarkLogic {
	return &GetRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询留言
func (l *GetRemarkLogic) GetRemark(in *remarkrpc.IdReq) (*remarkrpc.RemarkDetails, error) {
	entity, err := l.svcCtx.RemarkModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return ConvertRemarkOut(entity), nil
}
