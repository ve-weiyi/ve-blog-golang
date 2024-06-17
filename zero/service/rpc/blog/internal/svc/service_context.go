package svc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-zero/core/collection"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/gitee"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/github"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/kit/infra/captcha"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/constant"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/gormlogger"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/mail"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/feishu"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/qq"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/oauth/weibo"
	"github.com/ve-weiyi/ve-blog-golang/kit/infra/rabbitmq"
	"github.com/ve-weiyi/ve-blog-golang/zero/internal/gormlogx"
	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/config"
)

type ServiceContext struct {
	Config        config.Config
	Gorm          *gorm.DB
	Redis         *redis.Client
	LocalCache    *collection.Cache
	CaptchaHolder *captcha.CaptchaHolder
	EmailMQ       *rabbitmq.RabbitmqConn
	Oauth         map[string]oauth.Oauth

	UserAccountModel      model.UserAccountModel
	UserOauthModel        model.UserOauthModel
	UserLoginHistoryModel model.UserLoginHistoryModel
	RoleModel             model.RoleModel
	ApiModel              model.ApiModel
	MenuModel             model.MenuModel
	UserRoleModel         model.UserRoleModel
	RoleApiModel          model.RoleApiModel
	RoleMenuModel         model.RoleMenuModel

	// blog models
	WebsiteConfigModel model.WebsiteConfigModel
	ArticleModel       model.ArticleModel
	CategoryModel      model.CategoryModel
	TagModel           model.TagModel
	ArticleTagModel    model.ArticleTagModel

	CommentModel      model.CommentModel
	RemarkModel       model.RemarkModel
	FriendModel       model.FriendModel
	TalkModel         model.TalkModel
	PhotoModel        model.PhotoModel
	AlbumModel        model.AlbumModel
	BannerModel       model.BannerModel
	VisitHistoryModel model.VisitHistoryModel

	OperationLogModel model.OperationLogModel
	ChatRecordModel   model.ChatRecordModel
	UploadRecordModel model.UploadRecordModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := ConnectGorm(c.MysqlConf, c.Log)
	if err != nil {
		panic(err)
	}

	rds, err := ConnectRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}

	mq, err := ConnectRabbitMq(c.RabbitMQConf)
	if err != nil {
		panic(err)
	}

	go SubscribeMessage(c)

	cache, err := collection.NewCache(60 * time.Minute)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:                c,
		Gorm:                  db,
		Redis:                 rds,
		LocalCache:            cache,
		CaptchaHolder:         captcha.NewCaptchaHolder(captcha.WithRedisStore(rds)),
		EmailMQ:               mq,
		Oauth:                 InitOauth(c.OauthConfList),
		UserAccountModel:      model.NewUserAccountModel(db, rds),
		UserOauthModel:        model.NewUserOauthModel(db, rds),
		UserLoginHistoryModel: model.NewUserLoginHistoryModel(db, rds),
		RoleModel:             model.NewRoleModel(db, rds),
		ApiModel:              model.NewApiModel(db, rds),
		MenuModel:             model.NewMenuModel(db, rds),
		UserRoleModel:         model.NewUserRoleModel(db, rds),
		RoleApiModel:          model.NewRoleApiModel(db, rds),
		RoleMenuModel:         model.NewRoleMenuModel(db, rds),

		// blog models
		WebsiteConfigModel: model.NewWebsiteConfigModel(db, rds),
		ArticleModel:       model.NewArticleModel(db, rds),
		CategoryModel:      model.NewCategoryModel(db, rds),
		TagModel:           model.NewTagModel(db, rds),
		ArticleTagModel:    model.NewArticleTagModel(db, rds),

		CommentModel:      model.NewCommentModel(db, rds),
		RemarkModel:       model.NewRemarkModel(db, rds),
		FriendModel:       model.NewFriendModel(db, rds),
		TalkModel:         model.NewTalkModel(db, rds),
		PhotoModel:        model.NewPhotoModel(db, rds),
		AlbumModel:        model.NewAlbumModel(db, rds),
		BannerModel:       model.NewBannerModel(db, rds),
		VisitHistoryModel: model.NewVisitHistoryModel(db, rds),

		OperationLogModel: model.NewOperationLogModel(db, rds),
		ChatRecordModel:   model.NewChatRecordModel(db, rds),
		UploadRecordModel: model.NewUploadRecordModel(db, rds),
	}
}

func ConnectGorm(c config.MysqlConf, l logx.LogConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Dbname, c.Config)

	var lg logger.Interface
	if l.Mode == "console" && l.Encoding == "plain" {
		// 跟随gorm的日志输出格式
		lg = logger.New(
			gormlogger.NewGormWriter(),
			logger.Config{
				SlowThreshold:             500 * time.Millisecond, // 慢 SQL 阈值，超过会提前结束
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false, // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,  // 彩色打印
				ParameterizedQueries:      false, // 使用参数化查询 (true时，会将参数值替换为?)
			})
	} else {
		// 跟随go-zero的日志输出格式
		lg = gormlogx.New(
			logger.Config{
				SlowThreshold:             500 * time.Millisecond, // 慢 SQL 阈值，超过会提前结束
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: false, // 忽略ErrRecordNotFound（记录未找到）错误
				Colorful:                  true,  // 彩色打印
				ParameterizedQueries:      false, // 使用参数化查询 (true时，会将参数值替换为?)
			},
		)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// PrepareStmt:            true, // 缓存预编译语句
		// 外键约束
		// DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 表前缀
			TablePrefix: "",
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
		// gorm日志模式
		Logger: lg,
		// Logger: logger.Default,
	})

	if err != nil {
		return nil, fmt.Errorf("GORM 数据库连接失败: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("SQL 数据库连接失败: %v", err)
	}

	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(64)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(64)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Minute)

	return db, nil
}

func ConnectRedis(c config.RedisConf) (*redis.Client, error) {
	address := c.Host + ":" + c.Port
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Username: "",
		Password: c.Password, // no password set
		DB:       c.DB,       // use default DB
	})

	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, fmt.Errorf("redis 连接失败: %v", err)
	}

	client.Set(context.Background(), fmt.Sprintf("redis:rpc:%s", pong), time.Now().String(), -1)
	return client, nil
}

func ConnectRabbitMq(c config.RabbitMQConf) (*rabbitmq.RabbitmqConn, error) {
	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", c.Username, c.Password, c.Host, c.Port)

	// 消息发布者只需要声明交换机
	mq := rabbitmq.NewRabbitmqConn(url,
		rabbitmq.Exchange(rabbitmq.ExchangeOptions{
			Name:    constant.EmailExchange,
			Type:    rabbitmq.ExchangeTypeFanout,
			Durable: true,
		}),
		rabbitmq.DisableAutoAck(),
		rabbitmq.Requeue(),
	)

	err := mq.Connect(nil)
	if err != nil {
		return nil, fmt.Errorf("rabbitmq 初始化失败: %v", err)
	}

	return mq, nil
}

// 订阅消息
func SubscribeMessage(c config.Config) {
	r := c.RabbitMQConf

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/", r.Username, r.Password, r.Host, r.Port)
	// 消息订阅者需要声明交换机和队列
	mq := rabbitmq.NewRabbitmqConn(url,
		rabbitmq.Queue(rabbitmq.QueueOptions{
			Name:    constant.EmailQueue,
			Durable: true,
			Args:    nil,
		}),
		rabbitmq.Exchange(rabbitmq.ExchangeOptions{
			Name:    constant.EmailExchange,
			Type:    rabbitmq.ExchangeTypeFanout,
			Durable: true,
		}),
		rabbitmq.Key("email"),
	)
	err := mq.Connect(nil)
	if err != nil {
		log.Fatal("rabbitmq 初始化失败!", err)
	}

	e := c.EmailConf
	emailSender := mail.NewEmailSender(
		mail.WithHost(e.Host),
		mail.WithPort(e.Port),
		mail.WithUsername(e.Username),
		mail.WithPassword(e.Password),
		mail.WithNickname(e.Nickname),
		mail.WithDeliver(e.Deliver),
		mail.WithIsSSL(e.IsSSL),
	)

	// 订阅消息队列，发送邮件
	err = mq.SubscribeMessage(func(message []byte) (err error) {
		var msg mail.EmailMessage
		err = json.Unmarshal(message, &msg)
		if err != nil {
			return err
		}

		err = emailSender.SendEmailMessage(msg)
		if err != nil {
			log.Println("邮件发送失败!", err)
		}
		return err
	})
	if err != nil {
		log.Fatal("订阅消息失败!", err)
	}
}

func InitOauth(c map[string]config.OauthConf) map[string]oauth.Oauth {
	var om = make(map[string]oauth.Oauth)

	for k, v := range c {
		conf := &oauth.AuthConfig{
			ClientId:     v.ClientId,
			ClientSecret: v.ClientSecret,
			RedirectUri:  v.RedirectUri,
		}
		switch k {
		case "qq":
			auth := qq.NewAuthQq(conf)
			om["qq"] = auth
		case "weibo":
			auth := weibo.NewAuthWb(conf)
			om["weibo"] = auth
		case "feishu":
			auth := feishu.NewAuthFeishu(conf)
			om["feishu"] = auth
		case "github":
			auth := github.NewAuthGithub(conf)
			om["github"] = auth
		case "gitee":
			auth := gitee.NewAuthGitee(conf)
			om["gitee"] = auth
		}
	}
	return om
}
