package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/service/svc"
)

type WebsiteConfigService struct {
	svcCtx *svc.ServiceContext
}

func NewWebsiteConfigService(svcCtx *svc.ServiceContext) *WebsiteConfigService {
	return &WebsiteConfigService{
		svcCtx: svcCtx,
	}
}

func (s *WebsiteConfigService) GetAboutMe(reqCtx *request.Context, req interface{}) (resp string, err error) {
	config, err := s.svcCtx.WebsiteConfigRepository.FindWebsiteConfig(reqCtx, "about")
	if err != nil {
		return "", err
	}

	return config.Config, err
}

func (s *WebsiteConfigService) UpdateAboutMe(reqCtx *request.Context, req string) (resp string, err error) {
	config, err := s.svcCtx.WebsiteConfigRepository.FindWebsiteConfig(reqCtx, "about")
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
