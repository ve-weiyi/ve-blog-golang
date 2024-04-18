package menu

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMenuLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMenuLogic {
	return &FindMenuLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindMenuLogic) FindMenu(req *types.IdReq) (resp *types.MenuDetailsDTO, err error) {
	in := convert.ConvertIdReq(req)

	out, err := l.svcCtx.MenuRpc.FindMenu(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return convert.ConvertMenuTypes(out), nil
}
