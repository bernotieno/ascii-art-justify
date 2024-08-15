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
                filename: "standard.txt",
            },
            want: " * \n/ \\\n/ * \\\n/ ___ \\\n/_/ \\_\\\n",
        },
        {
            name: "Multiple characters input",
            args: args{
                input:    "Hello",
                filename: "standard.txt",
            },
            want: " *  *     *  *     *     \n| |_| |   | |_| |   | |    | |   | |\n|  _  |   |  _  |   | |    | |   | |\n| | | |   | | | |   | |    | |   | |\n|_| |_|   |_| |_|   |_|    |_|   |_|\n",
        },
        {
            name: "Input with newline",
            args: args{
                input:    "Hi\nWorld",
                filename: "standard.txt",
            },
            want: " *     * \n| |_| | | |\n|  _  | | |\n| | | | |_|\n|_| |_| (_)\n\n *       *     *     *     * \n| |  | | / __ \\ | |   | | | |\n| |  | |/ /  \\ \\| |   | | | |\n| |/\\| | |  | || |   | | | |\n\\__/\\__|\\_\\  /_/\\_|   |_| |_|\n",
        },
        {
            name: "Empty input",
            args: args{
                input:    "",
                filename: "standard.txt",
            },
            want: "",
        },
        {
            name: "Tampered file",
            args: args{
                input:    "Test",
                filename: "tampered.txt",
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
