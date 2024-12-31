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
}

func (f *Files) FoundAllFilesInDir(path string, hidden bool, filetype string) {
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

		if filetype != "" && filetype == filepath.Ext(path) {
		}

		if hidden || (!hidden && filepath.Base(path)[0] != '.') {
			lineCount, err := f.processFile(path)
			if err != nil {
				return err
			}

			if lineCount > 0 {
				f.Name = append(f.Name, path)
				f.Lines = append(f.Lines, lineCount)
				f.TotalLines += lineCount
			}
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

func sortByFT(filetype string, path string) string {
	if filetype == filepath.Ext(path) {
		return path
	}
	return path
}

func sortHidden(path string) string {
	if filepath.Base(path)[0] != '.' {
		return path
	}
	return path
}
