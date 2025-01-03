package forms

import (
	"fmt"
	"path/filepath"
	"strings"
	"tammy/internal/files"

	"github.com/charmbracelet/lipgloss/tree"
)

func TreeOutput(expandedPath string, flags files.Flags, enumerator string) {
	files := files.Files{}
	files.FoundAllFilesInDir(expandedPath, flags)

	t := tree.Root(expandedPath)

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

	for i, fileName := range files.Name {
		t.Child(
			filepath.Base(fileName),
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
