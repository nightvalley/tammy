package forms

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
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

		return relative
	}

	return filepath.Base(path)
}
