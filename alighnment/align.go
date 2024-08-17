package align

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"ascii-art-justify/Ascii"
)

// Helper function to get ASCII art lines for the given text
func getAsciiLines(text, banner string) []string {
	// words := strings.Fields(text)
	lines := make([]string, 8)

	for i := 0; i < 8; i++ {
		for _, letter := range text {
			line := Ascii.GetLine(1+int(letter-' ')*9+i, banner)
			lines[i] += line
		}
	}
	return lines
}

func AlignJustify(text, banner string, width int) {
	words := strings.Fields(text)
	var asciiLength int

	for _, word := range words {
		for _, letter := range word {
			asciiLength += len(Ascii.GetLine(1+int(letter-' ')*9, banner))
		}
	}

	totalSpace := width - asciiLength
	if totalSpace < 0 {
		fmt.Println("Text too long to fit in terminal width.")
		return
	}
	if len(words) == 1 {
		lines := getAsciiLines(words[0], banner)
		for i := 0; i < 8; i++ {
			fmt.Println(lines[i] + strings.Repeat(" ", totalSpace))
		}
		return
	}

	spaceBetween := totalSpace / (len(words) - 1)
	extraSpace := totalSpace % (len(words) - 1)

	if len(words) > 1 {
		lines := make([]string, 8)
		for i := 0; i < 8; i++ {
			var lineOutput string
			for j, word := range words {
				for _, letter := range word {
					line := Ascii.GetLine(1+int(letter-' ')*9+i, banner)
					lineOutput += line
				}
				if j < len(words)-1 {
					lineOutput += strings.Repeat(" ", spaceBetween)
					if j < extraSpace {
						lineOutput += " "
					}
				}
			}
			lines[i] = lineOutput
		}
		for i := 0; i < 8; i++ {
			fmt.Println(lines[i])
		}
	}
}

func AlignRight(text, banner string, width int) {
	var asciiLength int
	for _, letter := range text {
		asciiLength += len(Ascii.GetLine(1+int(letter-' ')*9, banner))
	}

	totalSpace := width - asciiLength
	// Checks if the ascii length is greater than terminal width
	if totalSpace < 0 {
		fmt.Println("Text too long to fit in terminal width.")
		return
	}

	lines := getAsciiLines(text, banner)
	for i := 0; i < 8; i++ {
		fmt.Println(strings.Repeat(" ", totalSpace) + lines[i])
	}
}

func AlignLeft(text, banner string) {
	lines := getAsciiLines(text, banner)
	for i := 0; i < 8; i++ {
		fmt.Println(lines[i])
	}
}

func AlignCenter(text, banner string, width int) {
	var asciiLength int
	for _, letter := range text {
		asciiLength += len(Ascii.GetLine(1+int(letter-' ')*9, banner))
	}

	totalSpace := width - asciiLength
	// Checks if the ascii length is greater than terminal width
	if totalSpace < 0 {
		fmt.Println("Text too long to fit in terminal width.")
		return
	}

	leftSpacing := totalSpace / 2
	rightSpacing := totalSpace - leftSpacing

	lines := getAsciiLines(text, banner)
	for i := 0; i < 8; i++ {
		fmt.Println(strings.Repeat(" ", leftSpacing) + lines[i] + strings.Repeat(" ", rightSpacing))
	}
}

func GetTerminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		fmt.Println("Error getting terminal size:", err)
		os.Exit(1)
	}

	strNum := strings.Fields(string(out))

	num, err := strconv.Atoi(strNum[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return num
}
