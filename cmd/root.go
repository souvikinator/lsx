package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/souvikinator/lsx/lsx"
	"github.com/spf13/cobra"
)

var App lsx.Lsx

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lsx",
	Short: " A command line utility written in golang which beautifies and extends ls command",
	Run: func(cmd *cobra.Command, args []string) {

		var path, prev, moveTo, startPath string
		// flags
		// App.Init()
		// App.DisplayIcons()
		// App.AllMode, _ = cmd.Flags().GetBool("all")
		// App.DirMode, _ = cmd.Flags().GetBool("dir")
		// App.LinkMode, _ = cmd.Flags().GetBool("link")
		// App.FileMode, _ = cmd.Flags().GetBool("file")

		templates := &promptui.SelectTemplates{
			Label:    "📍 {{ . | magenta | italic | underline }}:",
			Active:   "⯈ {{ . | green | bold }}",
			Inactive: "  {{ . | cyan | bold}}",
			Details: `
--------------------
[i] select ".." to move to previous directory
[i] ctrl+c to exit
`,
		}

		path, _ = os.Getwd()
		prev = path
		startPath = path
		home, _ := os.UserHomeDir()

		// replace home dir in path with ~
		for {
			App.ClearDirs()
			App.GetPathContent(path)
			dirs := App.GetDirs()

			shortPath := strings.Replace(path, home, "~", -1)

			if startPath != path {
				dirs = append([]string{".."}, dirs...)
			}

			// if no flag passed then lsx will be in exporer mode
			prompt := promptui.Select{
				Label:     shortPath,
				Items:     dirs,
				Templates: templates,
				Size:      7,
			}
			_, selectedDir, err := prompt.Run()

			if err != nil {
				// handling ctrl+c
				if fmt.Sprint(err) == "^C" {
					fmt.Println("Exiting...")
					return
				} else {
					fmt.Printf("some error occured %v\n", err)
					return
				}
			}
			if selectedDir == ".." {
				moveTo = prev
			} else {
				moveTo = filepath.Join(path, selectedDir)
				prev = path
			}
			path = moveTo
		}

	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// func init() {
// 	rootCmd.Flags().BoolP("file", "f", false, "Only display files")
// 	rootCmd.Flags().BoolP("dir", "d", false, "Only display directories")
// 	rootCmd.Flags().BoolP("symlink", "s", false, "Only display symlink")
// 	rootCmd.Flags().BoolP("all", "a", false, "Display hidden (dotfiles) files as well")
// }
