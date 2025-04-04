package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var APIUrl string
var APIkey string
var Verbose bool

func init() {
	RootCmd.PersistentFlags().StringVar(&APIUrl, "url", "http://localhost:3000", "URL for the API")
	RootCmd.PersistentFlags().StringVarP(&APIkey, "api-key", "k", "", "API key")
	RootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose logging")

}

var RootCmd = &cobra.Command{
	Use:     "champ-cli",
	Short:   "champ-cli is used to convert Champlain's API responses to our open source models",
	Version: "1.0.0",

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
