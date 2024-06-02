package mysql

import (
	"fmt"
	"go-auth/internal/app/apiserver/repo"
	"go-auth/internal/app/pkg/option"
	"gorm.io/gorm"
	"sync"
)

type dataset struct {
	db *gorm.DB
}

func (ds *dataset) User() repo.UserRepo {
	return newUserRepo(ds)
}

func (ds *dataset) Close() error {
	sqlDB, err := ds.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}

var (
	mysqlRepos repo.Factory
	once       sync.Once
)

// GetMySqlFactory create mysql factory with the given config.
func GetMySqlFactory(opts *option.MySqlOptions) (repo.Factory, error) {
	if opts == nil && mysqlRepos == nil {
		return nil, fmt.Errorf("failed to get mysql store fatory: mysql options is nil")
	}

	var err error
	var dbIns *gorm.DB

	once.Do(func() {
		dbIns, err = opts.NewClient()
		mysqlRepos = &dataset{db: dbIns}
	})

	if mysqlRepos == nil || err != nil {
		return nil, fmt.Errorf("failed to get mysql store fatory, mysqlFactory: %+v, error: %w", mysqlRepos, err)
	}

	return mysqlRepos, err
}
