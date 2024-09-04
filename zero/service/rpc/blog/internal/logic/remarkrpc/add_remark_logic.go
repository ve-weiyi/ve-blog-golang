package remarkrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
	"github.com/ve-weiyi/ve-blog-golang/zero/internal/rpcutil"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/remarkrpc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddRemarkLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddRemarkLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddRemarkLogic {
	return &AddRemarkLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建留言
func (l *AddRemarkLogic) AddRemark(in *remarkrpc.RemarkNewReq) (*remarkrpc.RemarkDetails, error) {
	ip, _ := rpcutil.GetRPCClientIP(l.ctx)
	is, _ := ipx.GetIpInfoByBaidu(ip)

	entity := &model.Remark{
		Id:             0,
		Nickname:       in.Nickname,
		Avatar:         in.Avatar,
		MessageContent: in.MessageContent,
		IpAddress:      ip,
		IpSource:       is.Location,
		Time:           0,
		IsReview:       1,
		//CreatedAt:      time.Time{},
		//UpdatedAt:      time.Time{},
	}

	_, err := l.svcCtx.RemarkModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convertRemarkOut(entity), nil
}

func convertRemarkOut(in *model.Remark) (out *remarkrpc.RemarkDetails) {
	out = &remarkrpc.RemarkDetails{
		Id:             in.Id,
		Nickname:       in.Nickname,
		Avatar:         in.Avatar,
		MessageContent: in.MessageContent,
		IpAddress:      in.IpAddress,
		IpSource:       in.IpSource,
		Time:           in.Time,
		IsReview:       in.IsReview,
		CreatedAt:      in.CreatedAt.Unix(),
		UpdatedAt:      in.UpdatedAt.Unix(),
	}

	return out
}
