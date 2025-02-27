package run

import (
	"context"
	"qqlx/base/conf"
	"qqlx/base/constant"
	"qqlx/base/helpers"
	"qqlx/base/logger"
	"qqlx/cmd"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var RunCmd = &cobra.Command{
	Use:   "run",
	Long:  "start the go web framework",
	Short: "start the go web framework",
	PreRun: func(cmd *cobra.Command, args []string) {
		helpers.PreRun(cmd)
	},
	Run: func(cmd *cobra.Command, args []string) {
		var (
			cf     string
			casbin string
			err    error
		)
		cf, err = cmd.Flags().GetString(constant.FlagConfigPath)
		if err != nil {
			panic(err)
		}
		casbin, err = cmd.Flags().GetString(constant.FlagCasbinModePath)
		if err != nil {
			panic(err)
		}
		runApp(cf, casbin)
	},
}

func runApp(configPath, casbinModePath string) {
	err := conf.LoadConfig(configPath)
	if err != nil {
		panic(err)
	}
	logger.InitLogger()
	defer zap.S().Sync()
	ctx := context.Background()
	app, cleanup, err := cmd.InitApplication(ctx, casbinModePath)
	defer cleanup()
	if err != nil {
		panic(err)
	}
	if err := app.Run(ctx); err != nil {
		panic(err)
	}
}
