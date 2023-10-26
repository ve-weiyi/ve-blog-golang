package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/sqlx"
	"github.com/ve-weiyi/ve-blog-golang/server/utils"
)

type WebsiteService struct {
	svcCtx *svc.ServiceContext
}

func NewWebsiteService(svcCtx *svc.ServiceContext) *WebsiteService {
	return &WebsiteService{
		svcCtx: svcCtx,
	}
}

func (s *WebsiteService) GetWebsiteAdminHomeInfo(reqCtx *request.Context, data interface{}) (resp *response.WebsiteAdminHomeInfo, err error) {
	page := &request.PageQuery{}
	// 查询消息数量
	msgCount, err := s.svcCtx.RemarkRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, err
	}

	// 查询用户数量
	userCount, err := s.svcCtx.UserAccountRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, err
	}

	// 查询文章数量
	articles, err := s.svcCtx.ArticleRepository.FindArticleList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, err
	}

	// 查询分类数量
	categories, err := s.svcCtx.CategoryRepository.FindCategoryList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, err
	}

	// 查询标签数量
	tags, err := s.svcCtx.TagRepository.FindTagList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, err
	}

	uniqueViews, err := s.svcCtx.UniqueViewRepository.FindUniqueViewList(reqCtx, &page.PageLimit, page.Sorts, page.Conditions...)
	if err != nil {
		return nil, err
	}

	articleCount, err := s.svcCtx.ArticleRepository.Count(reqCtx, page.Conditions...)
	if err != nil {
		return nil, err
	}
	resp = &response.WebsiteAdminHomeInfo{
		ViewsCount:            10,
		MessageCount:          msgCount,
		UserCount:             userCount,
		ArticleCount:          articleCount,
		CategoryDTOList:       convertCategoryList(categories),
		TagDTOList:            convertTagList(tags),
		ArticleStatisticsList: convertArticleStatisticsList(articles),
		UniqueViewDTOList:     convertUniqueViewList(uniqueViews),
		ArticleRankDTOList:    convertArticleRankList(articles),
	}

	return resp, err
}

func (s *WebsiteService) GetSystemState(reqCtx *request.Context, req interface{}) (server *utils.Server, err error) {
	var sv utils.Server
	sv.Os = utils.InitOS()
	if sv.Cpu, err = utils.InitCPU(); err != nil {
		return &sv, err
	}
	if sv.Ram, err = utils.InitRAM(); err != nil {
		return &sv, err
	}
	if sv.Disk, err = utils.InitDisk(); err != nil {
		return &sv, err
	}

	return &sv, nil
}

func (s *WebsiteService) GetAboutMe(reqCtx *request.Context, req interface{}) (resp string, err error) {
	config, err := s.svcCtx.WebsiteConfigRepository.FindWebsiteConfig(reqCtx, &sqlx.Condition{Field: "key", Value: "about"})
	if err != nil {
		return "", err
	}

	return config.Config, err
}

func (s *WebsiteService) UpdateAboutMe(reqCtx *request.Context, req string) (resp string, err error) {
	config, err := s.svcCtx.WebsiteConfigRepository.FindWebsiteConfig(reqCtx, &sqlx.Condition{Field: "key", Value: "about"})
	if err != nil {
		return "", err
	}
	// 更新
	config.Config = req
	_, err = s.svcCtx.WebsiteConfigRepository.UpdateWebsiteConfig(reqCtx, config)
	if err != nil {
		return "", err
	}

	return config.Config, err
}

func (s *WebsiteService) GetConfig(reqCtx *request.Context, req *request.WebsiteConfigRequest) (resp string, err error) {
	config, err := s.svcCtx.WebsiteConfigRepository.FindWebsiteConfig(reqCtx, &sqlx.Condition{Field: "key", Value: req.Key})
	if err != nil {
		return "", err
	}

	return config.Config, err
}

func (s *WebsiteService) UpdateConfig(reqCtx *request.Context, req *request.WebsiteConfigRequest) (resp string, err error) {
	config, err := s.svcCtx.WebsiteConfigRepository.FindWebsiteConfig(reqCtx, &sqlx.Condition{Field: "key", Value: req.Key})
	if err != nil {
		return "", err
	}
	// 更新
	config.Config = req.Value
	_, err = s.svcCtx.WebsiteConfigRepository.UpdateWebsiteConfig(reqCtx, config)
	if err != nil {
		return "", err
	}

	return config.Config, err
}
