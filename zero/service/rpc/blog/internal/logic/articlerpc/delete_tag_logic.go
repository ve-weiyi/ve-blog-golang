package tagrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagLogic {
	return &DeleteTagLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除标签
func (l *DeleteTagLogic) DeleteTag(in *blog.IdReq) (*blog.BatchResp, error) {
	rows, err := l.svcCtx.TagModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: rows,
	}, nil
}
