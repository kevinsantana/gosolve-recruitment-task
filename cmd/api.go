package cmd

import (
	envconfig "github.com/kevinsantana/gosolve-recruitment-task/internal/config"
	"github.com/kevinsantana/gosolve-recruitment-task/internal/server"
	"github.com/kevinsantana/gosolve-recruitment-task/pkg/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run the http server.",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		log.WithField("project_version", version.PROJECT_VERSION)

		conf := envconfig.InitConfig(ctx)
		server.Run(ctx, server.HttpConfig{
			Cfg: conf,
		})
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)
}
