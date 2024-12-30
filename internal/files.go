package files

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
)

type Files struct {
	Name       []string
	TotalLines int
	Lines      []int
	mu         sync.Mutex
}

var wg sync.WaitGroup

func (f *Files) FoundAllFilesInDir(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error reading files: %s", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)

			go func(fileName string) {
				defer wg.Done()

				f.processFile(path + "/" + fileName)
			}(file.Name())
		} else if file.IsDir() {
			if file.Name() != ".git" {
				wg.Add(1)

				go func(fileName string) {
					defer wg.Done()

					f.processDirectory(path + "/" + fileName)
				}(file.Name())
			}
		}
	}

	wg.Wait()
}

func (f *Files) processDirectory(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalf("Error reading subdirectory: %s", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)
			go func(fileName string) {
				defer wg.Done()

				f.processFile(fileName)
			}(directory + "/" + file.Name())
		}
	}
}

func (f *Files) processFile(filepath string) {
	fileBytes, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}
	defer fileBytes.Close()

	lineCount := lineCounter(fileBytes)

	f.mu.Lock()
	f.Name = append(f.Name, filepath)
	f.Lines = append(f.Lines, lineCount)
	f.mu.Unlock()
}

func lineCounter(r io.Reader) int {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count

		case err != nil:
			return count
		}
	}
}
