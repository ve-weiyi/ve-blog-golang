package talk

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
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

func (l *AddTalkLogic) AddTalk(req *types.TalkNew) (resp *types.TalkDetails, err error) {
	in := ConvertTalkPb(req)

	out, err := l.svcCtx.TalkRpc.UpdateTalk(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return ConvertTalkTypes(out), nil
}

func ConvertTalkPb(in *types.TalkNew) (out *talkrpc.TalkNew) {
	out = &talkrpc.TalkNew{
		Id:      in.Id,
		UserId:  in.UserId,
		Content: in.Content,
		Images:  jsonconv.ObjectToJson(in.ImgList),
		IsTop:   in.IsTop,
		Status:  in.Status,
	}

	return
}

func ConvertTalkTypes(in *talkrpc.TalkDetails) (out *types.TalkDetails) {
	out = &types.TalkDetails{
		Id:           in.Id,
		UserId:       in.UserId,
		Nickname:     in.Nickname,
		Avatar:       in.Avatar,
		Content:      in.Content,
		ImgList:      in.ImgList,
		IsTop:        in.IsTop,
		Status:       in.Status,
		LikeCount:    in.LikeCount,
		CommentCount: 0,
		CreatedAt:    in.CreatedAt,
		UpdatedAt:    in.UpdatedAt,
	}

	return
}
