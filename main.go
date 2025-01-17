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
	if strings.HasPrefix(args[0], "-") {
		if !strings.HasPrefix(args[0], "--align=") {
			flags.Usage()
			return
		}
		alignment = flags.CheckFlag(args)
		if alignment == "" {
			// CheckFlag will have printed the usage message
			return
		}
		args = args[1:] // Remove the alignment flag from args
	}

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

	// Iterate over each line of the input.
	for _, line := range lines {
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
