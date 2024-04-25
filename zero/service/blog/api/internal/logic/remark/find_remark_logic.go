package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询留言
func NewFindRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindRemarkLogic {
	return &FindRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindRemarkLogic) FindRemark(req *types.IdReq) (resp *types.Remark, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.RemarkRpc.FindRemark(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertRemarkTypes(out), nil
}
