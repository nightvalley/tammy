package files

import (
	"bytes"
	"io"
	"log"
	"os"
	"sync"
)

type Files struct {
	FileTypes string
	Name      []string
	Lines     int
}

func (f *Files) FoundAllFilesInDir(path string) {
	var wg sync.WaitGroup

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error reading files: %s", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			wg.Add(1)

			go func() {
				defer wg.Done()

				f.processFile(path + "/" + file.Name())
			}()
		} else if file.IsDir() {
			if file.Name() != ".git" {
				wg.Add(1)

				go func() {
					defer wg.Done()

					f.processDirectory(path + "/" + file.Name())
				}()
			}
		}
	}

	wg.Wait()

	// fmt.Println(f.Name)

	// for _, name := range f.Name {
	// 	fmt.Println(name)
	// }
}

func (f *Files) processDirectory(directory string) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Fatalf("Error reading subdirectory: %s", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			f.processFile(directory + "/" + file.Name())
		}
	}
}

func (f *Files) processFile(filepath string) {
	fileBytes, err := os.OpenFile(filepath, os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("Error opening file: %s", err)
	}

	defer fileBytes.Close()
	f.Name = append(f.Name, filepath)
	f.Lines += lineCounter(fileBytes)
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
