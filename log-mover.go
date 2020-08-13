package main

import (
	"flag"
	"fmt"
	"github.com/agrison/go-commons-lang/stringUtils"
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

	if stringUtils.IsAnyBlank(*srcDir, *destDir, *appName, *suffix) {
		fmt.Println("must pass srcDir, destDir and appName flag")
		return
	}

	stringUtils.AppendIfMissing(*srcDir, "/", "/")
	stringUtils.AppendIfMissing(*destDir, "/", "/")

	filepath.Walk(*srcDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, *suffix) && strings.Contains(path, *appName) {
			os.Rename(path, *destDir+filepath.Base(path))
		}
		fmt.Println("move log complete")
		return nil
	})
}
