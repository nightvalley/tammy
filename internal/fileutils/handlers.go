package fileutils

import (
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/mackerelio/go-osstat/memory"
)

type FileStatistics interface {
	fileSize(path string) FileSize
	lineCounter(r io.Reader) int
}

type CollectFileStats interface {
	ExploreDirectory(path string, flags Flags)
	processFile(filepath string) (int, error)
}

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
	ShowSize              bool
	Hidden                bool
	FileType              string
	Form                  string
	IgnoredFileExtensions string
}

const (
	Kilobyte = 1024
	Megabyte = 1024 * Kilobyte
	Gigabyte = 1024 * Megabyte
)

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

		if flags.ignoreFile(path) {
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

func (flags *Flags) ignoreFile(path string) bool {
	ignoredFileExtensions := []string{
		".png", ".jpg", ".jpeg", ".gif", ".ico",
		".bmp", ".tiff", ".svg", ".mp3", ".wav",
		".flac", ".mp4", ".avi", ".mkv", ".zip",
		".rar", ".tar", ".exe", ".dll", ".bin",
		".dat", ".ttf", ".otf",
		".xls", ".xlsx",

		".pdf",
		".doc", ".docx",
	}

	if isBinary(path) {
		return true
	}

	if flags.IgnoredFileExtensions != "" && filepath.Ext(path) == flags.IgnoredFileExtensions {
		return true
	}

	if flags.FileType != "" && flags.FileType != filepath.Ext(path) {
		return true
	}

	if !flags.Hidden && filepath.Base(path)[0] == '.' {
		return true
	}

	for _, ext := range ignoredFileExtensions {
		if filepath.Ext(path) == ext {
			return true
		}
	}

	return false
}

func isBinary(path string) bool {
	return false
}

func lineCounter(r io.Reader) int {
	buf := make([]byte, Megabyte)
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

func fileSize(path string) (float64, string) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		log.Error(err)
		return 0, ""
	}

	sizeInBytes := fileInfo.Size()
	var size float64
	var unit string

	switch {
	case sizeInBytes < Kilobyte:
		size = float64(sizeInBytes)
		unit = "b"
	case sizeInBytes < Megabyte:
		size = float64(sizeInBytes) / 1024
		unit = "KB"
	case sizeInBytes < Gigabyte:
		size = float64(sizeInBytes) / (1024 * 1024)
		unit = "MB"
	default:
		size = float64(sizeInBytes) / Gigabyte
		unit = "GB"
	}

	return size, unit
}

func calculateGoroutines(files string) (int, error) {
	memory, err := memory.Get()
	if err != nil {
		return 0, fmt.Errorf("failed to get memory: %v", err)
	}

	maxGoroutines := int(memory.Free / (1 << 20))
	if maxGoroutines > len(files) {
		maxGoroutines = len(files)
	}

	return maxGoroutines, nil
}
