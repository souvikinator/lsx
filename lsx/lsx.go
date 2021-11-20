package lsx

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
	activeDir string
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
	app.directory = make([]string, 0)
	app.activeDir = ""

	app.AllMode = false
	app.Version = "v0.1.5"
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
	app.ClearDirs()
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
	app.activeDir = path
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
// FIXME: not sorting all the dirs
func (app *Lsx) GetDirs() []string {
	rankedDirList := make([]string, 0)

	// not the actual cwd,
	// cwd for the process
	cwd := app.activeDir

	if len(app.AccessRecordFile) < 1 || len(app.directory) < 1 {
		return app.directory
	}
	//clear frequency record
	app.FrecencyRecords = make(entries, 0)
	// recalculate
	app.CalculateFrecency()
	// fmt.Println("fr:", app.FrecencyRecords)
	absPathDirs := utils.GetAbsPathSlice(cwd, app.directory)
	// fmt.Println("absPathDirs: ", absPathDirs)
	for _, ob := range app.FrecencyRecords {
		i := sort.SearchStrings(absPathDirs, ob.key)
		// fmt.Println(">>", i, ob.key)
		if i < len(absPathDirs) && ob.key == absPathDirs[i] {
			//remove absPath from directory
			rankedDirList = append(rankedDirList, strings.ReplaceAll(absPathDirs[i], cwd+"/", ""))
			//remove from absPathDirs
			absPathDirs = utils.Remove(absPathDirs, i)
			// fmt.Println(">>", absPathDirs)
			// fmt.Println("<<", rankedDirList)
		}
	}
	// fmt.Println("absPathDirs(after): ", absPathDirs)
	//merge remaining absPathDirs to rankedDirList
	for _, dir := range absPathDirs {
		rankedDirList = append(rankedDirList, strings.ReplaceAll(dir, cwd+"/", ""))
	}
	// fmt.Println("rankedDirList: ", rankedDirList)
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
	err = yaml.Unmarshal([]byte(f), &app.Alias)
	utils.CheckError(err)
}

func (app *Lsx) ReadAccessRecordFile() {
	f, err := ioutil.ReadFile(app.AccessRecordFile)
	utils.CheckError(err)
	err = yaml.Unmarshal([]byte(f), &app.AccessRecords)
	utils.CheckError(err)
}

func (app *Lsx) WriteAccessRecordFile() {
	d, err := yaml.Marshal(app.AccessRecords)
	utils.CheckError(err)
	err = ioutil.WriteFile(app.AccessRecordFile, d, 0644)
	utils.CheckError(err)
}
