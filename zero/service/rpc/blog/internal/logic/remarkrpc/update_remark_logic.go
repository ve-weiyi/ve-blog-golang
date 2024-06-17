package remarkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/remarkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRemarkLogic {
	return &UpdateRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新留言
func (l *UpdateRemarkLogic) UpdateRemark(in *remarkrpc.RemarkNewReq) (*remarkrpc.RemarkDetails, error) {
	entity := convertRemarkIn(in)

	_, err := l.svcCtx.RemarkModel.Update(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertRemarkOut(entity), nil
}
