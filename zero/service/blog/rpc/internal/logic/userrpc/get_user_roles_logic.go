package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserRolesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserRolesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserRolesLogic {
	return &GetUserRolesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户角色信息
func (l *GetUserRolesLogic) GetUserRoles(in *blog.UserReq) (*blog.RolePageResp, error) {
	uid := in.UserId

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

	// 查角色
	rs, err := l.svcCtx.RoleModel.FindALL(l.ctx, "id in (?)", ids)
	if err != nil {
		return nil, err
	}

	var list []*blog.RoleDetails
	for _, v := range rs {
		list = append(list, convert.ConvertRoleModelToDetailPb(v))
	}

	out := &blog.RolePageResp{}
	out.Total = int64(len(list))
	out.List = list
	return out, nil
}
