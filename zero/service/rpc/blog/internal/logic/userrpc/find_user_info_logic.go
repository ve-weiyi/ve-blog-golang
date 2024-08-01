package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/pb/blog"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserInfoLogic {
	return &FindUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户信息
func (l *FindUserInfoLogic) FindUserInfo(in *blog.UserIdReq) (*blog.UserInfoResp, error) {
	uid := in.UserId

	// ua, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", uid)
	// if err != nil {
	//	return nil, err
	// }

	ui, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", uid)
	if err != nil {
		return nil, err
	}

	return convert.ConvertUserInfoModelToPb(ui), nil
}
