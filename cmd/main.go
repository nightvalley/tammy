package main

import (
	files "CountLines/internal"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	pathFlag := flag.String("p", ".", "path")
	flag.Parse()

	path, err := expandPath(*pathFlag)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	files := files.Files{}
	files.FoundAllFilesInDir(path)
}

func expandPath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}

		return filepath.Join(usr.HomeDir, path[1:]), nil
	} else {
		var err error
		path, err = os.Getwd()
		if err != nil {
			return "", err
		}

		return path, nil
	}
}
