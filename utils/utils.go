package utils

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
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
