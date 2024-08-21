package userrpclogic

import (
	"context"
	"time"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

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
func (l *FindUserOnlineListLogic) FindUserOnlineList(in *blog.FindUserListReq) (*blog.FindUserListResp, error) {
	page, size, sorts, conditions, params := convertQuery(in)
	if conditions != "" {
		conditions += " and "
	}
	conditions += "logout_at < ? and login_at < logout_at"
	params = append(params, time.Now().Add(-time.Hour*24))

	result, err := l.svcCtx.UserAccountModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.UserAccountModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	var userIds []int64
	for _, item := range result {
		userIds = append(userIds, item.Id)
	}

	// 查找用户角色
	urList, err := l.svcCtx.UserRoleModel.FindALL(l.ctx, "user_id in ?", userIds)
	if err != nil {
		return nil, err
	}

	var ursMap = make(map[int64][]int64)
	var roleIds []int64
	for _, item := range urList {
		roleIds = append(roleIds, item.RoleId)
		ursMap[item.UserId] = append(ursMap[item.UserId], item.RoleId)
	}

	// 查找角色信息
	rList, err := l.svcCtx.RoleModel.FindALL(l.ctx, "id in ?", roleIds)
	if err != nil {
		return nil, err
	}

	var list []*blog.UserDetails
	for _, item := range result {

		var roles []*model.Role
		ur, _ := ursMap[item.Id]
		for _, rid := range ur {
			for _, r := range rList {
				if r.Id == rid {
					roles = append(roles, r)
					break
				}
			}
		}

		list = append(list, convert.ConvertUserDetailsModelToPb(item, roles))
	}

	resp := &blog.FindUserListResp{}
	resp.Total = total
	resp.List = list

	return resp, nil
}
