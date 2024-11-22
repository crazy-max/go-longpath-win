package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	tmpDir := os.TempDir()
	wd, _ := os.Getwd()
	fmt.Printf("TempDir: %s\n", tmpDir)
	fmt.Printf("WorkingDir: %s\n", wd)

	tmpDirSym, _ := filepath.EvalSymlinks(tmpDir)
	wdSym, _ := filepath.EvalSymlinks(wd)
	fmt.Printf("TempDirEvalSymlinks: %s\n", tmpDirSym)
	fmt.Printf("WorkingDirEvalSymlinks: %s\n", wdSym)
}
