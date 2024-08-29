package talkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/talkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddTalkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddTalkLogic {
	return &AddTalkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建说说
func (l *AddTalkLogic) AddTalk(in *talkrpc.TalkNew) (*talkrpc.TalkDetails, error) {
	entity := ConvertTalkIn(in)

	_, err := l.svcCtx.TalkModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return ConvertTalkOut(entity), nil
}

func ConvertTalkIn(in *talkrpc.TalkNew) (out *model.Talk) {
	out = &model.Talk{
		Id:        in.Id,
		UserId:    in.UserId,
		Content:   in.Content,
		Images:    jsonconv.ObjectToJson(in.ImgList),
		IsTop:     in.IsTop,
		Status:    in.Status,
		LikeCount: 0,
	}

	return out
}

func ConvertTalkOut(in *model.Talk) (out *talkrpc.TalkDetails) {
	var images []string
	jsonconv.JsonToObject(in.Images, &images)

	out = &talkrpc.TalkDetails{
		Id:           in.Id,
		UserId:       in.UserId,
		Content:      in.Content,
		ImgList:      images,
		IsTop:        in.IsTop,
		Status:       in.Status,
		LikeCount:    in.LikeCount,
		CommentCount: 0,
		CreatedAt:    in.CreatedAt.Unix(),
		UpdatedAt:    in.UpdatedAt.Unix(),
	}

	return out
}
