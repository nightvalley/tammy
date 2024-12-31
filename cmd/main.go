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

func main() {
	availableForms := []string{"table", "list", "total"}

	var (
		formFlag   = flag.String("f", availableForms[0], "Available forms: "+strings.Join(availableForms, ", "))
		pathFlag   = flag.String("p", ".", "path")
		timeFlag   = flag.Bool("t", false, "time")
		showHidden = flag.Bool("h", false, "show hidden files")
		filetype   = flag.String("ft", "", "filetype")
	)
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

	if !*timeFlag {
		switch *formFlag {
		case availableForms[0]:
			forms.TableOutput(expandedPath, *showHidden, *filetype)
		case availableForms[1]:
			forms.ListOutput(expandedPath, *showHidden, *filetype)
		case availableForms[2]:
			f.FoundAllFilesInDir(expandedPath, *showHidden, *filetype)
			fmt.Println(f.TotalLines)
		}
	} else {
		switch *formFlag {
		case availableForms[0]:
			t := time.Now()
			forms.TableOutput(expandedPath, *showHidden, *filetype)
			duration := time.Since(t)
			fmt.Printf("\nExecution time: %v\n", duration)
		case availableForms[1]:
			t := time.Now()
			forms.ListOutput(expandedPath, *showHidden, *filetype)
			duration := time.Since(t)
			fmt.Printf("\nExecution time: %v\n", duration)
		case availableForms[2]:
			t := time.Now()
			f.FoundAllFilesInDir(path, *showHidden, *filetype)
			fmt.Println(f.TotalLines)
			duration := time.Since(t)
			fmt.Printf("\nExecution time: %v\n", duration)
		}
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
