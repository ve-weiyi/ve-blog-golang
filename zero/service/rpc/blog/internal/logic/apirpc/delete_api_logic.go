package apirpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteApiLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteApiLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteApiLogic {
	return &DeleteApiLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除接口
func (l *DeleteApiLogic) DeleteApi(in *blog.IdReq) (*blog.BatchResp, error) {
	rows, err := l.svcCtx.ApiModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	rows2, err := l.svcCtx.ApiModel.DeleteBatch(l.ctx, "parent_id = ? ", in.Id)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: rows + rows2,
	}, nil
}
