package main

import (
	files "CountLines/internal"
	"log"
	"os"
)

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

	files := files.Files{}
	files.FoundAllFilesInDir(path)
}
