package tag

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/articlerpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTagLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询标签
func NewGetTagLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTagLogic {
	return &GetTagLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTagLogic) GetTag(req *types.IdReq) (resp *types.TagBackDTO, err error) {
	in := &articlerpc.IdReq{
		Id: req.Id,
	}

	category, err := l.svcCtx.ArticleRpc.GetTag(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return ConvertTagTypes(category), nil
}
