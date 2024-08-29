package remark

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/client/remarkrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRemarkLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建留言
func NewAddRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRemarkLogic {
	return &AddRemarkLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddRemarkLogic) AddRemark(req *types.Remark) (resp *types.Remark, err error) {
	in := ConvertRemarkPb(req)
	out, err := l.svcCtx.RemarkRpc.AddRemark(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = ConvertRemarkTypes(out)
	return resp, nil
}

func ConvertRemarkPb(req *types.Remark) (out *remarkrpc.RemarkNew) {
	return &remarkrpc.RemarkNew{
		Id:             req.Id,
		Nickname:       req.Nickname,
		Avatar:         req.Avatar,
		MessageContent: req.MessageContent,
		IpAddress:      req.IpAddress,
		IpSource:       req.IpSource,
		Time:           req.Time,
		IsReview:       req.IsReview,
	}
}

func ConvertRemarkTypes(req *remarkrpc.RemarkDetails) (out *types.Remark) {
	return &types.Remark{
		Id:             req.Id,
		Nickname:       req.Nickname,
		Avatar:         req.Avatar,
		MessageContent: req.MessageContent,
		IpAddress:      req.IpAddress,
		IpSource:       req.IpSource,
		Time:           req.Time,
		IsReview:       req.IsReview,
		CreatedAt:      req.CreatedAt,
		UpdatedAt:      req.UpdatedAt,
	}
}
