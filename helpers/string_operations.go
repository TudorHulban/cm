package helpers

import "strings"

func ReplEOL(s string) string {
	return strings.ReplaceAll(s, "\n", " ")
}
