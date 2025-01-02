package forms

import (
	"CountLines/internal/files"
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss/tree"
)

func TreeOutput(expandedPath string, flags files.Flags) {
	f := files.Files{}
	f.FoundAllFilesInDir(expandedPath, flags)

	t := tree.Root(expandedPath).Enumerator(tree.RoundedEnumerator)

	for i, fileName := range f.Name {
		t.Child(
			filepath.Base(fileName),
			tree.New().Child(
				fmt.Sprintf("Lines: %d", f.Lines[i]),
			),
		)
		if flags.ShowSize {
			tree.New().Child(
				fmt.Sprintf("Size: %.2f %s", f.Size[i].Size, f.Size[i].Unit),
			)
		}
	}

	t.Child(fmt.Sprintf("Total lines: %d", f.TotalLines))
	fmt.Println(t)
}

// if flags.ShowSize {
// 	for i, fileName := range f.Name {
// 		t.Child(
// 			filepath.Base(fileName),
// 			tree.New().Child(
// 				fmt.Sprintf("Lines: %d", f.Lines[i]),
// 				fmt.Sprintf("Size: %.2f %s", f.Size[i].Size, f.Size[i].Unit),
// 			),
// 		)
// 	}
// } else {
// 	for i, fileName := range f.Name {
// 		t.Child(
// 			filepath.Base(fileName),
// 			tree.New().Child(
// 				fmt.Sprintf("Lines: %d", f.Lines[i]),
// 			),
// 		)
// 	}
// }
