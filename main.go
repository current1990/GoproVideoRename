package main

import (
	"os"
	"path/filepath"
	"strings"
)

func main() {
	workDir := "."
	files, e := os.ReadDir(workDir)
	if e != nil {
		panic(e)
	}

	for _, f := range files {
		if f.IsDir() {
			continue
		}
		filename := strings.Split(f.Name(), ".")[0]
		filetype := strings.ToLower(strings.Split(f.Name(), ".")[1])

		if strings.Contains(filename, "_") || filetype == "exe" {
			continue
		}

		vidPrefix := filename[:2]
		vidPart := filename[2:4]
		vidSeq := filename[4:]
		newName := vidPrefix + vidSeq + "_" + vidPart + ".mp4"
		//fmt.Printf("%v, %v, %v, %v\n", f.IsDir(), f.Type(), f.Name(), newName)
		e = os.Rename(filepath.Join(workDir, f.Name()), filepath.Join(workDir, newName))
		if e != nil {
			panic(e)
		}
	}
}
