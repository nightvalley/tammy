package forms

import (
	"CountLines/internal/files"
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss/list"
)

func ListOutput(expandedPath string, flags files.Flags) {
	f := files.Files{}
	f.FoundAllFilesInDir(expandedPath, flags)

	if !flags.ShowSize {
		l := list.New()
		for i, name := range f.Name {
			nameAndLines := fmt.Sprintf("%s: %d lines", filepath.Base(name), f.Lines[i])
			l.Item(nameAndLines)
		}
		l.Item("Total lines: " + fmt.Sprintf("%d", f.TotalLines))

		l.Enumerator(list.Roman)

		fmt.Println(l)
	} else {
		l := list.New()
		for i, name := range f.Name {
			size := f.Size[i]
			nameAndLines := fmt.Sprintf("%s: %d lines, %.2f %s",
				filepath.Base(name), f.Lines[i],
				size.Size, size.Unit)
			l.Item(nameAndLines)
		}
		l.Item("Total lines: " + fmt.Sprintf("%d", f.TotalLines))

		l.Enumerator(list.Roman)

		fmt.Println(l)
	}
}
