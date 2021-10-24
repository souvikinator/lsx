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
var LSX_VERSION string = "0.1.0"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "lsx",
	Short: " A command line utility, let's you navigate across the terminal like butter",
	Run: func(cmd *cobra.Command, args []string) {

		App.Init()
		App.AllMode, _ = cmd.Flags().GetBool("all")

		showVersion, _ := cmd.Flags().GetBool("version")
		if showVersion {
			fmt.Println("lsx,", LSX_VERSION)
			os.Exit(0)
		}

		templates := &promptui.SelectTemplates{
			Label:    "📍 {{ . | magenta | italic | underline }}:",
			Active:   "⯈ {{ . | green | bold | italic }}",
			Inactive: "  {{ . | cyan | bold }}",
			Details: `
_________________________
* ctrl+c to exit
* select ".." to move to previous directory
`,
		}

		var startPath string
		var pathStack utils.Stack
		pathStack.Init()

		platform := utils.GetOs()
		startPath, _ = os.Getwd()
		home, _ := os.UserHomeDir()
		pathStack.Push(startPath)

		var lsx_config_path string = filepath.Join(home, ".config", "lsx")
		var lsx_tmp_file string = filepath.Join(lsx_config_path, "lsx.tmp")

		err := os.MkdirAll(lsx_config_path, 0664)
		utils.CheckError(err)

		utils.ClearScreen(platform)
		for {
			// utils.ClearScreen(platform)
			App.ClearDirs()

			currentPath := pathStack.Top()

			// get all the directories from the current path
			App.GetPathContent(currentPath)
			dirs := App.GetDirs()

			// remove all directories starting with .
			// if -a/--all is not used
			if !App.AllMode {
				dirs = utils.GetNonDotDirs(dirs)
			}

			// replace home dir in path with ~
			shortCurrentPath := strings.Replace(currentPath, home, "~", -1)

			// if current path is startPath
			// then user can't go back as they started
			// from startPath (there's no going back!)
			if startPath != currentPath {
				dirs = append([]string{".."}, dirs...)
			}

			searcher := func(input string, index int) bool {
				dir := dirs[index]
				name := strings.Replace(strings.ToLower(dir), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)

				return strings.Contains(name, input)
			}

			prompt := promptui.Select{
				Label:        shortCurrentPath,
				Items:        dirs,
				Templates:    templates,
				Size:         11,
				Searcher:     searcher,
				HideSelected: true,
			}
			_, selectedDir, err := prompt.Run()

			// handle ctrl+c and error
			if err != nil {
				if utils.IsKeyboardInterrupt(err) {
					//write currentPath to ~/.config/lsx.yml
					utils.WriteToFile(lsx_tmp_file, currentPath)
					utils.ClearScreen(platform)
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

func init() {
	rootCmd.Flags().BoolP("all", "a", false, "Display hidden (dotdirs) durectories as well")
	rootCmd.Flags().BoolP("version", "v", false, "Display lsx version")
}
