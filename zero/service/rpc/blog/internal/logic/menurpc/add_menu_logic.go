package menurpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddMenuLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMenuLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMenuLogic {
	return &AddMenuLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建菜单
func (l *AddMenuLogic) AddMenu(in *blog.Menu) (*blog.Menu, error) {
	entity := convert.ConvertMenuPbToModel(in)

	_, err := l.svcCtx.MenuModel.Insert(l.ctx, entity)
	if err != nil {
		return nil, err
	}

	return convert.ConvertMenuModelToPb(entity), nil
}
