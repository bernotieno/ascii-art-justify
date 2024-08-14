package main

import (
	"fmt"
	"os"
	"strings"

	"ascii-art-justify/Ascii"
	align "ascii-art-justify/alighnment"
	flags "ascii-art-justify/flag"
)

func main() {
	// Check for the correct number of arguments.
	if len(os.Args) < 2 || len(os.Args) > 4 {
		flags.Usage()
		return
	}

	var banner string
	var alignment string
	var input string
	args := os.Args[1:]

	// Handle cases where alignment flag is provided.
	if len(args) == 2 && strings.HasPrefix(args[0], "--align=") {
		alignment = flags.CheckFlag(args)
		input = args[1]
		banner = "standard"
	} else if len(args) == 2 && !strings.Contains(args[0], "--") {
		input = args[0]
		banner = strings.ToLower(args[1])
		alignment = "left" // default alignment
	} else if len(args) == 3 {
		alignment = flags.CheckFlag([]string{args[0]})
		input = args[1]
		banner = strings.ToLower(args[2])
	} else {
		input = args[0]
		banner = "standard"
	}

	processInput(input, banner, alignment)
}

func processInput(input, banner, alignment string) {
	input = strings.Replace(input, "\\n", "\n", -1)

	Ascii.HandleSpecialCase(input)
	Ascii.Checksum(banner)

	if input == "\n" {
		fmt.Println()
		return
	} else if input == "" {
		return
	}

	// Split the input into lines based on newline characters.
	lines := strings.Split(input, "\n")
	spaceCount := 0

	// Iterate over each line of the input.
	for _, line := range lines {
		if line == "" {
			spaceCount++
			if spaceCount < len(lines) {
				fmt.Println()
				continue
			}
		} else {
			width := align.GetTerminalWidth()
			switch alignment {
			case "justify":
				align.AlignJustify(line, banner, width)
			case "right":
				align.AlignRight(line, banner, width)
			case "left":
				align.AlignLeft(line, banner)
			case "center":
				align.AlignCenter(line, banner, width)
			}
		}
	}
}
