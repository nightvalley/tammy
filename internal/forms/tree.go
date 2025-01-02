package forms

import (
	"CountLines/internal/files"
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss/tree"
)

func TreeOutput(expandedPath string, flags files.Flags) {
	files := files.Files{}
	files.FoundAllFilesInDir(expandedPath, flags)

	t := tree.Root(expandedPath).Enumerator(tree.RoundedEnumerator)

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
