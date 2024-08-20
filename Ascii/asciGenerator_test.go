package Ascii

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestGetLine(t *testing.T) {
	type args struct {
		num      int
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "First Line",
			args: args{num: 0, filename: "testfile"},
			want: "This is the first line.",
		},
		{
			name: "Middle Line",
			args: args{num: 2, filename: "testfile"},
			want: "This is the third line.",
		},
		{
			name: "Last Line",
			args: args{num: 3, filename: "testfile"},
			want: "This is the last line.",
		},
	}
	content := []byte("This is the first line.\nThis is the second line.\nThis is the third line.\nThis is the last line.")
	err := ioutil.WriteFile("testfile.txt", content, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove("testfile.txt")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetLine(tt.args.num, tt.args.filename); got != tt.want {
				t.Errorf("GetLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
