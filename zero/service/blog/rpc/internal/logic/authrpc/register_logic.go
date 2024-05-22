package authrpclogic

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/apierr"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/crypto"
	"github.com/ve-weiyi/ve-blog-golang/kit/utils/valid"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/convert"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/internal/svc"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/blog/rpc/pb/blog"

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
func (l *RegisterLogic) Register(in *blog.LoginReq) (*blog.UserInfoResp, error) {
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
	if !l.svcCtx.CaptchaHolder.VerifyCaptcha(key, in.Code) {
		return nil, apierr.ErrorCaptchaVerify
	}

	// 邮箱注册
	account := &model.UserAccount{
		Username:     in.Username,
		Password:     crypto.BcryptHash(in.Password),
		Status:       constant.UserStatusNormal,
		RegisterType: constant.LoginTypeEmail,
		IpAddress:    "",
		IpSource:     "",
	}

	info := &model.UserInformation{
		Email:    in.Username,
		Nickname: in.Username,
		Avatar:   "https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG",
		Phone:    "",
		Intro:    "这个人很神秘，什么都没有写！",
		Website:  "",
	}

	err = l.svcCtx.Gorm.Transaction(func(tx *gorm.DB) error {
		_, err = l.register(l.ctx, tx, account, info)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	ui, err := l.svcCtx.UserInformationModel.First(l.ctx, "user_id = ?", account.Id)
	if err != nil {
		return nil, err
	}

	return convert.ConvertUserInfoModelToPb(ui), nil
}

func (l *RegisterLogic) register(ctx context.Context, tx *gorm.DB, account *model.UserAccount, info *model.UserInformation) (out *model.UserAccount, err error) {

	/** 创建用户 **/
	_, err = l.svcCtx.UserAccountModel.WithTransaction(tx).Insert(ctx, account)
	if err != nil {
		return nil, err
	}

	/** 创建用户信息 **/
	info.UserId = account.Id
	_, err = l.svcCtx.UserInformationModel.WithTransaction(tx).Insert(ctx, info)
	if err != nil {
		return nil, err
	}

	/** 创建用户角色 end **/
	roles, err := l.svcCtx.RoleModel.WithTransaction(tx).FindALL(ctx, "is_default = ?", 1)
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

	_, err = l.svcCtx.UserRoleModel.WithTransaction(tx).InsertBatch(ctx, userRoles...)
	if err != nil {
		return nil, err
	}

	return account, nil
}
