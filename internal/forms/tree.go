package forms

import (
	"fmt"
	"strings"
	"tammy/internal/filehandlers"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

func TreeOutput(files filehandlers.Files, path string, enumerator string, relative, showSize bool) {
	// f := filehandlers.Files{}
	// files := f.ExploreDirectory(path)

	t := tree.Root(".").
		EnumeratorStyle(
			lipgloss.NewStyle().
				Foreground(firstColor).
				Align(lipgloss.Center).
				PaddingRight(1)).ItemStyle(
		lipgloss.NewStyle().Foreground(secondColor))

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
			cutPath(fName, relative),
			tree.New().Child(
				fmt.Sprintf("Lines: %d", files.Lines[i]),
			),
		)

		if showSize {
			c := tree.New().Child(
				fmt.Sprintf("Size: %.2f %s", files.Size[i].Size, files.Size[i].Unit),
			)
			t.Child(c)
		}
	}

	t.Child(fmt.Sprintf("Total lines: %d", files.TotalLines))
	fmt.Println(t)
}
