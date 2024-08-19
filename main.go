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
	var alignment string = "left" // Default alignment
	var input string
	args := os.Args[1:]

	// Determine if the first argument is an alignment flag
	if strings.HasPrefix(args[0], "--align=") {
		alignment = flags.CheckFlag(args)
		args = args[1:] // Remove the alignment flag from args
	} else if strings.HasPrefix(args[0], "--") {
		flags.Usage()
		return
	}

	// Handle cases based on the remaining number of arguments
	switch len(args) {
	case 1:
		input = args[0]
		banner = "standard"
	case 2:
		input = args[0]
		banner = strings.ToLower(args[1])
	default:
		flags.Usage()
		return
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
