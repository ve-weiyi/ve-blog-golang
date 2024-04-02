package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRemarkLogic {
	return &CreateRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRemarkLogic) CreateRemark(req *types.Remark) (resp *types.Remark, err error) {
	// todo: add your logic here and delete this line

	return
}
