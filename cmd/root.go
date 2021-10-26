package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
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

		
		platform := utils.GetOs()
		home, _ := os.UserHomeDir()
		var TMP_FILE string = filepath.Join(home, "lsx", ".lsx.tmp")
		
	
		var currentPath string

		currentPath, err := os.Getwd()
		if err != nil {
			panic(err)
		}

		// utils.ClearScreen(platform)
		for {
			// utils.ClearScreen(platform)

			// get all the directories from the current path
			App.ClearDirs()
			App.GetPathContent(currentPath)
			
			dirs := App.GetDirs()
			// remove all directories starting with .
			// if -a/--all is not used
			if !App.AllMode {
				dirs = utils.GetNonDotDirs(dirs)
			}

			if currentPath != filepath.VolumeName(currentPath) {
				dirs = append(dirs, "..")
			}


			searcher := func(input string, index int) bool {
				dir := dirs[index]
				name := strings.Replace(strings.ToLower(dir), " ", "", -1)
				input = strings.Replace(strings.ToLower(input), " ", "", -1)

				return strings.Contains(name, input)
			}

			prompt := promptui.Select{
				Label:        strings.Replace(currentPath, home, "~", -1),
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
					utils.WriteToFile(TMP_FILE, currentPath)
					utils.ClearScreen(platform)
					os.Exit(0)
				}
				fmt.Printf("some error occured %v\n", err)
				os.Exit(1)
			}

			fmt.Println("aaa", currentPath)
			parts := strings.Split(currentPath, string(os.PathSeparator))
			if selectedDir == ".." {
				parts = parts[:len(parts)-1]
			} else {
				parts = append(parts, selectedDir)
			}
			path := strings.Join(parts, string(os.PathSeparator))
			if path == filepath.VolumeName(currentPath) && runtime.GOOS == "windows" {
				currentPath = path + string(os.PathSeparator)
			} else {
				currentPath = path
			}
			currentPath = filepath.Clean(currentPath)
		
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
