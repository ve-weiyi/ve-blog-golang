package talk

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *FindTalkListLogic) FindTalkList(reqCtx *types.RestHeader, req *types.PageQuery) (resp *types.PageResp, err error) {
	in := convert.ConvertPageQuery(req)
	out, err := l.svcCtx.TalkRpc.FindTalkList(l.ctx, in)
	if err != nil {
		return nil, err
	}

	total, err := l.svcCtx.TalkRpc.FindTalkCount(l.ctx, in)
	if err != nil {
		return nil, err
	}

	var list []*types.TalkDetails
	for _, v := range out.List {

		m := convert.ConvertTalkTypes(v)
		user, _ := l.svcCtx.UserRpc.FindUserInfo(l.ctx, &blog.UserReq{UserId: v.UserId})
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
	resp.Total = total.Count
	resp.List = list
	return resp, nil
}
