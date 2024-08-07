package authrpclogic

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/pb/blog"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 注册
func (l *RegisterLogic) Register(in *blog.RegisterReq) (*blog.LoginResp, error) {
	// 校验邮箱格式
	if !valid.IsEmailValid(in.Username) {
		return nil, apierr.ErrorInvalidParam
	}

	// 获取用户
	exist, err := l.svcCtx.UserAccountModel.FindOneByUsername(l.ctx, in.Username)
	if exist != nil {
		return nil, apierr.ErrorUserAlreadyExist
	}

	// 验证code是否正确
	key := fmt.Sprintf("%s:%s", constant.Register, in.Username)
	if !l.svcCtx.CaptchaHolder.VerifyCaptcha(key, in.VerifyCode) {
		return nil, apierr.ErrorCaptchaVerify
	}

	var ua *model.UserAccount
	err = l.svcCtx.Gorm.Transaction(func(tx *gorm.DB) error {
		ua, err = l.register(tx, in)
		return err
	})
	if err != nil {
		return nil, err
	}

	account, err := l.svcCtx.UserAccountModel.First(l.ctx, "id = ?", ua.Id)
	if err != nil {
		return nil, err
	}

	resp := &blog.LoginResp{
		UserId:   account.Id,
		Username: account.Username,
		Nickname: account.Nickname,
		Avatar:   account.Avatar,
		Info:     account.Info,
	}

	return resp, nil
}

func (l *RegisterLogic) register(tx *gorm.DB, in *blog.RegisterReq) (out *model.UserAccount, err error) {
	// 邮箱注册
	account := &model.UserAccount{
		Id:        0,
		Username:  in.Username,
		Password:  crypto.BcryptHash(in.Password),
		Nickname:  in.Username,
		Avatar:    "https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG",
		Email:     in.Username,
		Phone:     "",
		Info:      "",
		Status:    constant.UserStatusNormal,
		LoginType: constant.LoginTypeEmail,
		IpAddress: "",
		IpSource:  "",
	}

	return onRegister(l.svcCtx, l.ctx, tx, account)
}

func onRegister(svcCtx *svc.ServiceContext, ctx context.Context, tx *gorm.DB, account *model.UserAccount) (out *model.UserAccount, err error) {
	/** 创建用户 **/
	_, err = svcCtx.UserAccountModel.WithTransaction(tx).Insert(ctx, account)
	if err != nil {
		return nil, err
	}

	// 查找默认用户角色
	roles, err := svcCtx.RoleModel.WithTransaction(tx).FindALL(ctx, "is_default = ?", 1)
	if err != nil {
		return nil, err
	}

	var userRoles []*model.UserRole
	for _, item := range roles {
		userRoles = append(userRoles, &model.UserRole{
			UserId: account.Id,
			RoleId: item.Id,
		})
	}

	/** 创建用户角色 **/
	_, err = svcCtx.UserRoleModel.WithTransaction(tx).InsertBatch(ctx, userRoles...)
	if err != nil {
		return nil, err
	}

	return account, nil
}
