package userrpclogic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLoginHistoryListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLoginHistoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLoginHistoryListLogic {
	return &FindUserLoginHistoryListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查询用户登录历史
func (l *FindUserLoginHistoryListLogic) FindUserLoginHistoryList(in *blog.FindLoginHistoryListReq) (*blog.FindLoginHistoryListResp, error) {
	page, size := int(in.Page), int(in.PageSize)
	sorts := ""
	conditions := "user_id = ?"
	params := []interface{}{in.UserId}

	result, err := l.svcCtx.UserLoginHistoryModel.FindList(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*blog.UserLoginHistory
	for _, item := range result {
		list = append(list, convert.ConvertUserLoginHistoryModelToPb(item))
	}

	total, err := l.svcCtx.UserLoginHistoryModel.FindCount(l.ctx, conditions, params...)
	if err != nil {
		return nil, err
	}

	resp := &blog.FindLoginHistoryListResp{}
	resp.List = list
	resp.Total = total
	return resp, nil
}
