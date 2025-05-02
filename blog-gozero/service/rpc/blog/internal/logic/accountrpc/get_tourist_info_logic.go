package accountrpclogic

import (
	"context"
	"time"

	"github.com/mssola/useragent"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/rpcutils"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTouristInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTouristInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTouristInfoLogic {
	return &GetTouristInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取游客身份
func (l *GetTouristInfoLogic) GetTouristInfo(in *accountrpc.EmptyReq) (*accountrpc.GetTouristInfoResp, error) {
	ip, err := rpcutils.GetRemoteIPFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	ua, err := rpcutils.GetRemoteAgentFromCtx(l.ctx)
	if err != nil {
		return nil, err
	}

	// 分割字符串，提取 IP 部分
	is, _ := ipx.GetIpSourceByBaidu(ip)
	os := useragent.New(ua).OS()
	browser, _ := useragent.New(ua).Browser()

	terminalId := crypto.Md5v(ip+os+browser, "")

	visitor := &model.TVisitor{
		Id:         0,
		TerminalId: terminalId,
		Os:         os,
		Browser:    browser,
		IpAddress:  ip,
		IpSource:   is,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	// 插入访客记录
	_, err = l.svcCtx.TVisitorModel.Insert(l.ctx, visitor)
	if err != nil {
		return nil, err
	}

	return &accountrpc.GetTouristInfoResp{
		TouristId: terminalId,
	}, nil
}
