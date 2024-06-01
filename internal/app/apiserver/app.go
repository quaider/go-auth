package apiserver

import (
	"go-auth/internal/app/pkg/option"
	"go-auth/pkg/app"
)

func NewApp() *app.App {
	config := &Config{}
	return app.NewApp("apiserver",
		"auth apiserver",
		app.WithConfig(config),
		app.WithStarter(run(config)),
	)
}

func run(config *Config) app.Starter {
	return func(shortName string) error {
		server := newApiServer(config)
		server.prepareToRun()

		return server.Run()
	}
}

var _ app.Config = (*Config)(nil)

type Config struct {
	Host         string               `json:"host" mapstructure:"host"`
	Port         int                  `json:"port" mapstructure:"port"`
	MysqlOptions *option.MySqlOptions `json:"mysql" mapstructure:"mysql"`
}
