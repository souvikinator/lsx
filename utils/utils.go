package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gookit/color"
	"golang.org/x/term"
)

type Filepath struct {
	parts []string
	root  string
}

func (fp *Filepath) String() string {
	return filepath.Clean(fp.root + strings.Join(fp.parts, string(os.PathSeparator)))
}

func (fp *Filepath) To(path string) {
	if filepath.IsAbs(path) {
		fp.root = filepath.VolumeName(path) + string(os.PathSeparator)
		fp.parts = strings.Split(path[len(fp.root):], string(os.PathSeparator))
	} else if path == ".." {
		if len(fp.parts) == 0 {
			return
		}
		fp.parts = fp.parts[:len(fp.parts)-1]
	} else {
		fp.parts = append(fp.parts, path)
	}

}

/*******misc utility functions******/

func HomeDir() string {
	home, err := os.UserHomeDir()
	CheckError(err)
	return home
}

func GetNonDotDirs(dirs []string) []string {
	nonDotDirs := make([]string, 0)
	for _, dir := range dirs {
		if !strings.HasPrefix(dir, ".") {
			nonDotDirs = append(nonDotDirs, dir)
		}
	}
	return nonDotDirs
}

func GetTerminalHeight() int {
	_, height, err := term.GetSize(0)
	CheckError(err)
	return height
}

// func ReadAccessRecordFile(filepath string, data *map[string][]int64) {
// 	f, err := ioutil.ReadFile(filepath)
// 	CheckError(err)
// 	err = yaml.Unmarshal([]byte(f), data)
// 	CheckError(err)
// }

// func WriteAccessRecordFile(filepath string, data map[string][]int64) {
// 	d, err := yaml.Marshal(data)
// 	CheckError(err)
// 	err = ioutil.WriteFile(filepath, d, 0644)
// 	CheckError(err)
// }

// func ReadAliasFile(filepath string, data *map[string]string) {
// 	f, err := ioutil.ReadFile(filepath)
// 	CheckError(err)
// 	err = yaml.Unmarshal([]byte(f), data)
// 	CheckError(err)
// }

// func WriteAliasFile(filepath string, data map[string]string) {
// 	d, err := yaml.Marshal(data)
// 	CheckError(err)
// 	err = ioutil.WriteFile(filepath, d, 0644)
// 	CheckError(err)
// }

func CreateFile(filepath string) {
	f, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0666)
	CheckError(err)
	defer f.Close()
}

func WriteToFile(filepath, data string) {
	err := ioutil.WriteFile(filepath, []byte(data), 0644)
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		fmt.Printf("lsx error: some error occured %v\n", err)
		os.Exit(1)
	}
}

func IsKeyboardInterrupt(err error) bool {
	return fmt.Sprint(err) == "^C"
}

func GetOs() string {
	return runtime.GOOS
}

func Chdir(path string) {
	cmd := exec.Command("cd", path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

func ClearScreen(Os string) {
	if Os == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()
	}
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	CheckError(err)
	return true
}

func PathIsFile(path string) bool {
	fileInfo, err := os.Stat(path)
	CheckError(err)
	return !fileInfo.IsDir()
}

func PathIsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	CheckError(err)
	return fileInfo.IsDir()
}

func PathIsLink(path string) bool {
	abs, err := filepath.Abs(path)
	CheckError(err)
	src, _ := filepath.EvalSymlinks(path)
	return src != abs
}

func IsDotFile(filename string) bool {
	return strings.HasPrefix(filename, ".")
}

func ResolveLink(path string) string {
	absPath, err := filepath.Abs(path)
	CheckError(err)
	src, _ := filepath.EvalSymlinks(absPath)
	return src
}

func Err(msg string) {
	color.Error.Prompt(msg)
	os.Exit(1)
}

func Warn(msg string) {
	color.Warn.Prompt(msg)
}

func Info(msg string) {
	color.Info.Prompt(msg)
}
