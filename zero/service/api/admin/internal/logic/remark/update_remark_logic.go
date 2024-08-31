package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/admin/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/remarkrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新留言
func NewUpdateRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRemarkLogic {
	return &UpdateRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRemarkLogic) UpdateRemark(req *types.RemarkNewReq) (resp *types.RemarkBackDTO, err error) {
	in := ConvertRemarkPb(req)
	out, err := l.svcCtx.RemarkRpc.UpdateRemark(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertRemarkTypes(out)
	return resp, nil
}

func ConvertRemarkPb(in *types.RemarkNewReq) (out *remarkrpc.RemarkNewReq) {
	out = &remarkrpc.RemarkNewReq{
		Id:             in.Id,
		Nickname:       in.Nickname,
		Avatar:         in.Avatar,
		MessageContent: in.MessageContent,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		Time:           in.Time,
		IsReview:       in.IsReview,
	}

	return
}

func ConvertRemarkTypes(in *remarkrpc.RemarkDetails) (out *types.RemarkBackDTO) {
	out = &types.RemarkBackDTO{
		Id:             in.Id,
		Nickname:       in.Nickname,
		Avatar:         in.Avatar,
		MessageContent: in.MessageContent,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		Time:           in.Time,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
	}

	return
}
