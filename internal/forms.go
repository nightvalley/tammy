package files

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func (f *Files) TableOutput(expandedPath string) {
	f.FoundAllFilesInDir(expandedPath)

	const (
		purple    = lipgloss.Color("99")
		gray      = lipgloss.Color("245")
		lightGray = lipgloss.Color("241")
	)

	re := lipgloss.NewRenderer(os.Stdout)
	var (
		HeaderStyle  = re.NewStyle().Foreground(purple).Bold(true).Align(lipgloss.Center)
		CellStyle    = re.NewStyle().Padding(0, 1).Width(14)
		OddRowStyle  = CellStyle.Foreground(gray)
		EvenRowStyle = CellStyle.Foreground(lightGray)
		BorderStyle  = lipgloss.NewStyle().Foreground(purple)
	)

	var rows [][]string
	for i, name := range f.Name {
		rows = append(rows, []string{name, fmt.Sprintf("%d", f.Lines[i])})
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

			if col == 1 {
				style = style.Width(22)
			}

			return style
		}).
		Headers("File name", "Lines").
		Rows(rows...)

	fmt.Println(t)
}

func (f *Files) TreeOutput(expandedPath string) {
	f.FoundAllFilesInDir(expandedPath)
}
