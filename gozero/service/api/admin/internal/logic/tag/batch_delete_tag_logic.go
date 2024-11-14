package tag

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/client/articlerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除标签
func NewBatchDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteTagLogic {
	return &BatchDeleteTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BatchDeleteTagLogic) BatchDeleteTag(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := &articlerpc.IdsReq{
		Ids: req.Ids,
	}

	out, err := l.svcCtx.ArticleRpc.DeleteTag(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}
	return resp, nil
}
