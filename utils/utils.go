package utils

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

// stack implementation to keep track of paths
type Stack struct {
	data []string
}

func (s *Stack) Init() {
	s.data = make([]string, 0)
}

func (s *Stack) Push(d string) {
	s.data = append(s.data, d)
}

func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		log.Panicln("underflow!")
		os.Exit(1)
	}
	toBePoped := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return toBePoped
}

func (s *Stack) Top() string {
	if len(s.data) == 0 {
		log.Panic("underflow!")
		os.Exit(1)
	}
	return s.data[len(s.data)-1]
}

/*******misc utility functions******/

func CheckError(err error) {
	if err != nil {
		fmt.Printf("some error occured %v\n", err)
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
		fmt.Println("\033[H\033[2J")
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
	// fmt.Println("abs: ", abs)
	CheckError(err)
	src, _ := filepath.EvalSymlinks(path)
	// fmt.Println("src: ", src)
	return src != abs
}

func IsDotFile(filename string) bool {
	return strings.HasPrefix(filename, ".")
}

func ResolveLink(path string) string {
	absPath, err := filepath.Abs(path)
	CheckError(err)
	// fmt.Println(">>")
	src, _ := filepath.EvalSymlinks(absPath)
	// fmt.Println(">>")
	return src
}
