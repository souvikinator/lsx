package cmd

import (
	"os"

	"github.com/gookit/color"
	"github.com/souvikinator/lsx/utils"
	"github.com/spf13/cobra"
)

var aliasCmd = &cobra.Command{
	Use:   "alias",
	Short: "lists all alias created by user",
	Run: func(cmd *cobra.Command, args []string) {

		if len(App.Alias) == 0 {
			utils.Warn("no alias found in records")
			utils.Warn("get started by using 'set-alias' command")
			os.Exit(0)
		}

		for alias, p := range App.Alias {
			color.Printf("%s --> %s\n", color.Magenta.Render(alias), color.Yellow.Render(p))
		}
		os.Exit(0)

	},
}
