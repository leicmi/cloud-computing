package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	apiURL string

	rootCmd = &cobra.Command{
		Use:   "lamq",
		Short: "A job processor based on AWS lambdas",
	}

	startCmd = &cobra.Command{
		Use:   "start [file to process]",
		Short: "Start processing the file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			start(apiURL, args[0])
		},
	}

	pendingCmd = &cobra.Command{
		Use:   "pending",
		Short: "Lists all the pending jobs",
		Run: func(cmd *cobra.Command, args []string) {
			pending(apiURL)
		},
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all jobs",
		Run: func(cmd *cobra.Command, args []string) {
			list(apiURL)
		},
	}
)

// Reads config.yml and extracts credentials
func readConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			panic("Config file not found")
		} else {
			panic(fmt.Errorf("Fatal error reading config file: %s\n", err))
		}
	}

	apiURL = viper.GetString("apiurl")
}

func Execute() error {
	rootCmd.AddCommand(startCmd, listCmd, pendingCmd)
	return rootCmd.Execute()
}
