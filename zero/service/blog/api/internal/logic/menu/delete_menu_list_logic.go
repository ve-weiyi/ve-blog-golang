package menu

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMenuListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteMenuListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMenuListLogic {
	return &DeleteMenuListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMenuListLogic) DeleteMenuList(reqCtx *types.RestHeader, req *types.IdsReq) (resp *types.BatchResp, err error) {
	in := convert.ConvertIdsReq(req)

	_, err = l.svcCtx.MenuRpc.DeleteMenuList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return &types.BatchResp{}, nil
}
