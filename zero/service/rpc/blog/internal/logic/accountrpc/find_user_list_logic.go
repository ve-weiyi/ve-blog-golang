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

	result, err := l.svcCtx.TUserModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.TUserModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	var uids []int64
	for _, item := range result {
		uids = append(uids, item.Id)
	}

	// 查找用户角色
	urList, err := l.svcCtx.TUserRoleModel.FindALL(l.ctx, "user_id in (?)", uids)
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
	rList, err := l.svcCtx.TRoleModel.FindALL(l.ctx, "id in (?)", roleIds)
	if err != nil {
		return nil, err
	}

	var list []*accountrpc.UserInfoResp
	for _, item := range result {

		var roles []*model.TRole
		ur, _ := ursMap[item.Id]
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

	if len(in.UserIds) != 0 {
		if conditions != "" {
			conditions += " and "
		}
		conditions += "id in (?)"
		params = append(params, in.UserIds)
	}

	return page, size, sorts, conditions, params
}
