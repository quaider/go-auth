package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

// Options defines options for mysql database.
type Options struct {
	Host                  string
	Username              string
	Password              string
	Database              string
	MaxIdleConnections    int
	MaxOpenConnections    int
	MaxConnectionLifeTime time.Duration
	Logger                struct {
		LogLevel                  logger.LogLevel
		SlowThreshold             time.Duration
		IgnoreRecordNotFoundError bool
	}
}

func New(opts *Options) (*gorm.DB, error) {
	dsn := fmt.Sprintf(`%s:%s@tcp(%s)/%s?charset=utf8&parseTime=%t&loc=%s`,
		opts.Username,
		opts.Password,
		opts.Host,
		opts.Database,
		true,
		"Local")

	// 先简单实现，后面再抽出去
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             opts.Logger.SlowThreshold,             // Slow SQL threshold
			LogLevel:                  opts.Logger.LogLevel,                  // Log level
			IgnoreRecordNotFoundError: opts.Logger.IgnoreRecordNotFoundError, // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,                                  // Don't include params in the SQL log
			Colorful:                  false,                                 // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 关闭表名复数形式
		},
	})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(opts.MaxOpenConnections)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(opts.MaxConnectionLifeTime)

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(opts.MaxIdleConnections)

	return db, nil
}
