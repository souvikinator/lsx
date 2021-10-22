package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/manifoldco/promptui"
	"github.com/souvikinator/lsx/lsx"
	"github.com/souvikinator/lsx/utils"
	"github.com/spf13/cobra"
)

var App lsx.Lsx

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lsx",
	Short: " A command line utility written in golang which beautifies and extends ls command",
	Run: func(cmd *cobra.Command, args []string) {

		templates := &promptui.SelectTemplates{
			Label:    "📍 {{ . | magenta | italic | underline }}:",
			Active:   "⯈ {{ . | green | bold }}",
			Inactive: "  {{ . | cyan | bold}}",
			Details: `
_________________________
[i] select ".." to move to previous directory
[i] ctrl+c to exit
`,
		}

		var startPath string
		var pathStack utils.Stack
		pathStack.Init()

		platform := utils.GetOs()
		startPath, _ = os.Getwd()
		home, _ := os.UserHomeDir()
		pathStack.Push(startPath)

		for {
			utils.ClearScreen(platform)
			App.ClearDirs()

			currentPath := pathStack.Top()

			// get all the directories from the current path
			App.GetPathContent(currentPath)
			dirs := App.GetDirs()

			// replace home dir in path with ~
			shortCurrentPath := strings.Replace(currentPath, home, "~", -1)

			// if current path is startPath
			// then user can't go back as they started
			// from startPath (there's no going back!)
			if startPath != currentPath {
				dirs = append([]string{".."}, dirs...)
			}

			prompt := promptui.Select{
				Label:     shortCurrentPath,
				Items:     dirs,
				Templates: templates,
				Size:      7,
			}
			_, selectedDir, err := prompt.Run()

			// handle ctrl+c and error
			if err != nil {
				if utils.IsKeyboardInterrupt(err) {
					// TODO: perform syscall to change dir
					os.Exit(0)
				}
				fmt.Printf("some error occured %v\n", err)
				os.Exit(1)
			}

			if selectedDir == ".." {
				pathStack.Pop()
			} else {
				moveTo := filepath.Join(currentPath, selectedDir)
				pathStack.Push(moveTo)
			}
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
