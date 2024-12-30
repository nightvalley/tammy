package main

import (
	"CountLines/internal/files"
	"CountLines/internal/forms"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

var StartTime time.Time

func main() {
	availableForms := []string{"table", "list", "total"}
	formFlag := flag.String("f", availableForms[0], "Available forms: "+strings.Join(availableForms, ", "))
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

	f := files.Files{}
	switch *formFlag {
	case availableForms[0]:
		forms.TableOutput(expandedPath)
	case availableForms[1]:
		forms.ListOutput(expandedPath)
	case availableForms[2]:
		f.FoundAllFilesInDir(path)
		fmt.Println(f.TotalLines)
	}
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
