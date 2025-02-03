package forms

import (
	"fmt"
	"os"

	"github.com/nightvalley/tammy/internal/filehandlers"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func TableOutput(files filehandlers.Files, path string, relative, showSize bool) {
	re := lipgloss.NewRenderer(os.Stdout)

	fileNameLen := 0
	for _, name := range files.Name {
		if len(cutPath(name, relative)) > fileNameLen {
			fileNameLen = len(cutPath(name, relative)) + 10
		}
	}

	HeaderStyle := re.NewStyle().Foreground(borderColor).Bold(true).Align(lipgloss.Center)
	CellStyle := re.NewStyle().Padding(0, 2)
	OddRowStyle := CellStyle.Foreground(itemColor)
	EvenRowStyle := CellStyle.Foreground(itemColor)
	BorderStyle := lipgloss.NewStyle().Foreground(borderColor)

	var rows [][]string
	if showSize {
		for i, name := range files.Name {
			size := files.Size[i]
			rows = append(rows, createRow(cutPath(name, relative), files.Lines[i], fmt.Sprintf("%.2f %s", size.Size, size.Unit)))
		}
	} else {
		for i, name := range files.Name {
			rows = append(rows, createRow(cutPath(name, relative), files.Lines[i]))
		}
	}

	rows = append(rows, []string{"", ""})
	rows = append(rows, createRow("Total lines", files.TotalLines))

	t := table.New().
		Border(lipgloss.ThickBorder()).
		BorderStyle(BorderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch {
			case row == table.HeaderRow:
				return HeaderStyle
			case row%2 == 0:
				return EvenRowStyle
			default:
				return OddRowStyle
			}
		}).
		Rows(rows...)

	if showSize {
		t.Headers(fileIcon+"File name", lineIcon+"Lines", sizeIcon+"Size")
	} else {
		t.Headers(fileIcon+"File name", lineIcon+"Lines")
	}

	fmt.Println(t)
}

func createRow(name string, lines int, size ...string) []string {
	if len(size) > 0 {
		return []string{name, fmt.Sprintf("%d", lines), size[0]}
	}
	return []string{name, fmt.Sprintf("%d", lines)}
}
