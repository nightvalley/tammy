package forms

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

const (
	firstColor  = lipgloss.ANSIColor(10)
	secondColor = lipgloss.ANSIColor(15)
)

func cutPath(path string, relative bool) string {
	if relative {
		current, err := os.Getwd()
		if err != nil {
			log.Error(err)
			return path
		}

		relative, err := filepath.Rel(current, path)
		if err != nil {
			log.Error(err)
			return path
		}

		clean := strings.ReplaceAll(relative, "../", "")
		return clean
	}

	return filepath.Base(path)
}
