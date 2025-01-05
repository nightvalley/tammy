package forms

import (
	"fmt"
	"os"
	"path/filepath"
	"tammy/internal/files"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const (
	firstColor  = lipgloss.Color("5")
	secondColor = lipgloss.Color("240")
	thirdColor  = lipgloss.Color("250")
)

func TableOutput(expandedPath string, flags files.Flags) {
	files := files.Files{}
	files.ExploreDirectory(expandedPath, flags)

	re := lipgloss.NewRenderer(os.Stdout)

	fileNameLen := 0
	for _, name := range files.Name {
		if len(filepath.Base(name)) > fileNameLen {
			fileNameLen = len(filepath.Base(name)) + 10
		}
	}

	HeaderStyle := re.NewStyle().Foreground(firstColor).Bold(true).Align(lipgloss.Center)
	CellStyle := re.NewStyle().Padding(0, 2)
	OddRowStyle := CellStyle.Foreground(secondColor)
	EvenRowStyle := CellStyle.Foreground(thirdColor)
	BorderStyle := lipgloss.NewStyle().Foreground(firstColor)

	var rows [][]string
	if flags.ShowSize {
		for i, name := range files.Name {
			size := files.Size[i]
			rows = append(rows, createRow(filepath.Base(name), files.Lines[i], fmt.Sprintf("%.2f %s", size.Size, size.Unit)))
		}
	} else {
		for i, name := range files.Name {
			rows = append(rows, createRow(filepath.Base(name), files.Lines[i]))
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
		Headers("File name", "Lines").
		Rows(rows...)

	fmt.Println(t)
}

func createRow(name string, lines int, size ...string) []string {
	if len(size) > 0 {
		return []string{name, fmt.Sprintf("%d", lines), size[0]}
	}
	return []string{name, fmt.Sprintf("%d", lines)}
}
