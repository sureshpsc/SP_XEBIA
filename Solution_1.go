package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(CapitalizeEveryThirdAlphanumericChar(strings.ToLower("Aspiration.!!ccom")))
}

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
