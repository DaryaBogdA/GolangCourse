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
	escape := false

	for _, char := range s {
		if escape {
			if char >= '0' && char <= '9' || char == '\\' {
				if hasPrev {
					result.WriteRune(prevChar)
				}
				prevChar = char
				hasPrev = true
				escape = false
				continue
			}
			return "", ErrInvalidString
		}

		if char == '\\' {
			escape = true
			continue
		}

		if char >= '0' && char <= '9' {
			if !hasPrev {
				return "", ErrInvalidString
			}

			count := int(char - '0')
			if count == 0 {
				hasPrev = false
			} else {
				for j := 0; j < count; j++ {
					result.WriteRune(prevChar)
				}
				hasPrev = false
			}
		} else {
			if hasPrev {
				result.WriteRune(prevChar)
			}
			prevChar = char
			hasPrev = true
		}
	}

	if escape {
		return "", ErrInvalidString
	}

	if hasPrev {
		result.WriteRune(prevChar)
	}

	return result.String(), nil
}
