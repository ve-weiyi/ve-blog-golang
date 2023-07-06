package rbac

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-admin-store/server/api/blog/model/entity"
)

type CachedEnforcer struct {
	*casbin.SyncedCachedEnforcer
	DB            *gorm.DB
	whiteListPath map[string]string
}

func NewCachedEnforcer(db *gorm.DB) *CachedEnforcer {

	//会自动创建数据库表并管理
	adapter, err := gormadapter.NewAdapterByDB(db)

	m, err := model.NewModelFromString(SubjectDomainObjectAction)
	if err != nil {
		log.Fatalln("字符串加载模型失败!", err)
	}

	syncedCachedEnforcer, _ := casbin.NewSyncedCachedEnforcer(m, adapter)
	syncedCachedEnforcer.SetExpireTime(60 * 60)
	_ = syncedCachedEnforcer.LoadPolicy()

	return &CachedEnforcer{
		SyncedCachedEnforcer: syncedCachedEnforcer,
		DB:                   db,
		whiteListPath:        make(map[string]string),
	}
}

func (s *CachedEnforcer) LoadWhileList(apis []*entity.Api) error {

	whileList := make(map[string]string)
	for _, item := range apis {
		whileList[item.Path] = item.Method
	}

	s.whiteListPath = whileList
	return nil
}

func (s *CachedEnforcer) IsWhileList(path string, method string) bool {
	value, ok := s.whiteListPath[path]
	if !ok {
		return false
	}

	if value != method {
		return false
	}
	return true
}
