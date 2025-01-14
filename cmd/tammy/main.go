package main

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	help "tammy/internal"
	"tammy/internal/fileutils"
	"tammy/internal/forms"
	"time"

	"github.com/charmbracelet/log"
)

func main() {
	t := time.Now()

	availableForms := []string{"table", "list", "total", "tree"}

	var allwaysShowHiddenFiles bool
	var allwaysDisplaySize bool
	var relativePath bool
	envars := make(map[string]string)
	envars["defaultForm"] = os.Getenv("DEFAULT_FORM")
	envars["allwaysDisplaySize"] = os.Getenv("ALLWAYS_DISPLAY_SIZE")
	envars["allwaysShowHiddenFiles"] = os.Getenv("ALLWAYS_SHOW_HIDDEN_FILES")
	envars["listEnumerator"] = os.Getenv("LIST_ENUMERATOR")
	envars["treeEnumerator"] = os.Getenv("TREE_ENUMERATOR")
	envars["relativePath"] = os.Getenv("RELATIVE_PATH")
	if envars["defaultForm"] == "" {
		envars["defaultForm"] = "table"
	}
	if envars["allwaysDisplaySize"] == "true" {
		allwaysShowHiddenFiles = true
	} else {
		allwaysDisplaySize = false
	}
	if envars["allwaysShowHiddenFiles"] == "true" {
		allwaysShowHiddenFiles = true
	} else {
		allwaysShowHiddenFiles = false
	}
	if envars["listEnumerator"] == "" {
		envars["listEnumerator"] = "default_enumerator"
	}
	if envars["treeEnumerator"] == "" {
		envars["treeEnumerator"] = "default_enumerator"
	}
	if envars["relativePath"] == "true" {
		relativePath = true
	} else {
		relativePath = false
	}

	var (
		formFlag          = flag.String("f", envars["defaultForm"], "Available forms: "+strings.Join(availableForms, ", "))
		pathFlag          = flag.String("p", ".", "Path")
		fileExtFlag       = flag.String("ft", "", "Count files with file type")
		ignoreFileExtFlag = flag.String("i", "", "Ignore files with file type")
		relative          = flag.Bool("rel", relativePath, "Use relative path to file name")
		timeFlag          = flag.Bool("time", false, "Benchmark")
		showHiddenFlag    = flag.Bool("h", allwaysShowHiddenFiles, "Show hidden files")
		fileSizeFlag      = flag.Bool("s", allwaysDisplaySize, "Show size of files")
		showHelpMessage   = flag.Bool("help", false, "Show help message")
		version           = flag.Bool("version", false, "Check version")
	)
	flag.Parse()

	if *version {
		updatesAvailable, err := help.CheckForUpdates()
		if err != nil {
			fmt.Println("Error checking for updates:", err)
			os.Exit(1)
		}
		if updatesAvailable {
			log.Info("New updates are available!\n Repo link: https://github.com/nightvalley/tammy")
		} else {
			log.Info("You are using the latest version.")
		}
		os.Exit(0)
	}

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

	fileType := formatFileType(fileExtFlag)
	ignoreFiles := formatFileType(ignoreFileExtFlag)

	f := fileutils.Files{}
	flags := fileutils.Flags{
		Hidden:                *showHiddenFlag,
		ShowSize:              *fileSizeFlag,
		Relative:              *relative,
		FileType:              fileType,
		IgnoredFileExtensions: ignoreFiles,
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
			f.ExploreDirectory(expandedPath, flags)
			fmt.Println(f.TotalLines)
		}
	} else {
		switch *formFlag {
		case availableForms[0]:
			forms.TableOutput(expandedPath, flags)
			duration := time.Since(t)

			fmt.Println()
			log.Infof("Execution time: %v", duration)
		case availableForms[1]:
			forms.ListOutput(expandedPath, flags, envars["listEnumerator"])
			duration := time.Since(t)

			fmt.Println()
			log.Infof("Execution time: %v", duration)
		case availableForms[2]:
			f.ExploreDirectory(path, flags)
			fmt.Println(f.TotalLines)
			duration := time.Since(t)

			fmt.Println()
			log.Infof("Execution time: %v", duration)
		case availableForms[3]:
			forms.TreeOutput(path, flags, envars["treeEnumerator"])
			duration := time.Since(t)

			fmt.Println()
			log.Info("Execution time: %v", duration)
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

func formatFileType(flag *string) string {
	if flag != nil && *flag != "" {
		if (*flag)[0] != '.' {
			return fmt.Sprintf(".%s", *flag)
		}
		return *flag
	}
	return ""
}
