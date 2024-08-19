package Ascii

import (
	"os"
	"testing"
)

func TestChecksum(t *testing.T) {
	// Setup: create temporary files with known content
	createTestFile := func(fileName, content string) {
		err := os.WriteFile(fileName+".txt", []byte(content), 0o644)
		if err != nil {
			t.Fatalf("Failed to write file %s: %v", fileName, err)
		}
	}

	// Cleanup: remove temporary files
	removeTestFile := func(fileName string) {
		err := os.Remove(fileName + ".txt")
		if err != nil {
			t.Fatalf("Failed to remove file %s: %v", fileName, err)
		}
	}

	tests := []struct {
		name    string
		file    string
		content string
		want    bool
	}{
		{
			name:    "file does not exist",
			file:    "nonexistent",
			content: "",
			want:    false,
		},
		{
			name:    "file with wrong checksum",
			file:    "...standard",
			content: "content with a different checksum",
			want:    false,
		},
		{
			name:    "file with wrong checksum",
			file:    "...shadow",
			content: "content with a different checksum",
			want:    false,
		},
		{
			name:    "file with wrong checksum",
			file:    "...thinkertoy",
			content: "content with a different checksum",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.want {
				// Create the file with the expected content
				createTestFile(tt.file, tt.content)
				defer removeTestFile(tt.file)
			}

			got := Checksum(tt.file)
			if got != tt.want {
				t.Errorf("Checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}
