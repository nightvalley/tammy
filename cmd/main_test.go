package main

import "testing"

func Test_expandPath(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		path    string
		want    string
		wantErr bool
	}{
		{
			name:    "home directory",
			path:    "/home/username",
			want:    "/home/username",
			wantErr: false,
		},
		{
			name:    "current directory",
			path:    "",
			want:    "/home/username/Development/Golang/Pet-Projects/Cli/CountLines",
			wantErr: false,
		},
		{
			name:    "absolute path",
			path:    "/home/username/Development/Golang/Pet-Projects/Cli/CountLines",
			want:    "/home/username/Development/Golang/Pet-Projects/Cli/CountLines",
			wantErr: false,
		},
		{
			name:    "relative path",
			path:    "internal",
			want:    "/home/username/Development/Golang/Pet-Projects/Cli/CountLines",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := expandPath(tt.path)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("expandPath() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("expandPath() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("expandPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
