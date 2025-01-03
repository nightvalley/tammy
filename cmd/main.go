package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	help "tammy/internal"
	"tammy/internal/files"
	"tammy/internal/forms"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	availableForms := []string{"table", "list", "total", "tree"}

	var allwaysShowHiddenFiles bool
	var allwaysDisplaySize bool
	envars := make(map[string]string)
	envars["defaultForm"] = os.Getenv("DEFAULT_FORM")
	envars["allwaysDisplaySize"] = os.Getenv("ALLWAYS_DISPLAY_SIZE")
	envars["allwaysShowHiddenFiles"] = os.Getenv("ALLWAYS_SHOW_HIDDEN_FILES")
	envars["listEnumerator"] = os.Getenv("LIST_ENUMERATOR")
	envars["treeEnumerator"] = os.Getenv("TREE_ENUMERATOR")
	if envars["defaultForm"] == "" {
		envars["defaultForm"] = "table"
	}
	if envars["allwaysDisplaySize"] == "" {
		allwaysShowHiddenFiles = false
	} else if envars["allwaysDisplaySize"] == "true" {
		allwaysDisplaySize = true
	}
	if envars["allwaysShowHiddenFiles"] == "" {
		allwaysShowHiddenFiles = false
	} else if envars["allwaysShowHiddenFiles"] == "true" {
		allwaysShowHiddenFiles = true
	}
	if envars["listEnumerator"] == "" {
		envars["listEnumerator"] = "default_enumerator"
	}
	if envars["treeEnumerator"] == "" {
		envars["treeEnumerator"] = "default_enumerator"
	}

	var (
		formFlag        = flag.String("f", envars["defaultForm"], "Available forms: "+strings.Join(availableForms, ", "))
		pathFlag        = flag.String("p", ".", "Path")
		filetypeFlag    = flag.String("ft", "", "Count files with file type")
		timeFlag        = flag.Bool("time", false, "Benchmark")
		showHiddenFlag  = flag.Bool("h", allwaysShowHiddenFiles, "Show hidden files")
		fileSizeFlag    = flag.Bool("s", allwaysDisplaySize, "Show size of files")
		showHelpMessage = flag.Bool("help", false, "Show help message")
		// version         = flag.Bool("version", false, "Check version")
	)
	flag.Parse()

	if *showHelpMessage {
		help.ShowHelpMessage()
		os.Exit(0)
	}

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
			forms.ListOutput(expandedPath, flags, envars["listEnumerator"])
		case availableForms[3]:
			forms.TreeOutput(expandedPath, flags, envars["treeEnumerator"])
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
			forms.ListOutput(expandedPath, flags, envars["listEnumerator"])
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
			forms.TreeOutput(path, flags, envars["treeEnumerator"])
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
