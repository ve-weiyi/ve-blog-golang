package logic

import (
	"context"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

// 获取所有匿名Api记录
func (s *ApiRepository) FindAllPublicApis(ctx context.Context) (list []*entity.Api, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)

	var apis []*entity.Api
	err = db.Where("access_type = ?", 1).Find(&apis).Error
	if err != nil {
		return nil, err
	}

	return apis, nil
}

// 获取Api记录
func (s *ApiRepository) FindApiRoles(ctx context.Context, apiId int) (list []*entity.Role, err error) {
	// 创建db
	db := s.DbEngin.WithContext(ctx)
	var roleApis []*entity.RoleApi

	err = db.Where("api_id = ?", apiId).Find(&roleApis).Error
	if err != nil {
		return nil, err
	}

	var roleIds []int
	for _, item := range roleApis {
		roleIds = append(roleIds, item.ApiID)
	}

	var roles []*entity.Role
	err = db.Where("id in (?)", roleIds).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}
