package page

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

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

func (l *FindPageLogic) FindPage(reqCtx *types.RestHeader, req *types.IdReq) (resp *types.Page, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.PageRpc.FindPage(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertPageTypes(out), nil
}
