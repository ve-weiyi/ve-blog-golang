package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建留言
func NewAddRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRemarkLogic {
	return &AddRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRemarkLogic) AddRemark(req *types.Remark) (resp *types.Remark, err error) {
	// todo: add your logic here and delete this line

	return
}
