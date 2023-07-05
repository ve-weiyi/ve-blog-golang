package logic

import (
	"context"
	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
)

// 获取所有匿名Api记录
func (s *ApiRepository) FindAllPublicApis(ctx context.Context) (list []*entity.Api, err error) {
	// 创建db
	db := s.DbEngin

	var apis []*entity.Api
	err = db.Where("access_type = ?", 1).Find(&apis).Error
	if err != nil {
		return nil, err
	}

	return apis, nil
}
