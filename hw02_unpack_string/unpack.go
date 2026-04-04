package hw02unpackstring

import (
	"errors"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var result strings.Builder
	var prevChar rune
	hasPrev := false

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			if !hasPrev {
				return "", ErrInvalidString
			}
			count := int(ch - '0')
			if count == 0 {
			} else {
				for i := 0; i < count; i++ {
					result.WriteRune(prevChar)
				}
			}
			hasPrev = false
		} else {
			if hasPrev {
				result.WriteRune(prevChar)
			}
			prevChar = ch
			hasPrev = true
		}
	}

	if hasPrev {
		result.WriteRune(prevChar)
	}

	return result.String(), nil
}
