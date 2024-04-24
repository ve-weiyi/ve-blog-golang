package tag

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTagListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 批量删除标签
func NewDeleteTagListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTagListLogic {
	return &DeleteTagListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteTagListLogic) DeleteTagList(req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := convert.ConvertIdsReq(req)

	out, err := l.svcCtx.TagRpc.DeleteTagList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{
		SuccessCount: out.SuccessCount,
	}, nil
}
