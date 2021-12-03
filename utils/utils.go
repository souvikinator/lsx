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
	"gopkg.in/yaml.v3"
)

func ReadYamlFile(filepath string, store *interface{}) {
	f, err := ioutil.ReadFile(filepath)
	CheckError(err)
	err = yaml.Unmarshal([]byte(f), &store)
	CheckError(err)
}

func WriteYamlFile(filepath string, data interface{}) {
	d, err := yaml.Marshal(data)
	CheckError(err)
	err = ioutil.WriteFile(filepath, d, 0644)
	CheckError(err)

}

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

func Remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}

func GetAbsPathSlice(root string, paths []string) []string {
	tmp := make([]string, 0)
	for _, dir := range paths {
		absPath := filepath.Join(root, dir)
		tmp = append(tmp, absPath)
	}
	return tmp
}

func CreateFile(filepath string) {
	f, err := os.OpenFile(filepath, os.O_RDONLY|os.O_CREATE, 0600)
	CheckError(err)
	defer f.Close()
}

func WriteToFile(filepath, data string) {
	err := ioutil.WriteFile(filepath, []byte(data), 0644)
	CheckError(err)
}

func CheckError(err error, msg ...interface{}) {
	if err != nil {
		fmt.Println(msg...)
		fmt.Printf("lsx error: some error occured\n %v\n", err)
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
