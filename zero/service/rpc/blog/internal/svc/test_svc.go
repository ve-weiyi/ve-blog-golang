package svc

import (
	"log"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/model"

	"github.com/ve-weiyi/ve-blog-golang/zero/service/rpc/blog/internal/config"
)

func NewTestConfig() config.Config {
	return config.Config{
		RpcServerConf: zrpc.RpcServerConf{
			ServiceConf: service.ServiceConf{
				Mode: "dev",
				Log: logx.LogConf{
					Mode:     "console",
					Encoding: "plain",
					Path:     "logs",
				},
			},
		},
		MysqlConf: config.MysqlConf{
			Host:     "127.0.0.1",
			Port:     "3306",
			Username: "root",
			Password: "mysql7914",
			Dbname:   "blog-veweiyi",
			Config:   "charset=utf8mb4&parseTime=True&loc=Local",
		},
		RedisConf: config.RedisConf{
			DB:       0,
			Host:     "127.0.0.1",
			Port:     "6379",
			Password: "redis7914",
		},
		RabbitMQConf: config.RabbitMQConf{
			Host:     "127.0.0.1",
			Port:     "5672",
			Username: "veweiyi",
			Password: "rabbitmq7914",
		},
	}
}

func NewTestServiceContext() *ServiceContext {
	log.SetFlags(log.LstdFlags | log.Llongfile)

	c := NewTestConfig()

	db, err := ConnectGorm(c.MysqlConf, c.Log)
	if err != nil {
		panic(err)
	}

	rds, err := ConnectRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config: c,
		Gorm:   db,

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

		CommentModel:    model.NewCommentModel(db, rds),
		RemarkModel:     model.NewRemarkModel(db, rds),
		FriendLinkModel: model.NewFriendLinkModel(db, rds),
		TalkModel:       model.NewTalkModel(db, rds),
		PhotoModel:      model.NewPhotoModel(db, rds),
		PhotoAlbumModel: model.NewPhotoAlbumModel(db, rds),
		PageModel:       model.NewPageModel(db, rds),

		OperationLogModel: model.NewOperationLogModel(db, rds),
		ChatRecordModel:   model.NewChatRecordModel(db, rds),
		UploadRecordModel: model.NewUploadRecordModel(db, rds),
	}
}
