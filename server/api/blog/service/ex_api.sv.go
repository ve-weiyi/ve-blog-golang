package service

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/dto"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/base/request"
)

// 分页获取Api记录
func (l *ApiService) FindApiDetailsList(reqCtx *request.Context, page *dto.PageQuery) (list []*dto.ApiDetailsDTO, total int64, err error) {
	cond, args := page.ConditionClause()
	// 查询api信息
	apis, err := l.svcCtx.ApiRepository.FindALL(reqCtx, cond, args...)
	if err != nil {
		return nil, 0, err
	}

	// to tree
	var tree dto.ApiDetailsDTO
	tree.Children = getApiChildren(tree, apis)

	list = tree.Children
	return list, int64(len(list)), nil
}

func (l *ApiService) SyncApiList(reqCtx *request.Context, req interface{}) (data int64, err error) {
	//ap := apiparser.NewSwaggerParser()
	//apis, err := ap.ParseApiDocsByRoots(files.GetRuntimeRoot() + "server/docs")
	//if err != nil {
	//	return 0, err
	//}
	//
	//var apiModels []*entity.Api
	//for _, api := range apis {
	//	if api.Router == "" {
	//		continue
	//	}
	//
	//	// 已存在则跳过
	//	exist, _ := l.svcCtx.ApiRepository.First(reqCtx, "path = ? and method = ?", api.Router, api.Method)
	//	if exist != nil {
	//		continue
	//	}
	//
	//	// 查找父分类，没有则创建
	//	parent, _ := l.svcCtx.ApiRepository.First(reqCtx, "name = ? and parent_id = ?", api.Tag, 0)
	//	if parent == nil {
	//		parent = &entity.Api{
	//			Group: api.Tag,
	//		}
	//		_, err = l.svcCtx.ApiRepository.Create(reqCtx, parent)
	//		if err != nil {
	//			return 0, err
	//		}
	//	}
	//
	//	var traceable int64
	//	if strings.ToUpper(api.Method) == "PUT" || strings.ToUpper(api.Method) == "DELETE" {
	//		traceable = 1
	//	}
	//	if strings.ToUpper(api.Method) == "POST" && !strings.Contains(api.Router, "list") {
	//		traceable = 1
	//	}
	//
	//	// 插入数据
	//	model := &entity.Api{
	//		Group:      api.Summary,
	//		Path:      api.Router,
	//		Method:    strings.ToUpper(api.Method),
	//		ParentId:  parent.Id,
	//		Traceable: traceable,
	//	}
	//
	//	apiModels = append(apiModels, model)
	//	//_, err = l.svcCtx.ApiRepository.Create(reqCtx, model)
	//	//if err != nil {
	//	//	return 0, err
	//	//}
	//	//data++
	//}
	//
	//// 批量插入，减少数据库压力
	//query := l.svcCtx.ApiRepository.DbEngin.CreateInBatches(apiModels, len(apiModels))
	//data = query.RowsAffected
	//err = query.Error
	//if err != nil {
	//	return 0, err
	//}
	return data, nil
}

func (l *MenuService) CleanApiList(reqCtx *request.Context, req interface{}) (data interface{}, err error) {
	return l.svcCtx.ApiRepository.CleanApis(reqCtx)
}
func getApiChildren(root dto.ApiDetailsDTO, list []*entity.Api) (leafs []*dto.ApiDetailsDTO) {
	for _, item := range list {
		if item.ParentId == root.Id {
			leaf := dto.ApiDetailsDTO{
				Api:      *item,
				Children: nil,
			}
			leaf.Children = getApiChildren(leaf, list)
			leafs = append(leafs, &leaf)
		}
	}
	return leafs
}
