package forms
import "CountLines/internal/files"

func ListOutput(expandedPath string) {
	f := files.Files{}
	f.FoundAllFilesInDir(expandedPath)
}
