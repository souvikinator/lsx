package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gookit/color"
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

		var path string

		// flags
		App.Init()
		App.DisplayIcons()
		App.AllMode, _ = cmd.Flags().GetBool("all")
		App.DirMode, _ = cmd.Flags().GetBool("dir")
		App.LinkMode, _ = cmd.Flags().GetBool("link")
		App.FileMode, _ = cmd.Flags().GetBool("file")

		// get current working directory if not path provided
		if len(args) == 0 {
			path, _ = os.Getwd()
		} else {
			path, _ = filepath.Abs(args[0])
			// fmt.Println(path)
		}

		if !utils.PathExists(path) {
			color.Danger.Printf("'%s' path not found!", path)
			os.Exit(0)
		}

		App.GetPathContent(path)

		// print dirs
		for _, p := range App.GetDirs() {
			// if -a not used and p is dotfiles then skip
			if !App.AllMode && utils.IsDotFile(p) {
				continue
			}
			color.Blue.Printf("%s  ", color.Bold.Render(p))
		}
		// print files
		for _, p := range App.GetFiles() {
			// if -a not used and p is dotfiles then skip
			if !App.AllMode && utils.IsDotFile(p) {
				continue
			}
			color.Printf("%s  ", p)
		}
		// print symlink
		for p, l := range App.GetLinks() {
			// if -a not used and p is dotfiles then skip
			if !App.AllMode && utils.IsDotFile(p) {
				continue
			}
			if len(l) == 0 {
				color.Printf("%s  ", color.Red.Render(p))
			} else {
				color.Printf("%s  ", color.Green.Render(p))
			}
		}

		for _, p := range App.GetMisc() {
			color.Gray.Printf("%s  ", p)
		}
		fmt.Printf("\n")
	},
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	rootCmd.Flags().BoolP("file", "f", false, "Only display files")
	rootCmd.Flags().BoolP("dir", "d", false, "Only display directories")
	rootCmd.Flags().BoolP("symlink", "s", false, "Only display symlink")
	rootCmd.Flags().BoolP("all", "a", false, "Display hidden (dotfiles) files as well")
}
