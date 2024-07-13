package account

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

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

	in := &blog.PageQuery{}
	// 查询用户数量
	users, err := l.svcCtx.UserRpc.FindUserList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	// 分类
	AreaMap := make(map[string]int64)
	for _, item := range users.List {
		key := item.IpSource
		if _, ok := AreaMap[key]; ok {
			AreaMap[key]++
		} else {
			AreaMap[key] = 1
		}
	}

	var list []*types.UserArea
	for k, v := range AreaMap {
		list = append(list, &types.UserArea{
			Name:  k,
			Value: v,
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
