package tagrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagListLogic {
	return &DeleteTagListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量删除标签
func (l *DeleteTagListLogic) DeleteTagList(in *blog.IdsReq) (*blog.BatchResp, error) {
	result, err := l.svcCtx.TagModel.DeleteBatch(l.ctx, "id in (?)", in.Ids)
	if err != nil {
		return nil, err
	}

	return &blog.BatchResp{
		SuccessCount: result,
	}, nil
}
