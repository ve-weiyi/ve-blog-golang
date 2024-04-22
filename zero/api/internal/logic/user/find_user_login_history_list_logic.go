package user

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLoginHistoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserLoginHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLoginHistoryListLogic {
	return &FindUserLoginHistoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserLoginHistoryListLogic) FindUserLoginHistoryList(req *types.PageQuery) (resp *types.PageResult, err error) {
	logx.Info("FindUserLoginHistoryListLogic", jsonconv.ObjectToJsonIndent(req))
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.UserRpc.FindUserLoginHistoryList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.LoginHistory
	for _, role := range out.List {
		list = append(list, convert.ConvertLoginHistoryTypes(role))
	}

	resp = &types.PageResult{}
	resp.Page = in.Limit.Page
	resp.PageSize = in.Limit.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
