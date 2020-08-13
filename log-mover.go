package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	srcDir := flag.String("srcDir", "", "source directory contains logs")
	destDir := flag.String("destDir", "", "destination directory for storing logs")
	appName := flag.String("appName", "", "application name")
	suffix := flag.String("suffix", ".log.gz", "log file extension")

	flag.Parse()

	if !isFlagPassed([]string{"srcDir", "destDir", "appName"}) {
		fmt.Println("must pass srcDir, destDir and appName flag")
		return
	}

	if !strings.HasSuffix(*srcDir, "/") {
		*srcDir = *srcDir + "/"
	}

	if !strings.HasSuffix(*destDir, "/") {
		*destDir = *destDir + "/"
	}

	filepath.Walk(*srcDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, *suffix) && strings.Contains(path, *appName) {
			os.Rename(path, *destDir+filepath.Base(path))
		}
		return nil
	})
}

func isFlagPassed(names []string) bool {
	flagLen, count := len(names), 0
	for _, name := range names {
		flag.Visit(func(f *flag.Flag) {
			if f.Name == name {
				count++
			}
		})
	}
	return flagLen == count
}
