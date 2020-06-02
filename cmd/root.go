package cmd

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	accessKeyID string
	secretKey   string
	bucket      string

	rootCmd = &cobra.Command{
		Use:   "lamq",
		Short: "A job processor based on AWS lambdas",
	}

	startCmd = &cobra.Command{
		Use:   "start [file to process]",
		Short: "Start processing the file",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			start(awsSession(), args[0], bucket)
		},
	}

	pendingCmd = &cobra.Command{
		Use:   "pending",
		Short: "Lists all the pending jobs",
		Run: func(cmd *cobra.Command, args []string) {
			pending(awsSession(), bucket)
		},
	}

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all jobs",
		Run: func(cmd *cobra.Command, args []string) {
			list(awsSession(), bucket)
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

	accessKeyID = viper.GetString("accesskeyid")
	secretKey = viper.GetString("secretkey")
	bucket = viper.GetString("bucket")
}

// Starts a new AWS session with credentials
func awsSession() *session.Session {
	readConfig()

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-central-1"),
		Credentials: credentials.NewStaticCredentials(accessKeyID, secretKey, ""),
	})
	if err != nil {
		panic(fmt.Errorf("Fatal error connecting to AWS: %s\n", err))
	}

	return sess
}

func Execute() error {
	rootCmd.AddCommand(startCmd, listCmd, pendingCmd)
	return rootCmd.Execute()
}
