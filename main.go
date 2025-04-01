package main

import (
	"github.com/champlain-api/champ-cli/cmd"
	_ "github.com/champlain-api/champ-cli/cmd/housing"
	_ "github.com/champlain-api/champ-cli/cmd/shuttles"
)

func main() {
	cmd.Execute()
}
