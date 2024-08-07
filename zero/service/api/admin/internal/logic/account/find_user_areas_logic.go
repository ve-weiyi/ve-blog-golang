package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserAreasLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserAreasLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserAreasLogic {
	return &FindUserAreasLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserAreasLogic) FindUserAreas(req *types.PageQuery) (resp *types.PageResp, err error) {

	in := &accountrpc.EmptyReq{}
	// 查询用户数量
	out, err := l.svcCtx.AccountRpc.FindUserRegionList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.UserArea
	for _, v := range out.List {
		list = append(list, &types.UserArea{
			Name:  v.Region,
			Value: v.Count,
		})
	}

	resp = &types.PageResp{
		Page:     0,
		PageSize: 0,
		Total:    0,
		List:     list,
	}

	return
}
