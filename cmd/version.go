package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: "version",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Version 1.0 -- using API models from 0425d32")
	},
}
