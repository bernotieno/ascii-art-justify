package Ascii

import "testing"

func TestAsciArtGenerator(t *testing.T) {
	type args struct {
		input    string
		filename string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
            name: "Single character input",
			args: args{
				input:    "A",
				filename: "../standard",
			},
			want: "",
		},

		{
			name: "Multiple characters input",
			args: args{
				input:    "Hello",
				filename: "../standard",
			},
			want: "",
		},
		{
			name: "Input with newline",
			args: args{
				input:    "Hi\nWorld",
				filename: "../standard",
			},
			want: "",
		},
		{
			name: "Empty input",
			args: args{
				input:    "",
				filename: "../standard",
			},
			want: "",
		},
		{
			name: "Tampered file",
			args: args{
				input:    "Test",
				filename: "../tampered",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AsciArtGenerator(tt.args.input, tt.args.filename); got != tt.want {
				t.Errorf("AsciArtGenerator() = %v, want %v", got, tt.want)
			}
		})
	}
}
