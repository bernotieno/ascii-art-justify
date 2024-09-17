package Ascii

import (
	"bufio"
	"fmt"
	"os"
)

func GetLine(num int, filename string) string {
	openFile, err := os.Open(filename + ".txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer openFile.Close()

	scanner := bufio.NewScanner(openFile)
	lineNum := 0
	line := ""

	for scanner.Scan() {
		if lineNum == num {
			line = scanner.Text()
			return line
		}
		lineNum++
	}

	// Check if the requested line number was found
	fmt.Println("Character not found")
	os.Exit(0)
	return ""
}
