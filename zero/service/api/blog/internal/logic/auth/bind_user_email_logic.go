package auth

import (
	"context"

	"github.com/spf13/cast"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/api/blog/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type BindUserEmailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 绑定邮箱
func NewBindUserEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindUserEmailLogic {
	return &BindUserEmailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BindUserEmailLogic) BindUserEmail(req *types.BindUserEmailReq) (resp *types.EmptyResp, err error) {

	in := &blog.BindUserEmailReq{
		UserId:     cast.ToInt64(l.ctx.Value("uid")),
		Email:      req.Email,
		VerifyCode: req.VerifyCode,
	}

	_, err = l.svcCtx.AuthRpc.BindUserEmail(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
