package accountrpclogic

import (
	"context"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/gozero/global/constant"
	"github.com/ve-weiyi/ve-blog-golang/gozero/internal/rpcutil"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/model"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/common/rediskey"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/pb/accountrpc"
	"github.com/ve-weiyi/ve-blog-golang/gozero/service/rpc/blog/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/biz/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/ipx"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
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
func (l *RegisterLogic) Register(in *accountrpc.RegisterReq) (*accountrpc.LoginResp, error) {
	// 校验邮箱格式
	if !valid.IsEmailValid(in.Username) {
		return nil, apierr.NewApiError(apierr.CodeInvalidParam, "邮箱格式不正确")
	}

	// 获取用户
	exist, err := l.svcCtx.TUserModel.FindOneByUsername(l.ctx, in.Username)
	if exist != nil {
		return nil, apierr.NewApiError(apierr.CodeUserAlreadyExist, "用户已存在")
	}

	// 验证code是否正确
	key := rediskey.GetCaptchaKey(constant.Register, in.Username)
	if !l.svcCtx.CaptchaHolder.VerifyCaptcha(key, in.VerifyCode) {
		return nil, apierr.NewApiError(apierr.CodeCaptchaVerify, "验证码错误")
	}

	var ua *model.TUser
	err = l.svcCtx.Gorm.Transaction(func(tx *gorm.DB) error {
		ua, err = l.register(tx, in)
		return err
	})
	if err != nil {
		return nil, err
	}

	user, err := l.svcCtx.TUserModel.FindOneByUserId(l.ctx, ua.UserId)
	if err != nil {
		return nil, err
	}

	resp := &accountrpc.LoginResp{
		UserId:   user.UserId,
		Username: user.Username,
		Nickname: user.Nickname,
		Avatar:   user.Avatar,
		Info:     user.Info,
	}

	return resp, nil
}

func (l *RegisterLogic) register(tx *gorm.DB, in *accountrpc.RegisterReq) (out *model.TUser, err error) {
	ip, _ := rpcutil.GetRPCClientIP(l.ctx)
	is, _ := ipx.GetIpSourceByBaidu(ip)

	// 邮箱注册
	user := &model.TUser{
		UserId:    uuid.NewString(),
		Username:  in.Username,
		Password:  crypto.BcryptHash(in.Password),
		Nickname:  in.Username,
		Avatar:    "https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG",
		Email:     in.Username,
		Phone:     "",
		Info:      "",
		Status:    constant.UserStatusNormal,
		LoginType: constant.LoginTypeEmail,
		IpAddress: ip,
		IpSource:  is,
	}

	return onRegister(l.ctx, l.svcCtx, tx, user)
}

func onRegister(ctx context.Context, svcCtx *svc.ServiceContext, tx *gorm.DB, user *model.TUser) (out *model.TUser, err error) {
	/** 创建用户 **/
	_, err = svcCtx.TUserModel.WithTransaction(tx).Insert(ctx, user)
	if err != nil {
		return nil, err
	}

	// 查找默认用户角色
	roles, err := svcCtx.TRoleModel.WithTransaction(tx).FindALL(ctx, "is_default = ?", 1)
	if err != nil {
		return nil, err
	}

	var userRoles []*model.TUserRole
	for _, item := range roles {
		userRoles = append(userRoles, &model.TUserRole{
			UserId: user.UserId,
			RoleId: item.Id,
		})
	}

	/** 创建用户角色 **/
	_, err = svcCtx.TUserRoleModel.WithTransaction(tx).Inserts(ctx, userRoles...)
	if err != nil {
		return nil, err
	}

	return user, nil
}
