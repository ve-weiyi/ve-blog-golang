package logic

import (
	"github.com/ve-weiyi/ve-blog-golang/server/api/model/request"
	"github.com/ve-weiyi/ve-blog-golang/server/api/service/svc"
	"github.com/ve-weiyi/ve-blog-golang/server/utils"
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
	config, err := s.svcCtx.WebsiteConfigRepository.FindWebsiteConfigByKey(reqCtx, "about")
	if err != nil {
		return "", err
	}

	return config.Config, err
}

func (s *WebsiteConfigService) UpdateAboutMe(reqCtx *request.Context, req string) (resp string, err error) {
	config, err := s.svcCtx.WebsiteConfigRepository.FindWebsiteConfigByKey(reqCtx, "about")
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

func (s *WebsiteConfigService) GetConfig(reqCtx *request.Context, req *request.WebsiteConfigRequest) (resp string, err error) {
	config, err := s.svcCtx.WebsiteConfigRepository.FindWebsiteConfigByKey(reqCtx, req.Key)
	if err != nil {
		return "", err
	}

	return config.Config, err
}

func (s *WebsiteConfigService) UpdateConfig(reqCtx *request.Context, req *request.WebsiteConfigRequest) (resp string, err error) {
	config, err := s.svcCtx.WebsiteConfigRepository.FindWebsiteConfigByKey(reqCtx, req.Key)
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

func (s *WebsiteConfigService) GetSystemState(reqCtx *request.Context, req interface{}) (server *utils.Server, err error) {
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
