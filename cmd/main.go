package main

import (
	files "CountLines/internal"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

var StartTime time.Time

func main() {
	pathFlag := flag.String("p", ".", "path")
	flag.Parse()

	var path string
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	} else {
		path = *pathFlag
	}

	expandedPath, err := ExpandPath(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// fmt.Println(expandedPath)

	files := files.Files{}
	files.FoundAllFilesInDir(expandedPath)
}

func ExpandPath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		usr, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("failed to get current user: %v", err)
		}
		return filepath.Join(usr.HomeDir, path[1:]), nil
	} else if path == "." {
		return os.Getwd()
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %v", err)
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("path does not exist")
	}

	return absPath, nil
}
