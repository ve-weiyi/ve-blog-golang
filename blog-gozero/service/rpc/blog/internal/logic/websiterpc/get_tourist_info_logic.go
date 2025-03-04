package websiterpclogic

import (
	"context"
	"strings"

	"github.com/mssola/useragent"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/common/metadatax"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/websiterpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"

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
func (l *GetTouristInfoLogic) GetTouristInfo(in *websiterpc.EmptyReq) (*websiterpc.GetTouristInfoResp, error) {
	ci, err := metadatax.GetRPCClientIP(l.ctx)
	if err != nil {
		return nil, err
	}

	ua, err := metadatax.GetRPCUserAgent(l.ctx)
	if err != nil {
		return nil, err
	}

	// 分割字符串，提取 IP 部分
	ip := strings.Split(ci, ":")[0]
	os := useragent.New(ua).OS()
	browser, _ := useragent.New(ua).Browser()

	identity := crypto.Md5v(ip+os+browser, "")

	return &websiterpc.GetTouristInfoResp{
		TouristId: identity,
	}, nil
}
