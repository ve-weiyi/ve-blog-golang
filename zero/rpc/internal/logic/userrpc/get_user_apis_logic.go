package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/rpcutils"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserApisLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserApisLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserApisLogic {
	return &GetUserApisLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户接口权限
func (l *GetUserApisLogic) GetUserApis(in *account.EmptyReq) (*account.ApiPageResp, error) {
	uid, err := rpcutils.GetRPCInnerXUserId(l.ctx)
	if err != nil {
		return nil, err
	}

	// 查用户
	//ua, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", uid)
	//if err != nil {
	//	return nil, err
	//}

	// 查用户角色
	urs, err := l.svcCtx.UserRoleModel.FindALL(l.ctx, "user_id = ?", uid)
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, v := range urs {
		ids = append(ids, v.RoleId)
	}

	// 查角色拥有的接口
	rs, err := l.svcCtx.RoleApiModel.FindALL(l.ctx, "id in (?)", ids)
	if err != nil {
		return nil, err
	}

	var apiIds []int64
	for _, v := range rs {
		apiIds = append(apiIds, v.ApiId)
	}

	// 查接口信息
	apis, err := l.svcCtx.ApiModel.FindALL(l.ctx, "id in (?)", apiIds)
	if err != nil {
		return nil, err
	}

	var list []*account.ApiDetailsDTO
	for _, v := range apis {
		list = append(list, convert.ConvertApiModelToDetailPb(v))
	}

	out := &account.ApiPageResp{}
	out.Total = int64(len(list))
	out.List = list
	return out, nil
}
