package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var APIUrl string
var APIkey string

func init() {
	rootCmd.PersistentFlags().StringVar(&APIUrl, "url", "http://localhost:3000", "URL for the API")
	rootCmd.PersistentFlags().StringVarP(&APIkey, "api-key", "k", "", "API key")
}

var rootCmd = &cobra.Command{
	Use:     "champ-cli",
	Short:   "champ-cli is used to convert Champlain's API responses to our open source models",
	Version: "1 -- using model definitions from commit 0425d32",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("champ-cli tool")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
