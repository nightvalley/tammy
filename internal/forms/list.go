package forms

import (
	"CountLines/internal/files"
	"fmt"
	"path/filepath"

	"github.com/charmbracelet/lipgloss/list"
)

func ListOutput(expandedPath string, flags files.Flags) {
	files := files.Files{}
	files.FoundAllFilesInDir(expandedPath, flags)

	l := list.New()
	for i, name := range files.Name {
		addFileInfoToList(l, name, files.Lines[i], files.Size[i], flags.ShowSize)
	}
	l.Item("Total lines: " + fmt.Sprintf("%d", files.TotalLines))
	l.Enumerator(list.Roman)

	fmt.Println(l)
}

func addFileInfoToList(l *list.List, name string, lines int, size files.FileSize, showSize bool) {
	if showSize {
		nameAndLines := fmt.Sprintf("%s: %d lines, %.2f %s",
			filepath.Base(name), lines,
			size.Size, size.Unit)
		l.Item(nameAndLines)
	} else {
		nameAndLines := fmt.Sprintf("%s: %d lines", filepath.Base(name), lines)
		l.Item(nameAndLines)
	}
}
