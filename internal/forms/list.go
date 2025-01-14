package forms

import (
	"fmt"
	"strings"
	"tammy/internal/filehandlers"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/list"
)

func ListOutput(files filehandlers.Files, path, enumerator string, relative, showSize bool) {
	l := list.New().EnumeratorStyle(
		lipgloss.NewStyle().
			Foreground(borderColor).
			BorderForeground(borderColor)).ItemStyle(
		lipgloss.NewStyle().Foreground(itemColor))
	for i, name := range files.Name {
		addFileInfoToList(l, name, files.Lines[i], files.Size[i], relative, showSize)
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

func addFileInfoToList(l *list.List, fName string, lines int, size filehandlers.FileSize, relative, showSize bool) {
	if showSize {
		nameAndLines := fmt.Sprintf("%s: %d lines, %.2f %s",
			cutPath(fName, relative), lines,
			size.Size, size.Unit)
		l.Item(nameAndLines)
	} else {
		nameAndLines := fmt.Sprintf("%s: %d lines", cutPath(fName, relative), lines)
		l.Item(nameAndLines)
	}
}
