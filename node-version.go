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
	fullPath, err := os.Executable()
	if err != nil {
		log.Fatalf("cannot get path name for the executable: %s", err)
	}
	dirPath := filepath.Dir(fullPath)
	dirs, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalf("cannot read current directory: %s", err)
	}
	// log.Printf("scanning %d objects", len(dirs))

	for _, dir := range dirs {
		// log.Println("obj", dir.Name())
		if !dir.IsDir() {
			continue
		}
		if dirName := dir.Name(); strings.HasPrefix(dirName, version) {
			// log.Printf("found installed version %s by prefix %s", dirName, version)
			versionPath := filepath.Join(dirPath, dirName)
			fmt.Println(versionPath)
			// log.Printf("setting NODE_VERSION=%s", versionPath)

			// TODO: if OS is windows
			// {
			// 	// cmd := fmt.Sprintf(`[Environment]::SetEnvironmentVariable("NODE_VERSION", "%s", [System.EnvironmentVariableTarget]::User)`, versionPath)
			// 	result, err := exec.Command("powershell", cmd).CombinedOutput()
			// 	if err != nil {
			// 		log.Fatalf("cannot set NODE_VERSION environment variable: %s", err)
			// 	}
			// 	log.Println("done", string(result))
			// }

			// {
			// 	result, err := exec.Command("node", "--version").CombinedOutput()
			// 	if err != nil {
			// 		log.Fatalf("cannot check node version: %s", err)
			// 	}
			// 	// log.Println("node --version", string(result))
			// }
			return
		}
	}
	log.Println("fullPath", fullPath)
	log.Fatalf("no directory for version %s found", version)
}
