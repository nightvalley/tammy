package files_test

import (
	"CountLines/internal/files"
	"fmt"
	"sort"
	"testing"
)

func TestFiles_FoundAllFilesInDir(t *testing.T) {
	tests := []struct {
		name     string
		filetype string
		path     string
		want     []string
		hidden   bool
	}{
		{
			name:   "without hidden files",
			hidden: false,
			path:   "/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles",
			want: []string{
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/with-lines/b.txt",
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/with-lines/a.json",
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/with-lines/sisya.pisya",
			},
		},
		{
			name:   "with hidden files",
			hidden: true,
			path:   "/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles",
			want: []string{
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/with-lines/b.txt",
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/with-lines/a.json",
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/with-lines/sisya.pisya",
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/with-lines/.pipiska.pipiska",
			},
		},
		{
			name:     "with file type",
			hidden:   false,
			filetype: ".json",
			path:     "/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles",
			want: []string{
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/with-lines/a.json",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f files.Files
			flags := files.Flags{
				Hidden:   tt.hidden,
				FileType: tt.filetype,
			}

			f.FoundAllFilesInDir(tt.path, flags)

			sort.Strings(f.Name)
			sort.Strings(tt.want)

			if len(f.Name) != len(tt.want) {
				t.Errorf("\nReturned: %v\n Want: %v", len(f.Name), len(tt.want))
				for _, file := range f.Name {
					fmt.Println("returned file: ", file)
				}
				return
			}

			for i, name := range f.Name {
				if name != tt.want[i] {
					t.Errorf("\nReturned: %v\nWant: %v", f.Name, tt.want)
					return
				}
			}
		})
	}
}
