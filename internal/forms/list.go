package forms

import (
	"fmt"
	"path/filepath"
	"strings"
	"tammy/internal/files"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

func ListOutput(expandedPath string, flags files.Flags, enumerator string) {
	files := files.Files{}
	files.FoundAllFilesInDir(expandedPath, flags)

	l := list.New().EnumeratorStyle(lipgloss.NewStyle().Foreground(firstColor).BorderForeground(firstColor))
	for i, name := range files.Name {
		addFileInfoToList(l, name, files.Lines[i], files.Size[i], flags.ShowSize)
	}
	l.Item("Total lines: " + fmt.Sprintf("%d", files.TotalLines))

	switch strings.ToLower(enumerator) {
	case "roman":
		l.Enumerator(list.Roman)
	case "arabic":
		l.Enumerator(list.Arabic)
	case "dash":
		l.Enumerator(list.Dash)
	case "alphabet":
		l.Enumerator(list.Alphabet)
	case "bullet":
		l.Enumerator(list.Bullet)
	case "asterisk":
		l.Enumerator(list.Asterisk)
	default:
		l.Enumerator(list.Roman)
	}

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
