package forms

import (
	"fmt"
	"strings"
	"tammy/internal/fileutils"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

func TreeOutput(expandedPath string, flags fileutils.Flags, enumerator string) {
	files := fileutils.Files{}
	files.ExploreDirectory(expandedPath, flags)

	t := tree.Root(".").
		EnumeratorStyle(
			lipgloss.NewStyle().
				Foreground(firstColor).
				BorderForeground(firstColor).
				Align(lipgloss.Center).PaddingRight(1))

	switch strings.ToLower(enumerator) {
	case "default_enumerator":
		t.Enumerator(tree.DefaultEnumerator)
	case "default_indenter":
		t.Enumerator(tree.DefaultIndenter)
	case "rounded":
		t.Enumerator(tree.RoundedEnumerator)
	default:
		t.Enumerator(tree.RoundedEnumerator)
	}

	for i, fName := range files.Name {
		t.Child(
			cutPath(fName, flags.Relative),
			tree.New().Child(
				fmt.Sprintf("Lines: %d", files.Lines[i]),
			),
		)

		if flags.ShowSize {
			c := tree.New().Child(
				fmt.Sprintf("Size: %.2f %s", files.Size[i].Size, files.Size[i].Unit),
			)
			t.Child(c)
		}
	}

	t.Child(fmt.Sprintf("Total lines: %d", files.TotalLines))
	fmt.Println(t)
}
