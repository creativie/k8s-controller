/*
Copyright © 2025 creativie <iam@creat.if.ua>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var logLevel string

var serverPort int
var finalPort int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "k8s-controller",
	Short: "HTTP server for k8s-controller",
	Long:  `For now it HTTP server with loglevel control and port configuration`,

	PreRun: func(cmd *cobra.Command, args []string) {
		level := parseLogLevel(logLevel)
		configureLogger(level)
		log.Info().Msgf("Log level set to: %s", logLevel)
		log.Info().Msg("This is an info log")
		log.Debug().Msg("This is a debug log")
		log.Trace().Msg("This is a trace log")
		log.Warn().Msg("This is a warn log")
		log.Error().Msg("This is an error log")
		fmt.Println("Welcome to k8s-controller-tutorial CLI!")

		viper.BindEnv("SERVER_PORT")
		env_port := viper.GetInt("SERVER_PORT")
		if env_port == 0 {
			log.Info().Msgf("Environment variable SERVER_PORT not set, using flag")
			finalPort = serverPort
		} else {
			log.Info().Msgf("Environment variable SERVER_PORT is set, using port %d", env_port)
			finalPort = env_port
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		StartHttpServer(finalPort)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVar(&logLevel, "log-level", "info", "Set the log level (trace, debug, info, warn, error)")
	rootCmd.Flags().IntVar(&serverPort, "port", 8080, "Set the server port (default is 8080)")
}

