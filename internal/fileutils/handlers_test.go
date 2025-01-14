package fileutils_test

import (
	"fmt"
	"sort"
	"tammy/internal/fileutils"
	"testing"
)

func TestFiles_ExploreDirectory(t *testing.T) {
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
			path:   "/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles",
			want: []string{
				"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/b.txt",
				"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/a.json",
				"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/sisya.pisya",
			},
		},
		{
			name:   "with hidden files",
			hidden: true,
			path:   "/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles",
			want: []string{
				"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/b.txt",
				"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/a.json",
				"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/sisya.pisya",
				"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/.pipiska.pipiska",
			},
		},
		{
			name:     "with file type",
			hidden:   false,
			filetype: ".json",
			path:     "/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles",
			want: []string{
				"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/a.json",
			},
		},
		// {
		// 	name:     "binary file",
		// 	hidden:   false,
		// 	filetype: "",
		// 	path:     "/home/username/Development/Golang/Cli/tammy/bin",
		// 	want: []string{
		// 		"",
		// 	},
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var f fileutils.Files
			flags := fileutils.Flags{
				Hidden:   tt.hidden,
				FileType: tt.filetype,
			}

			f.ExploreDirectory(tt.path, flags)

			sort.Strings(f.Name)
			sort.Strings(tt.want)

			if len(f.Name) != len(tt.want) {
				t.Errorf("\nReturned: %v\n Want: %v", len(f.Name), len(tt.want))
				for _, file := range f.Name {
					fmt.Println("returned file: ", file)
				}
				fmt.Println("")
				for _, file := range tt.want {
					fmt.Println("want: ", file)
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
