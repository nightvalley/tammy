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

	"github.com/charmbracelet/log"
)

func main() {
	availableForms := []string{"table", "list", "total", "tree"}

	var (
		formFlag       = flag.String("f", availableForms[0], "Available forms: "+strings.Join(availableForms, ", "))
		pathFlag       = flag.String("p", ".", "path")
		filetypeFlag   = flag.String("ft", "", "count files with file type")
		timeFlag       = flag.Bool("t", false, "benchmark")
		showHiddenFlag = flag.Bool("h", false, "show hidden files")
		fileSizeFlag   = flag.Bool("s", false, "count size of files")
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
		log.Error(err)
		return
	}

	var fileType string
	if filetypeFlag != nil && *filetypeFlag != "" && (*filetypeFlag)[0] != '.' {
		fileType = fmt.Sprintf(".%s", *filetypeFlag)
	} else if filetypeFlag != nil {
		fileType = *filetypeFlag
	}

	f := files.Files{}
	flags := files.Flags{
		Hidden:   *showHiddenFlag,
		FileType: fileType,
		ShowSize: *fileSizeFlag,
	}

	if !*timeFlag {
		switch *formFlag {
		case availableForms[0]:
			forms.TableOutput(expandedPath, flags)
		case availableForms[1]:
			forms.ListOutput(expandedPath, flags)
		case availableForms[3]:
			forms.TreeOutput(expandedPath, flags)
		case availableForms[2]:
			f.FoundAllFilesInDir(expandedPath, flags)
			fmt.Println(f.TotalLines)
		}
	} else {
		switch *formFlag {
		case availableForms[0]:
			t := time.Now()
			forms.TableOutput(expandedPath, flags)
			duration := time.Since(t)
			fmt.Println()
			log.Infof("Execution time: %v", duration)
		case availableForms[1]:
			t := time.Now()
			forms.ListOutput(expandedPath, flags)
			duration := time.Since(t)
			fmt.Println()
			log.Infof("Execution time: %v", duration)
		case availableForms[2]:
			t := time.Now()
			f.FoundAllFilesInDir(path, flags)
			fmt.Println(f.TotalLines)
			duration := time.Since(t)
			fmt.Println()
			log.Infof("Execution time: %v", duration)
		case availableForms[3]:
			t := time.Now()
			forms.TreeOutput(path, flags)
			duration := time.Since(t)
			fmt.Println()
			log.Infof("Execution time: %v", duration)
		}
	}
}

func ExpandPath(path string) (string, error) {
	if len(path) > 0 && path[0] == '~' {
		usr, err := user.Current()
		if err != nil {
			return "", fmt.Errorf("Failed to get current user: %v", err)
		}
		return filepath.Join(usr.HomeDir, path[1:]), nil
	} else if path == "." {
		return os.Getwd()
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("Failed to get absolute path: %v", err)
	}

	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		return "", fmt.Errorf("Path does not exist")
	}

	return absPath, nil
}
