package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建留言
func NewCreateRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRemarkLogic {
	return &CreateRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRemarkLogic) CreateRemark(req *types.Remark) (resp *types.Remark, err error) {
	in := convert.ConvertRemarkPb(req)
	out, err := l.svcCtx.RemarkRpc.CreateRemark(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = convert.ConvertRemarkTypes(out)
	return resp, nil
}
