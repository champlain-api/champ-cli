package shuttles

import (
	"github.com/champlain-api/champ-cli/cmd"
	"github.com/spf13/cobra"
)

func init() {
	cmd.RootCmd.AddCommand(shuttlesCommand)
	err := shuttlesCommand.MarkPersistentFlagRequired("api-key")
	if err != nil {
		return
	}

}

var shuttlesCommand = &cobra.Command{
	Use:   "shuttles",
	Short: "Shuttle management command.",

	Run: func(thisCmd *cobra.Command, args []string) {
		thisCmd.Help()
	},
}
