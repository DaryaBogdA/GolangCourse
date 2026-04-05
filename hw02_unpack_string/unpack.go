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
	hasSlash := false

	for _, char := range s {
		if hasSlash {
			if hasPrev {
				result.WriteRune(prevChar)
			}
			prevChar = char
			hasPrev = true
			hasSlash = false
			continue
		}
		if char == '\\' {
			hasSlash = true
			continue
		}
		if char >= '0' && char <= '9' {
			count := int(char - '0')
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
			prevChar = char
			hasPrev = true
		}
	}

	if hasPrev {
		result.WriteRune(prevChar)
	}

	return result.String(), nil
}
