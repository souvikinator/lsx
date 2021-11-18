package lsx

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"time"

	"github.com/souvikinator/lsx/utils"
	"gopkg.in/yaml.v3"
)

type entry struct {
	val int
	key string
}

type entries []entry
type Lsx struct {
	directory []string

	// modes
	AllMode bool

	Version          string
	ConfigDir        string
	AliasFile        string
	TempFile         string
	AccessRecordFile string

	Alias         map[string]string
	AccessRecords map[string][]int64

	FrecencyRecords entries
}

func (app *Lsx) Init() {
	//data
	app.AllMode = false
	app.Version = "v0.1.5"
	app.directory = make([]string, 0)
	app.Alias = make(map[string]string)
	app.AccessRecords = make(map[string][]int64)
	app.FrecencyRecords = make(entries, 0)

	home := utils.HomeDir()
	app.ConfigDir = filepath.Join(home, ".config", "lsx")
	app.AliasFile = filepath.Join(app.ConfigDir, "alias.yaml")
	app.TempFile = filepath.Join(app.ConfigDir, "lsx.tmp")
	app.AccessRecordFile = filepath.Join(app.ConfigDir, "access-record.yaml")

	// create configDir if doesn't exist
	err := os.MkdirAll(app.ConfigDir, 0664)
	utils.CheckError(err)

	// create alias file , temp file and access record file
	utils.CreateFile(app.AliasFile)
	utils.CreateFile(app.TempFile)
	utils.CreateFile(app.AccessRecordFile)

	app.ReadAliasFile()
	app.ReadAccessRecordFile()

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

		if !utils.PathIsLink(absPath) && utils.PathIsDir(absPath) {
			app.directory = append(app.directory, fileName)
		}
	}
}

func (app *Lsx) CalculateFrecency() {
	for dir, stats := range app.AccessRecords {
		age := stats[1] - time.Now().Unix()
		score := int(utils.FrecencyScore(stats[0], age))
		// TODO: remove entry from accress records if score
		// less than some value
		// TODO: remove dir if does not exists
		app.FrecencyRecords = append(app.FrecencyRecords, entry{val: score, key: dir})
	}
	app.SortFrecencyRecord()
}

func (app *Lsx) SortFrecencyRecord() {
	sort.Slice(app.FrecencyRecords, func(i, j int) bool {
		return app.FrecencyRecords[i].val > app.FrecencyRecords[j].val
	})
}

// sorts directory as per frecency scores
func (app *Lsx) GetDirs() []string {
	rankedDirList := make([]string, 0)
	if len(app.FrecencyRecords) < 1 {
		return app.directory
	}
	//calculate frecency scores
	app.CalculateFrecency()
	//TODO: how to sort dirs as per ranks
	return rankedDirList
}

func (app *Lsx) ClearDirs() {
	app.directory = make([]string, 0)
}

func (app *Lsx) WriteAliasFile() {
	d, err := yaml.Marshal(app.Alias)
	utils.CheckError(err)
	err = ioutil.WriteFile(app.AliasFile, d, 0644)
	utils.CheckError(err)
}

func (app *Lsx) ReadAliasFile() {
	f, err := ioutil.ReadFile(app.AliasFile)
	utils.CheckError(err)
	err = yaml.Unmarshal([]byte(f), app.Alias)
	utils.CheckError(err)
}

func (app *Lsx) ReadAccessRecordFile() {
	f, err := ioutil.ReadFile(app.AccessRecordFile)
	utils.CheckError(err)
	err = yaml.Unmarshal([]byte(f), app.AccessRecords)
	utils.CheckError(err)
}

func (app *Lsx) WriteAccessRecordFile() {
	d, err := yaml.Marshal(app.AccessRecords)
	utils.CheckError(err)
	err = ioutil.WriteFile(app.AccessRecordFile, d, 0644)
	utils.CheckError(err)
}
