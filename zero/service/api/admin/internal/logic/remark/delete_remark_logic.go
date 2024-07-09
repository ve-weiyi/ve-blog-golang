package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除留言
func NewDeleteRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteRemarkLogic {
	return &DeleteRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteRemarkLogic) DeleteRemark(req *types.IdReq) (resp *types.BatchResp, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.RemarkRpc.DeleteRemark(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}, nil
}
