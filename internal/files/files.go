package files

import (
	"bytes"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
)

type Files struct {
	Name       []string
	TotalLines int
	Lines      []int
	Size       []int
}

type Flags struct {
	ShowSize bool
	Hidden   bool
	FileType string
}

func (files *Files) FoundAllFilesInDir(path string, flags Flags) {
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			if d.Name() == ".git" {
				return fs.SkipDir
			}
			return nil
		}

		if flags.FileType != "" && flags.FileType != filepath.Ext(path) {
			return nil
		}
		if !flags.Hidden && filepath.Base(path)[0] == '.' {
			return nil
		}

		lineCount, err := files.processFile(path)
		if err != nil {
			return err
		}

		if lineCount > 0 {
			files.Name = append(files.Name, path)
			files.Lines = append(files.Lines, lineCount)
			files.TotalLines += lineCount
		}
		return nil
	})
	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func (f *Files) processFile(filepath string) (int, error) {
	fileBytes, err := os.Open(filepath)
	if err != nil {
		return 0, err
	}
	defer fileBytes.Close()

	return lineCounter(fileBytes), nil
}

func lineCounter(r io.Reader) int {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		if err == io.EOF {
			return count
		}
		if err != nil {
			return count
		}
	}
}
