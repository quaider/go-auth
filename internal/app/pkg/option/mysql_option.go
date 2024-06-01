package option

import (
	"go-auth/pkg/db"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

type MySqlOptions struct {
	Host                  string           `json:"host,omitempty"                     mapstructure:"host"`
	Username              string           `json:"username,omitempty"                 mapstructure:"username"`
	Password              string           `json:"-"                                  mapstructure:"password"`
	Database              string           `json:"database"                           mapstructure:"database"`
	MaxIdleConnections    int              `json:"max-idle-connections,omitempty"     mapstructure:"max-idle-connections"`
	MaxOpenConnections    int              `json:"max-open-connections,omitempty"     mapstructure:"max-open-connections"`
	MaxConnectionLifeTime int              `json:"max-connection-life-time,omitempty" mapstructure:"max-connection-life-time"`
	Logger                *MySqlLogOptions `json:"logger,omitempty"                   mapstructure:"logger"`
}

type MySqlLogOptions struct {
	LogLevel                  int           `json:"log-level"                                  mapstructure:"log-level"`
	SlowThreshold             time.Duration `json:"slow-threshold,omitempty"                   mapstructure:"slow-threshold"`
	IgnoreRecordNotFoundError bool          `json:"ignore-record-not-found-error,omitempty"    mapstructure:"ignore-record-not-found-error"`
}

// NewClient 创建mysql客户端
func (o *MySqlOptions) NewClient() (*gorm.DB, error) {
	opts := &db.Options{
		Host:                  o.Host,
		Username:              o.Username,
		Password:              o.Password,
		Database:              o.Database,
		MaxIdleConnections:    o.MaxIdleConnections,
		MaxOpenConnections:    o.MaxOpenConnections,
		MaxConnectionLifeTime: time.Duration(o.MaxConnectionLifeTime) * time.Second,
		Logger: struct {
			LogLevel                  logger.LogLevel
			SlowThreshold             time.Duration
			IgnoreRecordNotFoundError bool
		}{
			LogLevel:                  logger.LogLevel(o.Logger.LogLevel),
			SlowThreshold:             o.Logger.SlowThreshold,
			IgnoreRecordNotFoundError: o.Logger.IgnoreRecordNotFoundError,
		},
	}

	return db.New(opts)
}
