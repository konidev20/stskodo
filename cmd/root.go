/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "stskodo",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	RunE: func(cmd *cobra.Command, args []string) error {
		endpoint := args[0]
		if endpoint == "" {
			return errors.New("endpoint missing")
		}

		accessKeyID := os.Getenv("AWS_ACCESS_KEY_ID")
		if accessKeyID == "" {
			return errors.New("ACCESS_KEY_ID missing")
		}

		secretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
		if secretAccessKey == "" {
			return errors.New("SECRET_ACCESS_KEY missing")
		}

		sess, err := session.NewSession(&aws.Config{
			Endpoint:    aws.String(endpoint),
			Region:      aws.String("us-east-1"),
			Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
		})
		if err != nil {
			return err
		}

		sts_client := sts.New(sess)

		params := &sts.GetSessionTokenInput{}

		req, resp := sts_client.GetSessionTokenRequest(params)
		err = req.Send()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Credentials)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stskodo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
