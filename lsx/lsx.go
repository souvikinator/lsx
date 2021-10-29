package lsx

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/souvikinator/lsx/utils"
)

type Lsx struct {
	raw       []string
	directory []string
	file      []string
	links     map[string]string
	misc      []string
	// executables []string

	// modes
	DirMode  bool
	FileMode bool
	LinkMode bool
	AllMode  bool

	Version   string
	ConfigDir string
	AliasFile string
	TempFile  string

	Alias map[string]string
}

func (app *Lsx) Init() {
	//data
	app.raw = make([]string, 0)
	app.directory = make([]string, 0)
	app.file = make([]string, 0)
	app.misc = make([]string, 0)
	app.links = make(map[string]string)
	app.Alias = make(map[string]string)
	app.AllMode = false
	app.DirMode = false
	app.FileMode = false
	app.LinkMode = false

	app.Version = "v0.1.3"
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
	return (!app.AllMode && !app.DirMode && !app.FileMode && !app.LinkMode)
}

func (app *Lsx) GetPathContent(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fileName := f.Name()
		absPath, _ := filepath.Abs(filepath.Join(path, fileName))

		if utils.PathIsLink(absPath) {
			src := utils.ResolveLink(absPath)
			app.links[fileName] = src

		} else if utils.PathIsDir(absPath) {
			app.directory = append(app.directory, fileName)

		} else if !utils.PathIsDir(absPath) {
			app.file = append(app.file, fileName)

		} else {
			app.misc = append(app.misc, fileName)

		}

		app.raw = append(app.raw, fileName)
	}
}

func (app *Lsx) GetDirs() []string {
	return app.directory
}

func (app *Lsx) ClearDirs() {
	app.directory = make([]string, 0)
}

func (app *Lsx) GetFiles() []string {
	return app.file
}

func (app *Lsx) ClearFiles() {
	app.file = nil
}

func (app *Lsx) GetLinks() map[string]string {
	return app.links
}

func (app *Lsx) ClearLinks() {
	app.links = nil
}

func (app *Lsx) GetMisc() []string {
	return app.misc
}

func (app *Lsx) ClearMisc() {
	app.misc = nil
}

func (app *Lsx) GetRaw() []string {
	return app.raw
}

func (app *Lsx) ClearRaw() {
	app.raw = nil
}
