package talk

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/common/apiutils"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/socialservice"
)

type GetTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取说说详情
func NewGetTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTalkLogic {
	return &GetTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTalkLogic) GetTalk(req *types.GetTalkReq) (resp *types.TalkVO, err error) {
	out, err := l.svcCtx.SocialService.GetTalk(l.ctx, &socialservice.GetTalkRequest{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	usm, err := apiutils.GetUserInfos(l.ctx, l.svcCtx, []string{out.Talk.UserId})
	if err != nil {
		return nil, err
	}

	return &types.TalkVO{
		Id:           out.Talk.Id,
		UserId:       out.Talk.UserId,
		Content:      out.Talk.Content,
		ImgList:      out.Talk.Images,
		IsTop:        out.Talk.IsTop,
		Status:       out.Talk.Status,
		LikeCount:    out.Talk.LikeCount,
		CommentCount: out.Talk.CommentCount,
		CreatedAt:    out.Talk.CreatedAt,
		UpdatedAt:    out.Talk.UpdatedAt,
		UserInfo:     usm[out.Talk.UserId],
	}, nil
}
