package forms

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"tammy/internal/fileutils"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
	"github.com/charmbracelet/log"
)

func TreeOutput(expandedPath string, flags fileutils.Flags, enumerator string) {
	files := fileutils.Files{}
	files.ExploreDirectory(expandedPath, flags)

	t := tree.Root(".").
		EnumeratorStyle(
			lipgloss.NewStyle().
				Foreground(firstColor).
				BorderForeground(firstColor).
				Align(lipgloss.Center).PaddingRight(1))

	switch strings.ToLower(enumerator) {
	case "default_enumerator":
		t.Enumerator(tree.DefaultEnumerator)
	case "default_indenter":
		t.Enumerator(tree.DefaultIndenter)
	case "rounded":
		t.Enumerator(tree.RoundedEnumerator)
	default:
		t.Enumerator(tree.RoundedEnumerator)
	}

	for i, fName := range files.Name {
		t.Child(
			cutPath(fName),
			tree.New().Child(
				fmt.Sprintf("Lines: %d", files.Lines[i]),
			),
		)
		if flags.ShowSize {
			tree.New().Child(
				fmt.Sprintf("Size: %.2f %s", files.Size[i].Size, files.Size[i].Unit),
			)
		}
	}

	t.Child(fmt.Sprintf("Total lines: %d", files.TotalLines))
	fmt.Println(t)
}

func cutPath(path string) string {
	current, err := os.Getwd()
	if err != nil {
		log.Error(err)
		return path
	}

	relative, err := filepath.Rel(current, path)
	if err != nil {
		log.Error(err)
		return path
	}

	return relative
}
