package commandline

import (
	"flag"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/nightvalley/tammy/internal/filehandlers"
	"github.com/nightvalley/tammy/internal/forms"

	help "github.com/nightvalley/tammy/internal/help"

	"github.com/charmbracelet/log"
)

type Flags struct {
	ShowSize              bool
	Hidden                bool
	Relative              bool
	FileType              string
	IgnoredFileExtensions string
}

func (flags Flags) Launch() {
	availableForms := []string{"table", "list", "total", "tree"}
	envars := make(map[string]string)
	var allwaysShowHiddenFiles bool
	var allwaysDisplaySize bool
	var relativePath bool

	envars["defaultForm"] = os.Getenv("DEFAULT_FORM")
	envars["allwaysDisplaySize"] = os.Getenv("ALLWAYS_DISPLAY_SIZE")
	envars["allwaysShowHiddenFiles"] = os.Getenv("ALLWAYS_SHOW_HIDDEN_FILES")
	envars["listEnumerator"] = os.Getenv("LIST_ENUMERATOR")
	envars["treeEnumerator"] = os.Getenv("TREE_ENUMERATOR")
	envars["relativePath"] = os.Getenv("RELATIVE_PATH")

	if envars["defaultForm"] == "" {
		envars["defaultForm"] = "tree"
	}
	if envars["allwaysDisplaySize"] == "true" {
		allwaysDisplaySize = true
	} else {
		allwaysDisplaySize = false
	}
	if envars["allwaysShowHiddenFiles"] == "true" {
		allwaysShowHiddenFiles = true
	} else {
		allwaysShowHiddenFiles = false
	}
	if envars["relativePath"] == "true" {
		relativePath = true
	} else {
		relativePath = false
	}
	if envars["listEnumerator"] == "" {
		envars["listEnumerator"] = "default_enumerator"
	}
	if envars["treeEnumerator"] == "" {
		envars["treeEnumerator"] = "rounded"
	}

	var (
		formFlag          = flag.String("f", envars["defaultForm"], "Available forms: "+strings.Join(availableForms, ", "))
		pathFlag          = flag.String("p", ".", "Path")
		fileExtFlag       = flag.String("e", "", "Count files with file extension")
		ignoreFileExtFlag = flag.String("i", "", "Ignore files with file type")
		showHiddenFlag    = flag.Bool("h", allwaysShowHiddenFiles, "Show hidden files")
		relativePathFlag  = flag.Bool("r", relativePath, "Show relative path")
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
			log.Info(`Install new version:

git clone https://github.com/nightvalley/tammy
cd tammy 
make build
cd .. && rm -rf tammy
`)
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

	path, err := ExpandPath(path)
	if err != nil {
		log.Error(err)
		return
	}

	fileType := formatFileType(fileExtFlag)
	ignoreFiles := formatFileType(ignoreFileExtFlag)

	flags = Flags{
		Hidden:                *showHiddenFlag,
		ShowSize:              *fileSizeFlag,
		Relative:              *relativePathFlag,
		FileType:              fileType,
		IgnoredFileExtensions: ignoreFiles,
	}

	f := filehandlers.Files{}
	f = f.ExploreDirectory(path, flags.FileType, flags.IgnoredFileExtensions, flags.Hidden)
	if f.TotalLines == 0 {
		log.Fatalf("Directory '%s' does not contain any files in which lines can be counted.", path)
	}

	switch *formFlag {
	case availableForms[0]:
		forms.TableOutput(f, path, flags.Relative, flags.ShowSize)
	case availableForms[1]:
		forms.ListOutput(f, path, envars["listEnumerator"], flags.Relative, flags.ShowSize)
	case availableForms[3]:
		forms.TreeOutput(f, path, envars["treeEnumerator"], flags.Relative, flags.ShowSize)
	case availableForms[2]:
		fmt.Println(f.TotalLines)
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

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", fmt.Errorf("Path does not exist")
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("Failed to get absolute path: %v", err)
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
