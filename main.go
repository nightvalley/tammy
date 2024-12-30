package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type Files struct {
	Name []string
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
			f.lineCounter(fileBytes)
		}
	}
}

func (f *Files) lineCounter(r io.Reader) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			fmt.Printf("%s: %d lines\n", f.Name, count)
			return

		case err != nil:
			fmt.Printf("%s: %d lines\n", f.Name, count)
			return
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
