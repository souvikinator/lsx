package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	currDir, _ := os.Getwd()
	moveTo := filepath.Join(currDir, "scripts")
	moveToAbs, _ := filepath.Abs(moveTo)
	fmt.Println("currDir: ", currDir)
	fmt.Println("moveto : ", moveTo)
	fmt.Println("moveto abs: ", moveToAbs)
}
