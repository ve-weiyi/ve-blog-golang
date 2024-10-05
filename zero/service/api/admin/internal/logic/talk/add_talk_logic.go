package talk

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/talkrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建说说
func NewAddTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTalkLogic {
	return &AddTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddTalkLogic) AddTalk(req *types.TalkNewReq) (resp *types.TalkBackDTO, err error) {
	in := ConvertTalkPb(req)
	in.UserId = cast.ToInt64(l.ctx.Value("uid"))
	out, err := l.svcCtx.TalkRpc.AddTalk(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertTalkTypes(out, nil)
	return resp, nil
}

func ConvertTalkPb(in *types.TalkNewReq) (out *talkrpc.TalkNewReq) {
	out = &talkrpc.TalkNewReq{
		Id:      in.Id,
		UserId:  0,
		Content: in.Content,
		ImgList: in.ImgList,
		IsTop:   in.IsTop,
		Status:  in.Status,
	}

	return
}

func ConvertTalkTypes(in *talkrpc.TalkDetails, usm map[int64]*accountrpc.User) (out *types.TalkBackDTO) {
	out = &types.TalkBackDTO{
		Id:           in.Id,
		UserId:       in.UserId,
		Nickname:     "",
		Avatar:       "",
		Content:      in.Content,
		ImgList:      in.ImgList,
		IsTop:        in.IsTop,
		Status:       in.Status,
		LikeCount:    in.LikeCount,
		CommentCount: in.CommentCount,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}

	// 用户信息
	if out.UserId != 0 {
		user, ok := usm[out.UserId]
		if ok && user != nil {
			out.Nickname = user.Nickname
			out.Avatar = user.Avatar
		}
	}

	return
}
