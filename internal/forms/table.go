package forms

import (
	"CountLines/internal/files"
	"fmt"
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"golang.org/x/term"
)

const (
	firstColor  = lipgloss.Color("5")
	secondColor = lipgloss.Color("240")
	thirdColor  = lipgloss.Color("250")
)

func TableOutput(expandedPath string, flags files.Flags) {
	f := files.Files{}
	f.FoundAllFilesInDir(expandedPath, flags)

	re := lipgloss.NewRenderer(os.Stdout)

	fileNameLen := 0
	for _, name := range f.Name {
		if len(filepath.Base(name)) > fileNameLen {
			fileNameLen = len(filepath.Base(name))
		}
	}

	termWidth, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		termWidth = 80
	}

	maxFileNameWidth := min(fileNameLen, termWidth/3)
	maxLinesWidth := 22
	maxSizeWidth := 22

	var (
		HeaderStyle  = re.NewStyle().Foreground(firstColor).Bold(true).Align(lipgloss.Center)
		CellStyle    = re.NewStyle().Padding(0, 2)
		OddRowStyle  = CellStyle.Foreground(secondColor)
		EvenRowStyle = CellStyle.Foreground(thirdColor)
		BorderStyle  = lipgloss.NewStyle().Foreground(firstColor)
	)

	if !flags.ShowSize {
		var rows [][]string
		for i, name := range f.Name {
			rows = append(rows, []string{filepath.Base(name), fmt.Sprintf("%d", f.Lines[i])})
		}

		t := table.New().
			Border(lipgloss.ThickBorder()).
			BorderStyle(BorderStyle).
			StyleFunc(func(row, col int) lipgloss.Style {
				var style lipgloss.Style

				switch {
				case row == table.HeaderRow:
					return HeaderStyle
				case row%2 == 0:
					style = EvenRowStyle
				default:
					style = OddRowStyle
				}

				if col == 0 {
					style = style.Width(maxFileNameWidth)
				} else if col == 1 {
					style = style.Width(maxLinesWidth)
				}

				return style
			}).
			Headers("File name", "Lines").
			Rows(rows...)

		fmt.Println(t)
		fmt.Println("Total lines: ", f.TotalLines)
	} else {
		var rows [][]string
		for i, name := range f.Name {
			size := f.Size[i]
			rows = append(rows, []string{
				filepath.Base(name),
				fmt.Sprintf("%d", f.Lines[i]),
				fmt.Sprintf("%.2f %s", size.Size, size.Unit),
			})
		}

		t := table.New().
			Border(lipgloss.ThickBorder()).
			BorderStyle(BorderStyle).
			StyleFunc(func(row, col int) lipgloss.Style {
				var style lipgloss.Style

				switch {
				case row == table.HeaderRow:
					return HeaderStyle
				case row%2 == 0:
					style = EvenRowStyle
				default:
					style = OddRowStyle
				}

				if col == 0 {
					style = style.Width(maxFileNameWidth)
				} else if col == 1 {
					style = style.Width(maxLinesWidth)
				} else if col == 2 {
					style = style.Width(maxSizeWidth)
				}

				return style
			}).
			Headers("File name", "Lines", "Size").
			Rows(rows...)

		fmt.Println(t)
		fmt.Println("Total lines: ", f.TotalLines)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
