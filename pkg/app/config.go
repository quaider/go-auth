package app

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
)

const configFlagName = "config"

var cfgFile string

func init() {
	pflag.StringVarP(&cfgFile, "config", "c", "", "配置文件路径")
}

type Config interface {
}

func addConfigFlag(configName string, flagSet *pflag.FlagSet) {
	flagSet.AddFlag(pflag.Lookup(configFlagName)) // 另一种绑定viper配置和flag的方式
	viper.AutomaticEnv()

	cobra.OnInitialize(func() {
		if cfgFile != "" {
			viper.SetConfigFile(cfgFile)
		} else {
			viper.AddConfigPath("./config")
			viper.SetConfigType("yaml")
			viper.SetConfigName(configName)
		}

		if err := viper.ReadInConfig(); err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Error: failed to read configuration file(%s): %v\n", cfgFile, err)
			os.Exit(1)
		}
	})
}
