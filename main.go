package main

import (
	"fmt"
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
			// fmt.Println("if ", file)
			fmt.Println("else ", file)
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
