package cmd

import (
	"os"
	"path/filepath"
	"regexp"

	"github.com/souvikinator/lsx/utils"
	"github.com/spf13/cobra"
)

var pathName, dirPath string
var alias_regexp = `^[a-zA-Z0-9-_]+$`

// setAliasCmd represents the setAlias command
var setAliasCmd = &cobra.Command{
	Use:   "set-alias",
	Short: "Use this command to set or update an alias",
	Long: `- set-alias can be used to create a new alias or update an existing alias
- path and alias should be provided using flags --path-name/-n and --path/-p respectively
- make sure the path provided is a directory which exists and the alias can only contain alphanumeric characters along with underscore and hyphens`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(pathName) == 0 || len(dirPath) == 0 {
			utils.Warn(" alias and directory path not provided! use --path-name and --path respectively")
			utils.Warn(" use  'lsx set-alias --help' to know more")
			os.Exit(1)
		}

		// path-name should contain only alphanumeric and underscore, dash
		isValidAlias, err := regexp.MatchString(alias_regexp, pathName)
		utils.CheckError(err)
		if !isValidAlias {
			utils.Warn("'" + pathName + "' is not a valid alias! ")
			utils.Warn("alias can only contain alphanumeric, underscore and dash characters")
			os.Exit(1)
		}

		dirPath, err = filepath.Abs(dirPath)
		utils.CheckError(err)

		// path exists
		if !utils.PathExists(dirPath) {
			utils.Err("'" + dirPath + "' not found, make sure that the path exists and is a directory")
		}

		// check if dirPath exists in the records
		for a, p := range App.Alias {
			if p == dirPath {
				utils.Err("'" + dirPath + "' is already associated with alias '" + a + "'")
			}
		}

		// path has to be a directory
		if !utils.PathIsDir(dirPath) {
			utils.Err("'" + dirPath + "' in not a directory, make sure that the path exists and is a directory")
			os.Exit(1)
		}

		// all good. update alias file
		App.Alias[pathName] = dirPath
		utils.WriteYamlFile(App.AliasFile, App.Alias)
		utils.Info("done")
	},
}

func init() {
	setAliasCmd.Flags().StringVarP(&pathName, "path-name", "n", "", "alias for a path")
	setAliasCmd.Flags().StringVarP(&dirPath, "path", "p", "", "takes in path to directory")
}
