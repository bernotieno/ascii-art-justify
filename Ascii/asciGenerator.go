package Ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetLine(num int, filename string) string {
	openFile, err := os.Open(filename + ".txt")
	if err != nil {
		fmt.Println("fffff")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(openFile)
	lineNum := 0
	line := ""
	for scanner.Scan() {
		if lineNum == num {
			line = scanner.Text()
		}
		lineNum++
	}
	// fmt.Printf("|%s|+", line)
	return line
}

func AsciArtGenerator(input string, filename string) string {
	var asciiArt string
	var asciLine string

	HandleSpecialCase(input)
	if !Checksum(filename) {
		fmt.Println("Error: The file has been tampered with")
		return ""
	}
	input = strings.Replace(input, "\\n", "\n", -1)
	words := strings.Split(input, "\n")
	spaceCount := 0
	for _, word := range words {
		for i := 0; i < 8; i++ {
			if word == "" {
				spaceCount++
				if spaceCount < len(words) {
					asciiArt += "\n"
					continue
				}
			} else {
				for _, letter := range word {
					asciLine = GetLine(1+int(letter-' ')*9+i, filename)
					asciiArt += asciLine
				}
				asciiArt += "\n"
			}
		}
	}
	return asciiArt
}
