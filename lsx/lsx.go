package lsx

import (
	"encoding/json"
	"fmt"
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

	iconsFile string
	Icons     map[string]string
}

func (app *Lsx) Init() {
	app.raw = make([]string, 0)
	app.directory = make([]string, 0)
	app.file = make([]string, 0)
	app.misc = make([]string, 0)
	app.links = make(map[string]string)
	app.Icons = make(map[string]string)

	app.AllMode = false
	app.DirMode = false
	app.FileMode = false
	app.LinkMode = false

	app.iconsFile, _ = filepath.Abs("icons/files.json")
	//read json file
	jsonFile, err := os.Open(app.iconsFile)
	utils.CheckError(err)
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &app.Icons)
}

func (app *Lsx) DisplayIcons() {
	for key, icon := range app.Icons {
		fmt.Printf("%s 	%s\n", key, icon)
	}
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
		//TODO: check for executable

		app.raw = append(app.raw, fileName)
	}
}

func (app *Lsx) GetDirs() []string {
	return app.directory
}

func (app *Lsx) ClearDirs() {
	app.directory = nil
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
