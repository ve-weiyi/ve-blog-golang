package logic

import (
	"context"

	entity2 "github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
)

// 获取所有匿名Api记录
func (s *ApiRepository) FindAllPublicApis(ctx context.Context) (list []*entity2.Api, err error) {
	// 创建db
	db := s.DbEngin

	var apis []*entity2.Api
	err = db.Where("access_type = ?", 1).Find(&apis).Error
	if err != nil {
		return nil, err
	}

	return apis, nil
}

// 获取Api记录
func (s *ApiRepository) FindApiRoles(apiId int) (list []*entity2.Role, err error) {
	// 创建db
	db := s.DbEngin
	var roleApis []*entity2.RoleApi

	err = db.Where("api_id = ?", apiId).Find(&roleApis).Error
	if err != nil {
		return nil, err
	}

	var roleIds []int
	for _, item := range roleApis {
		roleIds = append(roleIds, item.ApiID)
	}

	var roles []*entity2.Role
	err = db.Where("id in (?)", roleIds).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}
