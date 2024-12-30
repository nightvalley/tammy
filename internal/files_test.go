package files_test

import (
	files "CountLines/internal"
	"fmt"
	"sort"
	"testing"
)

func TestFiles_FoundAllFilesInDir(t *testing.T) {
	tests := []struct {
		name string
		path string
		want []string
	}{
		{
			name: "test files",
			path: "/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles",
			want: []string{
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/b.txt",
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/a.json",
				"/home/username/Development/Golang/Pet-Projects/Cli/CountLines/cmd/testfiles/pipiska.jopa",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f files.Files
			f.FoundAllFilesInDir(tt.path)

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
