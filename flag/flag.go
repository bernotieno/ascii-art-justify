package flags

import (
    "fmt"
    "strings"
)

func CheckFlag(args []string) string {
    if len(args) == 0 || !strings.HasPrefix(args[0], "--align=") {
        Usage()
        return ""
    }

    parts := strings.SplitN(args[0], "=", 2)
    if len(parts) != 2 || parts[0] != "--align" {
        Usage()
        return ""
    }

    alignment := strings.TrimSpace(strings.ToLower(parts[1]))
    if alignment == "" || (alignment != "right" && alignment != "justify" && alignment != "left" && alignment != "center") {
        Usage()
        return ""
    }

    return alignment
}

func Usage() {
    fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
}
