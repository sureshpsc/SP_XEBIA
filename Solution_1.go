package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(CapitalizeEveryThirdAlphanumericChar(strings.ToLower("Aspiration.!!ccom")))
}

//Below function would take string as a input paramater and converts each 3rd positioned character to upper case by ignoring the special character positions.
func CapitalizeEveryThirdAlphanumericChar(s string) string {
	rs, len := []rune(s), 0
	for i, r := range rs {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			len += 1
			if len%3 == 0 {
				rs[i] = unicode.ToUpper(r)
			}
		}
	}
	return string(rs)
}
