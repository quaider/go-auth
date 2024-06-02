package apiserver

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-auth/internal/app/apiserver/repo"
	"go-auth/internal/app/apiserver/repo/mysql"
	rdb2 "go-auth/pkg/rdb"
)

type apiServer struct {
	*gin.Engine
	config *Config
}

func newApiServer(config *Config) *apiServer {
	eg := gin.New()
	eg.Use(gin.Recovery())

	return &apiServer{
		Engine: eg,
		config: config,
	}
}

func (server *apiServer) prepareToRun() {
	initRouter(server.Engine)

	// 初始化仓储实现为mysql
	sqlFactory, err := mysql.GetMySqlFactory(server.config.MysqlOptions)
	if err != nil {
		panic(err)
	}

	repo.SetClient(sqlFactory)

	// 初始化 redis
	rdb2.NewRdbCli(server.config.Redis)
}

func (server *apiServer) Run() error {
	return server.Engine.Run(fmt.Sprintf("%s:%d", server.config.Host, server.config.Port))
}
