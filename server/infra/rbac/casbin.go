package rbac

import (
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/server/api/blog/model/entity"
)

// 访问实体 (Subject)，访问资源 (Object) 和访问方法 (Action)
const SubjectObjectAction = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
`

// 访问实体 (Subject)，领域(Domain)，访问资源 (Object) 和访问方法 (Action)
const SubjectDomainObjectAction = `
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.act
`

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

// 删除CasbinRule记录
func (s *CachedEnforcer) DeleteRolePolicy(roleName string, domain string) (result interface{}, err error) {
	rbac := s.SyncedCachedEnforcer
	return rbac.RemoveFilteredNamedPolicy("p", 0, roleName, domain)
}

// 添加CasbinRule记录
func (s *CachedEnforcer) AddRolePolicy(roleName string, domain string, apis []*entity.Api) (result interface{}, err error) {
	var policies [][]string
	for _, item := range apis {
		if item.Path == "" || item.Method == "" {
			continue
		}
		policies = append(policies, []string{roleName, domain, item.Path, item.Method})

	}

	rbac := s.SyncedCachedEnforcer
	_, err = rbac.AddPolicies(policies)
	if err != nil {
		return nil, err
	}
	return nil, rbac.SavePolicy()
}
