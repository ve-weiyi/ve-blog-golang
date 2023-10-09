package database

import (
	"log"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/ve-weiyi/ve-blog-golang/server/config/properties"
	"github.com/ve-weiyi/ve-blog-golang/server/infra/database/ormlog"
)

func Open(cfg properties.DsnProvider) *gorm.DB {
	dialector := cfg.GetConnector()

	db, err := gorm.Open(*dialector, Config(cfg))
	if err != nil {
		log.Printf("GORM 数据库连接失败: %v", err)
		return nil
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Printf("SQL 数据库连接失败: %v", err)
		return nil
	}

	sqlDB.SetMaxIdleConns(cfg.GetConfig().MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.GetConfig().MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db
}

// Config gorm 自定义配置
// Author [SliverHorn](https://github.com/SliverHorn)
func Config(cfg properties.DsnProvider) *gorm.Config {

	config := &gorm.Config{
		//PrepareStmt:            true, // 缓存预编译语句
		// gorm日志模式：silent
		//Logger: logger.Default.LogMode(logger.Info),
		// 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: cfg.GetConfig().Prefix,
			// 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	}

	_default := logger.New(ormlog.NewWriter(log.New(os.Stdout, "\r\n", log.LstdFlags)), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false, // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  true,  // 彩色打印
	})
	switch cfg.GetConfig().LogMode {
	case "silent", "Silent":
		config.Logger = _default.LogMode(logger.Silent)
	case "error", "Error":
		config.Logger = _default.LogMode(logger.Error)
	case "warn", "Warn":
		config.Logger = _default.LogMode(logger.Warn)
	case "info", "Info":
		config.Logger = _default.LogMode(logger.Info)
	default:
		config.Logger = _default.LogMode(logger.Info)
	}
	return config
}
