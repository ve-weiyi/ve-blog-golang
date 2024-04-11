package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/pb/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindMenuLogic {
	return &FindMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询菜单
func (l *FindMenuLogic) FindMenu(in *account.IdReq) (*account.Menu, error) {
	result, err := l.svcCtx.MenuModel.First(l.ctx, "id = ?", in.Id)
	if err != nil {
		return nil, err
	}

	return convert.ConvertMenuModelToPb(result), nil
}
