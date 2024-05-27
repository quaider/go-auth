package apiserver

import (
	"github.com/gin-gonic/gin"
	"go-auth/internal/app/apiserver/repo"
	"go-auth/internal/app/apiserver/repo/mysql"
	"go-auth/internal/app/pkg/option"
	"time"
)

type apiServer struct {
	*gin.Engine
}

func (s *apiServer) prepareToRun() {
	initRouter(s.Engine)

	// 初始化仓储实现为mysql
	sqlFactory, err := mysql.GetMySqlFactory(&option.MySqlOption{ // todo: 从配置文件解析
		Host:                  "127.0.0.1",
		Username:              "root",
		Password:              "123456",
		Database:              "auth",
		MaxIdleConnections:    20,
		MaxOpenConnections:    100,
		MaxConnectionLifeTime: time.Minute,
		Logger: &option.MySqlLogOption{
			LogLevel:                  1, // silence
			SlowThreshold:             3 * time.Second,
			IgnoreRecordNotFoundError: true,
		},
	})
	if err != nil {
		panic(err)
	}

	repo.SetClient(sqlFactory)
}

func Run() {
	eg := gin.New()
	eg.Use(gin.Recovery())

	s := &apiServer{eg}
	s.prepareToRun()

	_ = s.Engine.Run(":8000")
}
