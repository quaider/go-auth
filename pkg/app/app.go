package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// App 代表一个应用程序
type App struct {
	shortName string
	name      string
	config    Config
	starter   Starter
	rootCmd   *cobra.Command
}

// Option 定义App的初始化参数
type Option func(*App)

func WithConfig(config Config) Option {
	return func(app *App) {
		app.config = config
	}
}

func WithStarter(starter Starter) Option {
	return func(app *App) {
		app.starter = starter
	}
}

type Starter func(shortName string) error

// NewApp 初始化一个App
func NewApp(shortName, name string, opts ...Option) *App {
	app := &App{
		shortName: shortName,
		name:      name,
	}

	for _, opt := range opts {
		opt(app)
	}

	// 集成cobra命令行框架
	app.buildRootCmd()

	return app
}

func (app *App) runE(cmd *cobra.Command, args []string) error {
	printWorkingDir()

	// viper配置绑定flags
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}

	if err := viper.Unmarshal(&app.config); err != nil {
		return err
	}

	if app.starter != nil {
		return app.starter(app.shortName)
	}

	return nil
}

// Run 启动应用程序
func (app *App) Run() {
	if err := app.rootCmd.Execute(); err != nil {
		fmt.Printf("%s 启动失败, %v\n", app.name, err)
		os.Exit(1) // 非正常退出
	}
}

// buildRootCmd 构建 cobra 根命令
func (app *App) buildRootCmd() {
	cmd := &cobra.Command{
		Use:   app.shortName,
		Short: app.name,
	}

	if app.starter != nil {
		cmd.RunE = app.runE
	}

	addConfigFlag(app.shortName, cmd.Flags())

	app.rootCmd = cmd
}

func printWorkingDir() {
	wd, _ := os.Getwd()
	fmt.Printf("WorkingDir: %s\n", wd)
}
