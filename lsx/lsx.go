package lsx

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/souvikinator/lsx/utils"
)

type Lsx struct {
	directory []string

	// modes
	AllMode bool

	Version   string
	ConfigDir string
	AliasFile string
	TempFile  string

	Alias map[string]string
}

func (app *Lsx) Init() {
	//data
	app.directory = make([]string, 0)
	app.Alias = make(map[string]string)
	app.AllMode = false

	app.Version = "v0.1.4"
	home := utils.HomeDir()
	app.ConfigDir = filepath.Join(home, ".config", "lsx")
	app.AliasFile = filepath.Join(app.ConfigDir, "alias.yaml")
	app.TempFile = filepath.Join(app.ConfigDir, "lsx.tmp")

	// create configDir if doesn't exist
	err := os.MkdirAll(app.ConfigDir, 0664)
	utils.CheckError(err)

	// create alias file and temp file
	utils.CreateFile(app.AliasFile)
	utils.CreateFile(app.TempFile)

	// read data from alias file
	utils.ReadYamlFile(app.AliasFile, &app.Alias)

}

func (app *Lsx) NoFlagPassed() bool {
	return (!app.AllMode)
}

func (app *Lsx) GetPathContent(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileName := f.Name()
		absPath, _ := filepath.Abs(filepath.Join(path, fileName))

		if utils.PathIsDir(absPath) {
			app.directory = append(app.directory, fileName)
		}
	}
}

func (app *Lsx) GetDirs() []string {
	return app.directory
}

func (app *Lsx) ClearDirs() {
	app.directory = make([]string, 0)
}
