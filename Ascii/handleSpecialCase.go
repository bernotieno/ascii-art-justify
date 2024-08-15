package Ascii

import (
    "fmt"
    "strings"
)

func HandleSpecialCase(s string) error {
    cases := []string{"\\a", "\\b", "\\t", "\\f", "\\r", "\\v"}

    for _, specialCase := range cases {
        if strings.Contains(s, specialCase) {
            return fmt.Errorf("special Case \"%s\" is not supported", specialCase)
        }
    }
    return nil
}