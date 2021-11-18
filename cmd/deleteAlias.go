package cmd

import (
	"os"

	"github.com/manifoldco/promptui"
	"github.com/souvikinator/lsx/utils"
	"github.com/spf13/cobra"
)

// deleteAliasCmd represents the deleteAlias command
var removeAliasCmd = &cobra.Command{
	Use:   "remove-alias <alias>",
	Short: "remove existing alias from records",
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			utils.Err("no alias specified")
		}

		if len(App.Alias) == 0 {
			utils.Warn("no alias in records! set alias by using 'set-alias' command.")
			os.Exit(0)
		}

		deleteAlias(args[0])
		utils.Info("deleted")
	},
}

func deleteAlias(name string) {

	_, ok := App.Alias[name]
	if !ok {
		utils.Err("alias '" + name + "' not found in records. List all alias using 'alias' command")
	}

	prompt := promptui.Prompt{
		Label:     "are you sure, you want to delete this alias?",
		IsConfirm: true,
	}

	_, err := prompt.Run()

	if err != nil {
		utils.Info("aborting delete!")
		os.Exit(1)
	}

	delete(App.Alias, name)
	App.WriteAliasFile()

}
