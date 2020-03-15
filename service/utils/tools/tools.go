package tools

import (
	"bytes"
	"regexp"
)

func ConcatStrings(ss ...string) string {
	var buf bytes.Buffer
	for _, s := range ss {
		buf.WriteString(s)
	}
	return buf.String()
}

func MatchDigitsString(s string) (bool, error) {
	matched, err := regexp.MatchString(`^[0-9]*$`, s)
	return matched, err
}
