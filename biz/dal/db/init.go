package db

import (
	"crypto/tls"
	"database/sql"
	"fmt"
	"lonely-monitor/pkg/config"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once
	db   *gorm.DB
)

func GetDB() *gorm.DB {
	once.Do(initDB)
	return db
}

// initDB 初始化数据库连接
func initDB() {
	var err error
	dsn := config.GetConfig().Database.DSN()
	tlsCfg := config.GetConfig().Database.TLS
	if tlsCfg != "" {
		mysqlDriver.RegisterTLSConfig(tlsCfg, &tls.Config{
			MinVersion: tls.VersionTLS12,
			ServerName: config.GetConfig().Database.Host,
		})
	}

	sqlDB, err := sql.Open("mysql", dsn)

	db, err = gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(fmt.Errorf("连接数据库失败: %w", err))
	}

	//自动迁移
	// if err := DB.AutoMigrate(
	// 	&User{},
	// 	&AlertRecord{},
	// 	&ContactMethod{},
	// 	&NotifyRecord{},
	// ); err != nil {
	// 	hlog.Fatal("数据库迁移失败", err)
	// }

	hlog.Info("数据库初始化成功")
}
