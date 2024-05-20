package global

import (
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"github.com/orca-zhang/ecache"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/chatgpt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/upload"
	"github.com/ve-weiyi/ve-blog-golang/server/config"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/rbac"
)

var (
	VP        *viper.Viper
	DB        *gorm.DB
	DBList    map[string]*gorm.DB
	REDIS     *redis.Client
	CONFIG    config.Config
	JWT       *jjwt.JwtToken
	LOG       *glog.Glogger
	EmailMQ   *rabbitmq.RabbitmqConn
	Uploader  upload.Uploader
	AIChatGPT *chatgpt.AIChatGPT
	//Timer    timer.Timer = timer.NewTimerTask()
	//Concurrency_Control             = &singleflight.Group{}

	BlackCache *ecache.Cache
	lock       sync.RWMutex

	//RBAC角色访问控制器
	Permission rbac.RbacHolder
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则 使用默认db panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := DBList[dbname]
	if !ok || db == nil {
		return DB
		//panic("db no init")
	}
	return db
}

func GetRuntimeRoot() string {
	//获得当前方法运行文件名
	_, filename, _, _ := runtime.Caller(1)
	//找到 resource/language目录
	src := "server"
	index := strings.Index(filename, src)
	root := filename[:index]
	return root
}

func GetCurrentDir() string {
	// 获取当前函数的调用信息
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Failed to get caller information")
	}
	// 根据文件名获取绝对路径
	dir := filepath.Dir(filename)
	absDir, err := filepath.Abs(dir)
	if err != nil {
		panic(err)
	}
	return absDir
}
