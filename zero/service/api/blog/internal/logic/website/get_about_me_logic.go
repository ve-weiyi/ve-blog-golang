package website

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/jsonconv"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAboutMeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取关于我的信息
func NewGetAboutMeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAboutMeLogic {
	return &GetAboutMeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAboutMeLogic) GetAboutMe(req *types.EmptyReq) (resp *types.AboutMe, err error) {
	in := &blog.FindConfigReq{
		ConfigKey: "about_me",
	}

	out, err := l.svcCtx.ConfigRpc.FindConfig(l.ctx, in)
	if err != nil {
		return nil, err
	}

	resp = &types.AboutMe{}
	jsonconv.JsonToObject(out.ConfigValue, &resp)
	return resp, nil
}
