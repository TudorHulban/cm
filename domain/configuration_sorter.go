package domain

import (
	"strings"
)

// Implements the sort interface.
type JustWords []OSVariableName

func (w JustWords) Len() int {
	return len(w)
}

func (w JustWords) Swap(i, j int) {
	w[i], w[j] = w[j], w[i]
}

func (w JustWords) Less(i, j int) bool {
	return strings.ToLower(string(w[i])) < strings.ToLower(string(w[j]))
}
