package logic

import (
	jsoniter "github.com/json-iterator/go"

	"github.com/ve-weiyi/ve-blog-golang/kit/utils/system"

	"github.com/ve-weiyi/ve-blog-golang/server/api/model/entity"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/response"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
)

type WebsiteService struct {
	svcCtx *svc.ServiceContext
}

func NewWebsiteService(svcCtx *svc.ServiceContext) *WebsiteService {
	return &WebsiteService{
		svcCtx: svcCtx,
	}
}

func (l *WebsiteService) GetBlogHomeInfo(reqCtx *request.Context, data interface{}) (resp *response.BlogHomeInfo, err error) {
	articleCount, _ := l.svcCtx.ArticleRepository.Count(reqCtx, "`is_delete` = ?", entity.False)
	categoryCount, _ := l.svcCtx.CategoryRepository.Count(reqCtx, "")
	tagCount, _ := l.svcCtx.TagRepository.Count(reqCtx, "")
	pages, _ := l.svcCtx.PageRepository.FindALL(reqCtx, "")
	config, err := l.svcCtx.WebsiteConfigRepository.First(reqCtx, "`key` = ?", "website_config")

	resp = &response.BlogHomeInfo{
		ArticleCount:  articleCount,
		CategoryCount: categoryCount,
		TagCount:      tagCount,
		ViewsCount:    "1",
		PageList:      convertPageList(pages),
	}

	jsoniter.Unmarshal([]byte(config.Config), &resp.WebsiteConfig)
	return resp, err
}

func (l *WebsiteService) GetAdminHomeInfo(reqCtx *request.Context, data interface{}) (resp *response.AdminHomeInfo, err error) {
	// 查询消息数量
	msgCount, err := l.svcCtx.RemarkRepository.Count(reqCtx, "")
	if err != nil {
		return nil, err
	}

	// 查询用户数量
	userCount, err := l.svcCtx.UserAccountRepository.Count(reqCtx, "")
	if err != nil {
		return nil, err
	}

	// 查询文章数量
	articles, err := l.svcCtx.ArticleRepository.FindALL(reqCtx, "")
	if err != nil {
		return nil, err
	}

	// 查询分类数量
	categories, err := l.svcCtx.CategoryRepository.FindALL(reqCtx, "")
	if err != nil {
		return nil, err
	}

	// 查询标签数量
	tags, err := l.svcCtx.TagRepository.FindALL(reqCtx, "")
	if err != nil {
		return nil, err
	}

	uniqueViews, err := l.svcCtx.UniqueViewRepository.FindALL(reqCtx, "")
	if err != nil {
		return nil, err
	}

	articleCount, err := l.svcCtx.ArticleRepository.Count(reqCtx, "")
	if err != nil {
		return nil, err
	}
	resp = &response.AdminHomeInfo{
		ViewsCount:            10,
		MessageCount:          msgCount,
		UserCount:             userCount,
		ArticleCount:          articleCount,
		CategoryList:          convertCategoryList(categories),
		TagList:               convertTagList(tags),
		ArticleStatisticsList: convertArticleStatisticsList(articles),
		UniqueViewList:        convertUniqueViewList(uniqueViews),
		ArticleRankList:       convertArticleRankList(articles),
	}

	return resp, err
}

func (l *WebsiteService) GetSystemState(reqCtx *request.Context, req interface{}) (server *system.Server, err error) {
	var sv system.Server
	sv.Os = system.InitOS()
	if sv.Cpu, err = system.InitCPU(); err != nil {
		return &sv, err
	}
	if sv.Ram, err = system.InitRAM(); err != nil {
		return &sv, err
	}
	if sv.Disk, err = system.InitDisk(); err != nil {
		return &sv, err
	}

	return &sv, nil
}

func (l *WebsiteService) GetAboutMe(reqCtx *request.Context, req interface{}) (resp *response.AboutMeResp, err error) {
	config, err := l.svcCtx.WebsiteConfigRepository.First(reqCtx, "`key` = ?", "about_me")
	if err != nil {
		return nil, err
	}

	jsoniter.Unmarshal([]byte(config.Config), &resp)
	return resp, nil
}

func (l *WebsiteService) UpdateAboutMe(reqCtx *request.Context, req *request.AboutMeReq) (resp *response.AboutMeResp, err error) {
	config, err := l.svcCtx.WebsiteConfigRepository.First(reqCtx, "`key` = ?", "about_me")
	if err != nil {
		return nil, err
	}
	// 更新
	config.Config = req.Content
	_, err = l.svcCtx.WebsiteConfigRepository.Update(reqCtx, config)
	if err != nil {
		return nil, err
	}

	jsoniter.Unmarshal([]byte(config.Config), &resp)
	return resp, nil
}

func (l *WebsiteService) GetWebsiteConfig(reqCtx *request.Context, req interface{}) (resp *response.WebsiteConfigDTO, err error) {
	config, err := l.svcCtx.WebsiteConfigRepository.First(reqCtx, "`key` = ?", "website_config")
	if err != nil {
		return nil, err
	}

	jsoniter.Unmarshal([]byte(config.Config), &resp)
	return resp, nil
}

func (l *WebsiteService) UpdateWebsiteConfig(reqCtx *request.Context, req *request.WebsiteConfigDTO) (resp *response.WebsiteConfigDTO, err error) {
	config, err := l.svcCtx.WebsiteConfigRepository.First(reqCtx, "`key` = ?", "website_config")
	if err != nil {
		return nil, err
	}

	// 更新
	config.Config, _ = jsoniter.MarshalToString(req)
	_, err = l.svcCtx.WebsiteConfigRepository.Update(reqCtx, config)
	if err != nil {
		return nil, err
	}

	jsoniter.Unmarshal([]byte(config.Config), &resp)
	return resp, nil
}

func (l *WebsiteService) GetConfig(reqCtx *request.Context, req *request.WebsiteConfigReq) (resp string, err error) {
	config, err := l.svcCtx.WebsiteConfigRepository.First(reqCtx, "`key` = ?", req.Key)
	if err != nil {
		return "", err
	}

	return config.Config, err
}

func (l *WebsiteService) UpdateConfig(reqCtx *request.Context, req *request.WebsiteConfigReq) (resp string, err error) {
	config, err := l.svcCtx.WebsiteConfigRepository.First(reqCtx, "`key` = ?", req.Key)
	if err != nil {
		return "", err
	}
	// 更新
	config.Config = req.Value
	_, err = l.svcCtx.WebsiteConfigRepository.Update(reqCtx, config)
	if err != nil {
		return "", err
	}

	return config.Config, err
}
