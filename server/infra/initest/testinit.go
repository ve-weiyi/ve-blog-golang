package initest

import (
	"context"
	"fmt"
	"log"
	"path"
	"time"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/fsnotify/fsnotify"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"

	"github.com/ve-weiyi/ve-blog-golang/server/config/properties"
	"github.com/ve-weiyi/ve-blog-golang/server/global"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/database"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/glog"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/glog/zaplog"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/jjwt"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/copyutil"
	"github.com/ve-weiyi/ve-blog-golang/server/utils/files"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
}

// @title						Swagger Example API
// @version					0.0.1
// @description				This is a sample Server pets
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						x-token
// @BasePath					/
func Init(configPath ...string) {
	InitConfig(configPath...)
	// 初始化zap日志库
	Zap()
	// 初始化gorm数据库
	Gorm()
	// 初始化redis服务
	Redis()
	// 初始化jwt服务
	JwtToken()
	//RBAC()
}

func InitConfig(configPath ...string) {
	log.Println("let's go")
	var filepath string
	if len(configPath) > 1 {
		filepath = configPath[0]
	} else {
		filepath = path.Join(global.GetRuntimeRoot() + "server/config.yaml")
	}
	global.VP = Viper(filepath) // 初始化Viper
}

func Viper(config string) *viper.Viper {

	v := viper.New()
	v.SetConfigFile(config)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		log.Println("config file changed:", e.Name)
		if err = v.Unmarshal(&global.CONFIG); err != nil {
			log.Println(err)
		}
	})
	if err = v.Unmarshal(&global.CONFIG); err != nil {
		log.Println(err)
	}

	// root 适配性 根据root位置去找到对应迁移位置,保证root路径有效
	//global.CONFIG.AutoCode.Root = global.GetRuntimeRoot()
	return v
}

func Zap() {
	err := files.MkDir(global.CONFIG.Zap.CacheDir)
	if err != nil {
		log.Println(err)
	}
	cfg := zaplog.ZapConfig{}

	copyutil.DeepCopyByJson(global.CONFIG.Zap, &cfg)

	//glog.ReplaceDefaultLogger(cfg)
	global.LOG = glog.NewGlogger(1, cfg)

	global.LOG.Infof("日志组件初始化成功！")
	return
}

func Gorm() {
	var cfg properties.DsnProvider

	cfg = &global.CONFIG.Mysql
	global.DB = database.Open(cfg)

	log.Printf("Mysql 数据库连接成功！%s", cfg.Dsn())
}

func Redis() {
	redisCfg := global.CONFIG.Redis
	address := redisCfg.Host + ":" + redisCfg.Port
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Printf("Redis 连接失败, err:%v", err)
		return
	}
	client.Set(context.Background(), "connect", time.Now().String(), -1)
	global.REDIS = client

	log.Printf("Redis 连接成功%v! address:%v db:%v", pong, address, redisCfg.DB)
}

func JwtToken() {
	global.JWT = &jjwt.JwtToken{
		SigningKey:  []byte(global.CONFIG.JWT.SigningKey),
		TokenPrefix: "",
		Issuer:      "blog",
	}
}

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

func RBAC() {
	GORM := global.DB
	if GORM == nil {
		panic("db is null")
	}
	//会自动创建数据库表并管理
	adapter, err := gormadapter.NewAdapterByDB(GORM)

	m, err := model.NewModelFromString(SubjectDomainObjectAction)
	if err != nil {
		log.Fatalln("字符串加载模型失败!", err)
	}

	syncedCachedEnforcer, _ := casbin.NewSyncedCachedEnforcer(m, adapter)
	syncedCachedEnforcer.SetExpireTime(60 * 60)
	_ = syncedCachedEnforcer.LoadPolicy()

	//global.RbacEnforcer = rbac.NewCachedEnforcer(global.DB)
}
