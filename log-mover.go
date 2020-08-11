package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	srcDir := flag.String("srcDir", "", "source directory contains logs")
	destDir := flag.String("destDir", "", "target directory for move logs")
	appName := flag.String("appName", "", "application name")
	suffix := flag.String("suffix", ".log.gz", "file extension")

	flag.Parse()

	flagsExist := isFlagPassed([]string{"srcDir", "destDir", "appName"})

	if !flagsExist {
		return
	}

	filepath.Walk(*srcDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, *suffix) && strings.Contains(path, *appName) {
			os.Rename(path, *destDir+filepath.Base(path))
		}
		return nil
	})
}

func isFlagPassed(names []string) bool {
	flagLen := len(names)
	count := 0
	for _, name := range names {
		flag.Visit(func(f *flag.Flag) {
			if f.Name == name {
				count++
			}
		})
	}
	return flagLen == count
}
