package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	var b strings.Builder
	for i := 0; i < len(s); i++ {
		if unicode.IsDigit(rune(s[0])) {
			return "", ErrInvalidString
		}
		if unicode.IsDigit(rune(s[i])) {
			numForRepeat, err := strconv.Atoi(string(s[i]))
			if err != nil {
				return "", err
			}
			if unicode.IsNumber(rune(s[i+1])) {
				return "", ErrInvalidString
			}

			if numForRepeat == 0 {
				editString := strings.ReplaceAll(s, string(s[i-1])+string(s[i]), "")
				b.Reset()
				b.WriteString(editString)

				return b.String(), nil
			}

			stringForRepeat := strings.Repeat(string(s[i-1]), numForRepeat-1)
			editString := strings.ReplaceAll(stringForRepeat, "\n", `\n`)
			b.WriteString(editString)
		} else {
			fmt.Fprintf(&b, "%s", strings.ReplaceAll(string(s[i]), "\n", `\n`))
		}
	}
	return b.String(), nil
}
