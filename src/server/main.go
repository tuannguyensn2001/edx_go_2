package main

import (
	"edx_go_2/src/cmd"
	"edx_go_2/src/config"
	"edx_go_2/src/middlewares"
	"edx_go_2/src/routes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	config.Load()

	rootCmd := cmd.GetRoot(config.Conf)

	rootCmd.AddCommand(&cobra.Command{
		Use: "server",
		Run: func(cmd *cobra.Command, args []string) {
			
			logrus.SetOutput(os.Stdout)

			conf := config.Conf
			r := gin.Default()

			r.Use(middlewares.Recover)
			r.Use(middlewares.Cors)

			routes.DeclareRoute(r)

			err := r.Run(":" + conf.Port)
			if err != nil {
				return
			}
		},
	})

	err := rootCmd.Execute()
	if err != nil {
		return
	}

}
