package root

import (
	"os"
	"qqlx/base/conf"
	"qqlx/base/constant"
	"qqlx/cmd"
	"qqlx/cmd/root/init_data"
	"qqlx/cmd/root/run"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     conf.GetProjectName(),
	Long:    `go web framework`,
	Version: cmd.Version,
}

func init() {
	// 添加全局标志
	rootCmd.PersistentFlags().StringP(constant.FlagConfigPath, "C", "./config.yaml", "config file path")
	rootCmd.PersistentFlags().StringP(constant.FlagCasbinModePath, "M", "./model.conf", "casbin model file path")
	rootCmd.AddCommand(run.RunCmd, init_data.InitCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
