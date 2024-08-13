package flags

import (
	"fmt"
	"strings"
)

func CheckFlag(args []string) string {
	var count int
	for _, arg := range args {
		if strings.Contains(arg, "--") {
			count++
		}
	}
	var alignment string

	if strings.HasPrefix(args[0], "--align=") {
		alignment = strings.ToLower(strings.TrimPrefix(args[0], "--align="))
	}

	if alignment != "right" &&
		alignment != "justify" &&
		alignment != "left" &&
		alignment != "center" {
		Usage()
	}
	if count > 1 {
		Usage()
	}

	return alignment
}

func Usage() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
}
