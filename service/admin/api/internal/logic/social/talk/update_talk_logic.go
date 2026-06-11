package talk

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/service/admin/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/service/app/rpc/client/socialservice"
)

type UpdateTalkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新说说
func NewUpdateTalkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTalkLogic {
	return &UpdateTalkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateTalkLogic) UpdateTalk(req *types.UpdateTalkReq) (resp *types.EmptyResp, err error) {
	_, err = l.svcCtx.SocialService.UpdateTalk(l.ctx, &socialservice.UpdateTalkRequest{
		Id:      req.Id,
		Content: req.Content,
		Images:  req.ImgList,
		IsTop:   req.IsTop,
		Status:  req.Status,
	})
	if err != nil {
		return nil, err
	}

	return &types.EmptyResp{}, nil
}
