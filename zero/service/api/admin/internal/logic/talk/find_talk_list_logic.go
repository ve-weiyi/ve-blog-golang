package talk

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/talkrpc"
)

type FindTalkListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分页获取说说列表
func NewFindTalkListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindTalkListLogic {
	return &FindTalkListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindTalkListLogic) FindTalkList(req *types.TalkQuery) (resp *types.PageResp, err error) {
	in := &talkrpc.FindTalkListReq{
		Page:     req.Page,
		PageSize: req.PageSize,
	}

	out, err := l.svcCtx.TalkRpc.FindTalkList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.TalkDetails
	for _, v := range out.List {

		m := ConvertTalkTypes(v)
		user, _ := l.svcCtx.AccountRpc.GetUserInfo(l.ctx, &accountrpc.UserIdReq{UserId: v.UserId})
		if user != nil {
			m.UserId = user.UserId
			m.Nickname = user.Nickname
			m.Avatar = user.Avatar
		}

		list = append(list, m)
	}

	resp = &types.PageResp{}
	resp.Page = in.Page
	resp.PageSize = in.PageSize
	resp.Total = out.Total
	resp.List = list
	return resp, nil
}
