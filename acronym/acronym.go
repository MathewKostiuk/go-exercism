// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
	"bytes"
	"regexp"
)

// Abbreviate takes a string of words and produces an Acronym
// using the first character of each word
func Abbreviate(s string) string {
	re := regexp.MustCompile(`(?:\s|^|_|-)([a-zA-Z])`)
	sm := re.FindAllSubmatch([]byte(s), -1)
	var ba []byte

	for _, v := range sm {
		b := bytes.ToUpper(v[1])
		ba = append(ba, b...)
	}

	return string(ba)
}
