package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/souvikinator/lsx/lsx"
	"github.com/souvikinator/lsx/utils"
	"github.com/spf13/cobra"
)

var App lsx.Lsx

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "lsx",
	Version: App.Version,
	Args:    cobra.ArbitraryArgs,
	Short:   " A command line utility, let's you navigate across the terminal like butter",
	Long:    `lsx stands for "ls Xtended". It let's you navigate across the terminal using cursors along with search. One can also set aliases for paths.`,
	Run: func(cmd *cobra.Command, args []string) {

		home := utils.HomeDir()
		platform := utils.GetOs()

		App.AllMode, _ = cmd.Flags().GetBool("all")

		// if alias is passed
		if len(args) > 0 {
			pathAlias := args[0]
			ChdirToAlias(pathAlias)
			utils.ClearScreen(platform)
			os.Exit(0)
		}
		// if no args then prompt the user
		// #61D1C2  #7E6CFA how to use these bruh!?
		templates := &promptui.SelectTemplates{
			Label:    "📌 {{ . | magenta | italic | underline }}:",
			Active:   "> {{ . | yellow | bold }}",
			Inactive: "  {{ . | cyan }}",
			Help:     `{{ " ctrl+c to exit and ↑ ↓ → ← or h,j,k,l to navigate" | faint }}`,
		}

		var currentPath string

		currentPath, err := os.Getwd()
		utils.CheckError(err)

		//TODO: utils.ClearScreen(platform)
		for {
			// get all the directories from the current path
			App.GetPathContent(currentPath)

			dirs := App.GetDirs()
			// remove all directories starting with .
			// if -a/--all is not used
			if !App.AllMode {
				dirs = utils.GetNonDotDirs(dirs)
			}
			dirs = append([]string{".."}, dirs...)

			searcher := func(input string, index int) bool {
				dir := dirs[index]
				name := strings.Replace(strings.ToLower(dir), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)

				return strings.Contains(name, input)
			}

			prompt := promptui.Select{
				Label:        fmt.Sprintf("%s (%d)", strings.Replace(currentPath, home, "~", -1), len(dirs)-1),
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
					//write currentPath to temp file
					//for use after the process ends
					utils.WriteToFile(App.TempFile, currentPath)
					// write to access record file
					App.WriteAccessRecordFile()
					utils.ClearScreen(platform)
					os.Exit(0)
				}
				utils.CheckError(err)
			}

			currentPath = filepath.Join(currentPath, selectedDir)
			// TODO: record hit count and last access time in access record for selectedDir
			stats, exists := App.AccessRecords[currentPath]
			if !exists {
				App.AccessRecords[currentPath] = []int64{1, time.Now().Unix()}
			} else {
				App.AccessRecords[currentPath] = []int64{stats[0] + 1, time.Now().Unix()}
			}
		}

	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	App.Init()
	// commands
	rootCmd.AddCommand(aliasCmd)
	rootCmd.AddCommand(removeAliasCmd)
	rootCmd.AddCommand(setAliasCmd)

	//flags
	rootCmd.Flags().BoolP("all", "a", false, "display hidden (dotdirs) directories as well")
}

func ChdirToAlias(pathAlias string) {
	if len(App.Alias) == 0 {
		utils.Warn("no alias found in records")
		utils.Warn("get started by using 'set-alias' command")
		os.Exit(0)
	}

	moveToPath, ok := App.Alias[pathAlias]
	if !ok {
		utils.Err("alias '" + pathAlias + "' not found")
	}

	// path corresponding to alias exists?
	if !utils.PathExists(moveToPath) {
		utils.Err("'" + dirPath + "' not found, make sure that the path exists and is a directory")
	}

	utils.WriteToFile(App.TempFile, moveToPath)

}
