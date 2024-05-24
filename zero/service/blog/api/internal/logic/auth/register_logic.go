package auth

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/api/internal/types"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(reqCtx *types.RestHeader, req *types.LoginReq) (resp *types.EmptyResp, err error) {
	in := &blog.LoginReq{
		Username: req.Username,
		Password: req.Password,
		Code:     req.Code,
	}

	_, err = l.svcCtx.AuthRpc.Register(l.ctx, in)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
