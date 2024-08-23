package accountrpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
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
func (l *FindUserListLogic) FindUserList(in *accountrpc.FindUserListReq) (*accountrpc.FindUserListResp, error) {
	page, size, sorts, conditions, params := convertQuery(in)

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

	var list []*accountrpc.UserDetails
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

		list = append(list, ConvertUserDetailsOut(item, roles))
	}

	resp := &accountrpc.FindUserListResp{}
	resp.Total = total
	resp.List = list

	return resp, nil
}

func convertQuery(in *accountrpc.FindUserListReq) (page int, size int, sorts string, conditions string, params []interface{}) {
	page = int(in.Page)
	size = int(in.PageSize)

	if in.Username != "" {
		conditions += "username like ?"
		params = append(params, "%"+in.Username+"%")
	}

	if in.Nickname != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions += "nickname like ?"
		params = append(params, "%"+in.Nickname+"%")
	}

	if in.Email != "" {
		if conditions != "" {
			conditions += " and "
		}
		conditions += "email like ?"
		params = append(params, "%"+in.Email+"%")
	}

	if in.Status != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions += "status = ?"
		params = append(params, in.Status)
	}

	return page, size, sorts, conditions, params
}
