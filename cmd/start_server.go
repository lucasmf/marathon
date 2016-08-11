package cmd

import (
	"git.topfreegames.com/topfreegames/marathon/api"
	"github.com/spf13/cobra"
	"github.com/uber-go/zap"
)

var host string
var port int
var debug bool

// startCmd represents the start command
var startServerCmd = &cobra.Command{
	Use:   "start-server",
	Short: "starts the marathon API server",
	Long:  `Starts marathon server with the specified arguments. You can use environment variables to override configuration keys.`,
	Run: func(cmd *cobra.Command, args []string) {
		ll := zap.InfoLevel
		if debug {
			ll = zap.DebugLevel
		}
		l := zap.NewJSON(ll, zap.AddCaller())

		cmdL := l.With(
			zap.String("source", "startServerCmd"),
			zap.String("operation", "Run"),
			zap.String("host", host),
			zap.Int("port", port),
			zap.Bool("debug", debug),
		)

		cmdL.Debug("Creating application...")
		application := api.GetApplication(
			host,
			port,
			ConfigFile,
			debug,
			l,
		)
		cmdL.Debug("Application created successfully.")

		cmdL.Debug("Starting application...")
		application.Start()
	},
}

func init() {
	RootCmd.AddCommand(startServerCmd)

	startServerCmd.Flags().StringVarP(&host, "bind", "b", "0.0.0.0", "Host to bind marathon to")
	startServerCmd.Flags().IntVarP(&port, "port", "p", 8888, "Port to bind marathon to")
	startServerCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Debug mode")
}