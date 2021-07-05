// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

type Remark struct {
	remark string
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	r := newRemark(remark)

	switch {
	case r.isBlank():
		return "Fine. Be that way!"
	case r.isYelledQuestion():
		return "Calm down, I know what I'm doing!"
	case r.isQuestion():
		return "Sure."
	case r.isCaps():
		return "Whoa, chill out!"
	default:
		return "Whatever."
	}
}

func (r Remark) isBlank() bool {
	return r.remark == ""
}

func (r Remark) isQuestion() bool {
	return strings.HasSuffix(r.remark, "?")
}

func (r Remark) isCaps() bool {
	hasLetters := strings.IndexFunc(r.remark, unicode.IsLetter) >= 0
	isUpperCase := strings.ToUpper(r.remark) == r.remark

	return hasLetters && isUpperCase
}

func (r Remark) isYelledQuestion() bool {
	return r.isQuestion() && r.isCaps()
}

func newRemark(remark string) Remark {
	return Remark{strings.TrimSpace(remark)}
}
