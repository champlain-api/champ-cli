package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var APIUrl string

func init() {
	rootCmd.PersistentFlags().StringVar(&APIUrl, "url", "http://localhost:3000", "URL for the API")
}

var rootCmd = &cobra.Command{
	Use:   "champ-cli",
	Short: "champ-cli is used to convert Champlain's API responses to our open source models",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("value is: %s", APIUrl)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
