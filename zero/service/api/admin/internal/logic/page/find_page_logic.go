package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/blogrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindPageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询页面
func NewFindPageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindPageLogic {
	return &FindPageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindPageLogic) FindPage(req *types.IdReq) (resp *types.Page, err error) {
	in := &blogrpc.IdReq{
		Id: req.Id,
	}

	out, err := l.svcCtx.PageRpc.FindPage(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPageTypes(out), nil
}
