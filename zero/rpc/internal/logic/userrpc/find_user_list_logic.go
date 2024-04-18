package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/repository/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserListLogic {
	return &FindUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找用户列表
func (l *FindUserListLogic) FindUserList(in *account.PageQuery) (*account.PageUserInfoResp, error) {
	limit, offset, sorts, conditions, params := convert.ParsePageQuery(in)

	result, err := l.svcCtx.UserAccountModel.FindList(l.ctx, limit, offset, sorts, conditions, params)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.UserAccountModel.FindCount(l.ctx, conditions, params)
	if err != nil {
		return nil, err
	}

	var userIds []int64
	for _, item := range result {
		userIds = append(userIds, item.Id)
	}

	// 查找用户信息
	uiList, err := l.svcCtx.UserInformationModel.FindALL(l.ctx, "user_id in ?", userIds)
	if err != nil {
		return nil, err
	}

	// 将用户信息列表转换为map
	var uiMap = make(map[int64]*model.UserInformation)
	for _, item := range uiList {
		uiMap[item.UserId] = item
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

	var list []*account.UserDTO
	for _, item := range result {
		ui, ok := uiMap[item.Id]
		if !ok {
			continue
		}

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

		list = append(list, convert.ConvertUserDetailsModelToPb(item, ui, roles))
	}

	resp := &account.PageUserInfoResp{}
	resp.Total = total
	resp.List = list

	return resp, nil
}
