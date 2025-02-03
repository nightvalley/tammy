package forms

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
)

const (
	borderColor = lipgloss.ANSIColor(2)
	itemColor   = lipgloss.ANSIColor(4)
)

var (
	lineIcon       = "󰉸 "
	sizeIcon       = " "
	fileIcon       = "󰢪 "
	totalLinesIcon = " "
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

func SetIcon(path string) string {
	switch filepath.Ext(path) {
	case ".go", ".mod", ".sum":
		return " "
	case ".rs":
		return " "
	case ".md":
		return " "
	}

	switch filepath.Base(path) {
	case "Dockerfile":
		return " "
	case "Makefile":
		return " "
	}

	return fileIcon
}
