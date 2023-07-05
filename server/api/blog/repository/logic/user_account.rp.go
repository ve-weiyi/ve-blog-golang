package logic

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-admin-store/server/infra/constant"
	"strings"
)

// 注销用户
func (s *UserAccountRepository) Logoff(ctx context.Context, id int) (data any, err error) {
	// 创建db
	db := s.DbEngin
	var account entity.UserAccount

	// 开始事务
	tx := db.Begin()

	// 执行数据库操作
	db.Where("id = ?", id).First(&account)

	// 删除用户
	account.Status = constant.StatusDisable
	db.Save(&account)

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		// 处理提交事务错误
		tx.Rollback() // 回滚事务
		return nil, err
	}

	return account, err
}

// 注册
func (s *UserAccountRepository) Register(ctx context.Context, account *entity.UserAccount, info *entity.UserInformation) (u *entity.UserAccount, i *entity.UserInformation, err error) {
	//var role entity.Role

	db := s.DbEngin
	// 开始事务
	tx := db.Begin()
	// 执行数据库操作
	/** 创建用户 **/
	db.Create(&account)

	/** 创建用户信息 start**/
	info.UserID = account.ID

	// 默认邮箱
	if info.Email == "" {
		if account.RegisterType == constant.LoginEmail {
			info.Email = account.Username
		}
	}

	// 默认昵称
	if info.Nickname == "" {
		if info.Email != "" {
			info.Nickname = strings.Split(info.Email, "@")[0]
		} else {
			info.Nickname = fmt.Sprintf("游客%v", uuid.New().String()[:8])
		}
	}

	// 默认头像
	if info.Avatar == "" {
		info.Avatar = "https://mms1.baidu.com/it/u=2815887849,1501151317&fm=253&app=138&f=JPEG"
	}

	// 默认简介
	if info.Intro == "" {
		info.Intro = "这个人很神秘，什么都没有写！"
	}

	db.Create(&info)
	/** 创建用户信息 end **/

	/** 创建用户角色 end **/
	var roles []*entity.Role
	db.Where("is_default = ?", 1).Find(&roles)

	var userRoles []*entity.UserRole
	for _, item := range roles {
		userRoles = append(userRoles, &entity.UserRole{
			UserID: account.ID,
			RoleID: item.ID,
		})
	}

	db.Create(&userRoles)
	/** 创建用户角色 end **/

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		// 处理提交事务错误
		tx.Rollback() // 回滚事务
		return nil, nil, err
	}

	return account, info, nil
}

func (s *UserAccountRepository) FindUserMenus(userId int) (list []*entity.Menu, err error) {
	// 创建db
	db := s.DbEngin

	//查询用户信息
	var account entity.UserAccount
	err = db.Where("user_id = ?", userId).First(&account).Error
	if err != nil {
		return nil, err
	}

	//查询用户的角色
	var roleApis []*entity.RoleMenu
	err = db.Where("user_id = ?", userId).Find(&roleApis).Error
	if err != nil {
		return nil, err
	}

	var ids []int
	for _, item := range roleApis {
		ids = append(ids, item.MenuID)
	}

	var apis []*entity.Menu
	err = db.Where("id in (?)", ids).Find(&apis).Error
	if err != nil {
		return nil, err
	}

	return apis, nil
}

// 加载用户model
func (s *UserAccountRepository) LoadUserByUsername(username string) (data *entity.UserAccount, err error) {
	// 创建db
	db := s.DbEngin

	//查询用户信息
	err = db.Where("username = ?", username).First(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *UserAccountRepository) GetUserDetails(userid string) (data *entity.UserInformation, err error) {
	// 创建db
	db := s.DbEngin

	//查询用户信息
	err = db.Where("user_id = ?", userid).First(&data).Error
	if err != nil {
		return nil, err
	}

	return data, nil
}
