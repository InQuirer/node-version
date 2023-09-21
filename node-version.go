package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	flag.Parse()
	version := flag.Arg(0)
	if version == "" {
		log.Fatalf("version argument missing")
	}
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}
	executablePath, err := os.Executable()
	if err != nil {
		log.Fatalf("cannot get path name for the executable: %s", err)
	}
	dirPath := filepath.Dir(executablePath)
	dirs, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("cannot read current directory: %s", err)
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}
		if dirName := dir.Name(); strings.HasPrefix(dirName, version) {
			versionPath := filepath.Join(dirPath, dirName)
			fmt.Println(versionPath)
			return
		}
	}
	log.Fatalf("no directory for version %s found in %s", version, dirPath)
}
