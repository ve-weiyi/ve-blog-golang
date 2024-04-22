package svc

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/ve-weiyi/ve-blog-golang/server/infra/captcha"
	"github.com/ve-weiyi/ve-blog-golang/zero/repository/model"
	"github.com/ve-weiyi/ve-blog-golang/zero/rpc/internal/config"
)

type ServiceContext struct {
	Config            config.Config
	CaptchaRepository *captcha.CaptchaRepository

	UserAccountModel      model.UserAccountModel
	UserInformationModel  model.UserInformationModel
	UserLoginHistoryModel model.UserLoginHistoryModel
	RoleModel             model.RoleModel
	ApiModel              model.ApiModel
	MenuModel             model.MenuModel

	UserRoleModel model.UserRoleModel
	RoleApiModel  model.RoleApiModel
	RoleMenuModel model.RoleMenuModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := ConnectGorm(c.MysqlConf)
	if err != nil {
		panic(err)
	}

	rds, err := ConnectRedis(c.RedisConf)
	if err != nil {
		panic(err)
	}

	return &ServiceContext{
		Config:                c,
		UserAccountModel:      model.NewUserAccountModel(db, rds),
		UserInformationModel:  model.NewUserInformationModel(db, rds),
		UserLoginHistoryModel: model.NewUserLoginHistoryModel(db, rds),
		RoleModel:             model.NewRoleModel(db, rds),
		ApiModel:              model.NewApiModel(db, rds),
		MenuModel:             model.NewMenuModel(db, rds),
		UserRoleModel:         model.NewUserRoleModel(db, rds),
		RoleApiModel:          model.NewRoleApiModel(db, rds),
		RoleMenuModel:         model.NewRoleMenuModel(db, rds),
	}
}

func ConnectGorm(c config.MysqlConf) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s", c.Username, c.Password, c.Host, c.Port, c.Dbname, c.Config)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//PrepareStmt:            true, // 缓存预编译语句
		// 外键约束
		//DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// 表前缀
			TablePrefix: "",
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
		// gorm日志模式
		Logger: logger.Default.LogMode(logger.Info),
		//Logger: logger.Default,
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

	client.Set(context.Background(), fmt.Sprintf("redis:%s", pong), time.Now().String(), -1)
	return client, nil
}
