package accountrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserInfoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserInfoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserInfoListLogic {
	return &FindUserInfoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找用户信息列表
func (l *FindUserInfoListLogic) FindUserInfoList(in *accountrpc.FindUserListReq) (*accountrpc.FindUserInfoListResp, error) {
	page, size, sorts, conditions, params := convertQuery(in)

	result, err := l.svcCtx.TUserModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.TUserModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	var uids []string
	for _, item := range result {
		uids = append(uids, item.UserId)
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
	for _, item := range result {

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
