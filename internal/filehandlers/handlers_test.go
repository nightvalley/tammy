package filehandlers_test

import (
	"reflect"
	"testing"

	"github.com/nightvalley/tammy/internal/filehandlers"
)

func TestFiles_ExploreDirectory(t *testing.T) {
	tests := []struct {
		name      string // описание теста
		path      string
		fileExt   string
		ignoreExt string
		hidden    bool
		want      filehandlers.Files // ожидаемый результат
	}{
		{
			name:   "without hidden files",
			hidden: false,
			path:   "/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles",
			want: filehandlers.Files{
				Name: []string{
					"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/b.txt",
					"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/a.json",
					"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/sisya.pisya",
				},
				TotalLines: 10,             // Укажите ожидаемое количество строк
				Lines:      []int{3, 4, 3}, // Укажите ожидаемое количество строк для каждого файла
				Size: []filehandlers.FileSize{
					{Size: 1.2, Unit: "KB"}, // Укажите ожидаемый размер для каждого файла
					{Size: 2.5, Unit: "KB"},
					{Size: 0.8, Unit: "KB"},
				},
			},
		},
		{
			name:   "with hidden files",
			hidden: true,
			path:   "/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles",
			want: filehandlers.Files{
				Name: []string{
					"/home/username/Development/Golang/Cli/tammy/internal/commandline/testfiles/testfiles/with-lines/b.txt",
					"/home/username/Development/Golang/Cli/tammy/internal/commandline/testfiles/testfiles/with-lines/a.json",
					"/home/username/Development/Golang/Cli/tammy/internal/commandline/testfiles/testfiles/with-lines/sisya.pisya",
					"/home/username/Development/Golang/Cli/tammy/internal/commandline/testfiles/testfiles/with-lines/.pipiska.pipiska",
				},
				TotalLines: 12,                // Укажите ожидаемое количество строк
				Lines:      []int{3, 4, 3, 2}, // Укажите ожидаемое количество строк для каждого файла
				Size: []filehandlers.FileSize{
					{Size: 1.2, Unit: "KB"},
					{Size: 2.5, Unit: "KB"},
					{Size: 0.8, Unit: "KB"},
					{Size: 0.5, Unit: "KB"}, // Размер для скрытого файла
				},
			},
		},
		{
			name:    "with file type",
			hidden:  false,
			fileExt: ".json",
			path:    "/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles",
			want: filehandlers.Files{
				Name: []string{
					"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/a.json",
				},
				TotalLines: 4,        // Укажите ожидаемое количество строк
				Lines:      []int{4}, // Укажите ожидаемое количество строк для файла
				Size: []filehandlers.FileSize{
					{Size: 2.5, Unit: "KB"}, // Укажите ожидаемый размер
				},
			},
		},
		{
			name:    "binary file",
			hidden:  false,
			fileExt: "",
			path:    "/home/username/Development/Golang/Cli/tammy/bin",
			want: filehandlers.Files{
				Name:       []string{"/home/username/Development/Golang/Cli/tammy/bin/with-lines/sisya.pisya"},
				TotalLines: 3,        // Укажите ожидаемое количество строк
				Lines:      []int{3}, // Укажите ожидаемое количество строк для файла
				Size: []filehandlers.FileSize{
					{Size: 1.0, Unit: "KB"},
				},
			},
		},
		{
			name:      "ignore specific file extension",
			hidden:    false,
			fileExt:   "",
			ignoreExt: ".txt",
			path:      "/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles",
			want: filehandlers.Files{
				Name: []string{
					"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/a.json",
					"/home/username/Development/Golang/Cli/tammy/cmd/tammy/testfiles/with-lines/sisya.pisya",
				},
				TotalLines: 7,
				Lines:      []int{4, 3},
				Size: []filehandlers.FileSize{
					{Size: 2.5, Unit: "KB"},
					{Size: 0.8, Unit: "KB"},
				},
			},
		},
		{
			name:   "empty directory",
			hidden: false,
			path:   "/home/username/Development/Golang/Cli/tammy/cmd/tammy/emptydir",
			want: filehandlers.Files{
				Name:       []string{},
				TotalLines: 0,
				Lines:      []int{},
				Size:       []filehandlers.FileSize{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := filehandlers.Files{}
			got := f.ExploreDirectory(tt.path, tt.fileExt, tt.ignoreExt, tt.hidden)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExploreDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}
