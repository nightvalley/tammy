package main

import (
	"bytes"
	"io"
	"log"
	"os"
)

type Files struct {
	Name  []string
	Lines int
}

func (f *Files) FoundAllFilesInDir(path string) {
	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatalf("Error reading files: %s", err)
	}

	for _, file := range files {
		if !file.IsDir() {
			fileBytes, err := os.OpenFile(path+"/"+file.Name(), os.O_RDONLY, os.ModePerm)
			if err != nil {
				log.Fatalf("Error opening file: %s", err)
			}

			defer fileBytes.Close()
			f.Name = append(f.Name, file.Name())
			f.Lines += lineCounter(fileBytes)
		}
	}
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

func main() {
	var path string

	if len(os.Args) > 1 && os.Args[1] != "" {
		path = os.Args[1]
	} else {
		var err error
		path, err = os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current dir: %s", err)
		}
	}

	files := Files{}
	files.FoundAllFilesInDir(path)
}
