package accountrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserOnlineListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserOnlineListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserOnlineListLogic {
	return &FindUserOnlineListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找在线用户列表
func (l *FindUserOnlineListLogic) FindUserOnlineList(in *accountrpc.FindUserListReq) (*accountrpc.FindUserInfoListResp, error) {
	total, err := l.svcCtx.OnlineUserService.GetOnlineUserCount(l.ctx)
	if err != nil {
		return nil, err
	}

	offset := (in.Page - 1) * in.PageSize
	limit := in.PageSize

	// 查找在线用户
	uids, err := l.svcCtx.OnlineUserService.GetOnlineUsers(l.ctx, offset, limit)
	if err != nil {
		return nil, err
	}

	users, err := l.svcCtx.TUserModel.FindALL(l.ctx, "user_id in (?)", uids)
	if err != nil {
		return nil, err
	}

	// 查找用户角色
	urList, err := l.svcCtx.TUserRoleModel.FindALL(l.ctx, "user_id in (?)", uids)
	if err != nil {
		return nil, err
	}

	var ursMap = make(map[string][]int64)
	var roleIds []int64
	for _, item := range urList {
		roleIds = append(roleIds, item.RoleId)
		ursMap[item.UserId] = append(ursMap[item.UserId], item.RoleId)
	}

	// 查找角色信息
	rList, err := l.svcCtx.TRoleModel.FindALL(l.ctx, "id in (?)", roleIds)
	if err != nil {
		return nil, err
	}

	var list []*accountrpc.UserInfoResp
	for _, item := range users {

		var roles []*model.TRole
		ur, _ := ursMap[item.UserId]
		for _, rid := range ur {
			for _, r := range rList {
				if r.Id == rid {
					roles = append(roles, r)
					break
				}
			}
		}

		list = append(list, convertUserInfoOut(item, roles))
	}

	resp := &accountrpc.FindUserInfoListResp{}
	resp.Total = total
	resp.List = list

	return resp, nil
}
