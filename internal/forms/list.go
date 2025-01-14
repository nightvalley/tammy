package forms

import (
	"fmt"
	"strings"
	"tammy/internal/fileutils"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

func ListOutput(expandedPath string, flags fileutils.Flags, enumerator string) {
	files := fileutils.Files{}
	files.ExploreDirectory(expandedPath, flags)

	l := list.New().EnumeratorStyle(
		lipgloss.NewStyle().
			Foreground(firstColor).
			BorderForeground(firstColor)).ItemStyle(
		lipgloss.NewStyle().Foreground(secondColor))
	for i, name := range files.Name {
		addFileInfoToList(l, name, files.Lines[i], files.Size[i], flags)
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

func addFileInfoToList(l *list.List, fName string, lines int, size fileutils.FileSize, flags fileutils.Flags) {
	if flags.ShowSize {
		nameAndLines := fmt.Sprintf("%s: %d lines, %.2f %s",
			cutPath(fName, flags.Relative), lines,
			size.Size, size.Unit)
		l.Item(nameAndLines)
	} else {
		nameAndLines := fmt.Sprintf("%s: %d lines", cutPath(fName, flags.Relative), lines)
		l.Item(nameAndLines)
	}
}
