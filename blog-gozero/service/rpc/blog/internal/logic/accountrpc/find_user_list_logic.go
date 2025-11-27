package accountrpclogic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/common/query"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/blog-gozero/service/rpc/blog/internal/svc"
)

type FindUserListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserListLogic {
	return &FindUserListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 查找用户列表
func (l *FindUserListLogic) FindUserList(in *accountrpc.FindUserListReq) (*accountrpc.FindUserListResp, error) {
	page, size, sorts, conditions, params := convertUserQuery(in)

	records, total, err := l.svcCtx.TUserModel.FindListAndTotal(l.ctx, page, size, sorts, conditions, params...)
	if err != nil {
		return nil, err
	}

	var list []*accountrpc.User
	for _, item := range records {
		list = append(list, convertUserOut(item))
	}

	resp := &accountrpc.FindUserListResp{}
	resp.Total = total
	resp.List = list

	return resp, nil
}

func convertUserQuery(in *accountrpc.FindUserListReq) (page int, size int, sorts string, conditions string, params []interface{}) {
	opts := []query.Option{
		query.WithPage(int(in.Page)),
		query.WithSize(int(in.PageSize)),
		query.WithSorts(in.Sorts...),
	}

	if in.Username != "" {
		opts = append(opts, query.WithCondition("username like ?", "%"+in.Username+"%"))
	}

	if in.Nickname != "" {
		opts = append(opts, query.WithCondition("nickname like ?", "%"+in.Nickname+"%"))
	}

	if in.Email != "" {
		opts = append(opts, query.WithCondition("email like ?", "%"+in.Email+"%"))
	}

	if in.Status != 0 {
		opts = append(opts, query.WithCondition("status = ?", in.Status))
	}

	if len(in.UserIds) != 0 {
		opts = append(opts, query.WithCondition("user_id in (?)", in.UserIds))
	}

	return query.NewQueryBuilder(opts...).Build()
}

func convertUserOut(in *model.TUser) (out *accountrpc.User) {

	out = &accountrpc.User{
		UserId:       in.UserId,
		Username:     in.Username,
		Nickname:     in.Nickname,
		Avatar:       in.Avatar,
		Email:        in.Email,
		Phone:        in.Phone,
		Info:         in.Info,
		Status:       in.Status,
		RegisterType: in.RegisterType,
		IpAddress:    in.IpAddress,
		IpSource:     in.IpSource,
		CreatedAt:    in.CreatedAt.Unix(),
		UpdatedAt:    in.UpdatedAt.Unix(),
	}

	return out
}
