package files

import (
	"bytes"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
)

type Files struct {
	TotalLines int
	Name       []string
	Lines      []int
	Size       []FileSize
}

type FileSize struct {
	Size float64
	Unit string
}

type Flags struct {
	ShowSize bool
	Hidden   bool
	FileType string
}

func (files *Files) FoundAllFilesInDir(path string, flags Flags) {
	err := filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil
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
			return nil
		}

		if lineCount > 0 {
			size := fileSize(path)
			files.Size = append(files.Size, size)
			files.Name = append(files.Name, path)
			files.Lines = append(files.Lines, lineCount)
			files.TotalLines += lineCount
		}
		return nil
	})
	if err != nil {
		log.Error(err)
	}
}

func (files *Files) processFile(filepath string) (int, error) {
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

func fileSize(path string) FileSize {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Error(err)
	}

	sizeInBytes := fileInfo.Size()
	var size FileSize

	size.Size = float64(sizeInBytes)

	switch {
	case sizeInBytes < 1024:
		size.Unit = "b"
	case sizeInBytes < 1024*1024:
		size.Size /= 1024
		size.Unit = "KB"
	case sizeInBytes < 1024*1024*1024:
		size.Size /= (1024 * 1024)
		size.Unit = "MB"
	default:
		size.Size /= (1024 * 1024 * 1024)
		size.Unit = "GB"
	}

	return size
}
