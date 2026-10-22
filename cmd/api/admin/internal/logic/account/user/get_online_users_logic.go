package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/svc"
	"github.com/ve-weiyi/blog-cloud/cmd/api/admin/internal/types"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/guestservice"
	"github.com/ve-weiyi/blog-cloud/cmd/rpc/blog/client/userservice"
	"github.com/ve-weiyi/blog-cloud/infra/auth/identityx"
)

type GetOnlineUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取在线用户列表
func NewGetOnlineUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOnlineUsersLogic {
	return &GetOnlineUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOnlineUsersLogic) GetOnlineUsers(req *types.EmptyReq) (resp *types.OnlineUserListResp, err error) {
	// 从 blog:online:user 查询博客前台在线用户
	rawUsers, err := l.queryOnlineUsers(l.ctx)
	if err != nil {
		return nil, err
	}

	if len(rawUsers) == 0 {
		return &types.OnlineUserListResp{List: []*types.OnlineUserItem{}}, nil
	}

	// 用户是 36 位 UUIDv7，游客是 32 位小写 hex；按形态分流查询。
	// 形态判定只是提示：历史用户标识可能呈游客形态，查游客未命中的再回查用户。
	var userIDs, guestIDs, unknownIDs []string
	for _, u := range rawUsers {
		if u.Identity == "" {
			continue
		}
		switch identityx.KindOf(u.Identity) {
		case identityx.KindUser:
			userIDs = append(userIDs, u.Identity)
		case identityx.KindGuest:
			guestIDs = append(guestIDs, u.Identity)
		default:
			unknownIDs = append(unknownIDs, u.Identity)
		}
	}

	userQuery := append(append([]string{}, userIDs...), unknownIDs...)
	guestQuery := append(append([]string{}, guestIDs...), unknownIDs...)

	userInfoMap, err := l.listUsers(userQuery)
	if err != nil {
		return nil, err
	}
	guestInfoMap, err := l.listGuests(guestQuery)
	if err != nil {
		return nil, err
	}

	// 游客形态但游客表无记录，可能是历史用户 ID
	var unresolved []string
	for _, id := range guestIDs {
		if _, ok := guestInfoMap[id]; !ok {
			unresolved = append(unresolved, id)
		}
	}
	if len(unresolved) > 0 {
		fallback, err := l.listUsers(unresolved)
		if err != nil {
			return nil, err
		}
		for id, u := range fallback {
			userInfoMap[id] = u
		}
	}

	list := make([]*types.OnlineUserItem, 0, len(rawUsers))
	for _, u := range rawUsers {
		if u.Identity == "" {
			continue
		}

		item := &types.OnlineUserItem{
			LastActiveAt: u.LastActiveAt,
		}

		if ui, ok := userInfoMap[u.Identity]; ok {
			item.UserInfo = &types.UserInfoVO{
				UserId:   ui.UserId,
				Username: ui.Username,
				Nickname: ui.Nickname,
				Avatar:   ui.Avatar,
				UserType: "user",
			}
			list = append(list, item)
			continue
		}

		if ci, ok := guestInfoMap[u.Identity]; ok {
			item.GuestInfo = &types.GuestInfoVO{
				DeviceId:  ci.DeviceId,
				Os:        ci.Os,
				Browser:   ci.Browser,
				IpAddress: ci.IpAddress,
				IpSource:  ci.IpSource,
			}
		}

		list = append(list, item)
	}

	return &types.OnlineUserListResp{List: list}, nil
}

// onlineQueryBatchSize 按标识批量查询时每页的条数，需小于 queryx 的每页上限。
const onlineQueryBatchSize = 200

// listUsers 按用户 uid 分页取回在线用户资料；一页装不下时继续翻页，避免默认分页截断。
func (l *GetOnlineUsersLogic) listUsers(userIds []string) (map[string]*userservice.User, error) {
	userInfoMap := make(map[string]*userservice.User, len(userIds))
	pages := (len(userIds) + onlineQueryBatchSize - 1) / onlineQueryBatchSize
	for page := 1; page <= pages; page++ {
		out, err := l.svcCtx.UserService.ListUsers(l.ctx, &userservice.ListUsersRequest{
			ListQuery: &userservice.ListQuery{
				Page:     int64(page),
				PageSize: onlineQueryBatchSize,
			},
			UserIds: userIds,
		})
		if err != nil {
			return nil, err
		}
		for _, u := range out.List {
			userInfoMap[u.UserId] = u
		}
	}
	return userInfoMap, nil
}

// listGuests 按 device_id 分页取回在线游客资料。
func (l *GetOnlineUsersLogic) listGuests(deviceIds []string) (map[string]*guestservice.Guest, error) {
	guestInfoMap := make(map[string]*guestservice.Guest, len(deviceIds))
	pages := (len(deviceIds) + onlineQueryBatchSize - 1) / onlineQueryBatchSize
	for page := 1; page <= pages; page++ {
		out, err := l.svcCtx.GuestService.ListGuests(l.ctx, &guestservice.ListGuestsRequest{
			ListQuery: &guestservice.ListQuery{
				Page:     int64(page),
				PageSize: onlineQueryBatchSize,
			},
			DeviceIds: deviceIds,
		})
		if err != nil {
			return nil, err
		}
		for _, c := range out.List {
			guestInfoMap[c.DeviceId] = c
		}
	}
	return guestInfoMap, nil
}

// queryOnlineUsers 读博客前台的在线名册（归前台服务写，本服务只读）。
// 名册按连接记，读取经 presencex 按用户聚合 —— 一个用户开几个页面仍只出现一次。
func (l *GetOnlineUsersLogic) queryOnlineUsers(ctx context.Context) ([]*rawOnlineUser, error) {
	users, err := l.svcCtx.SitePresence.List(ctx)
	if err != nil {
		return nil, err
	}

	raw := make([]*rawOnlineUser, 0, len(users))
	for _, u := range users {
		raw = append(raw, &rawOnlineUser{
			Identity:     u.Login,
			LastActiveAt: u.LastActiveAt,
		})
	}
	return raw, nil
}

type rawOnlineUser struct {
	Identity     string // 注册用户的 uid 或游客的 device_id
	LastActiveAt int64
}
