package main

import (
	"fmt"
	"strings"
	"unicode"
)

type Interface interface {
	TransformRune(pos int)
	GetValueAsRuneSlice() []rune
}

type request struct {
	position int
	input    string
	length   int
}

func (req *request) GetValueAsRuneSlice() []rune {
	return []rune(req.input)
}

func (req *request) TransformRune(position int) {
	if unicode.IsLetter(rune(req.input[position])) || unicode.IsNumber(rune(req.input[position])) {
		req.length += 1
		if req.length%req.position == 0 {
			req.input = req.input[:position] + strings.ToUpper(string(req.input[position])) + req.input[position+1:]
		}
	}
}

func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

func NewSkipString(position int, input string) request {
	return request{position: position, input: strings.ToLower(input), length: 0}
}

func (req request) String() string {
	return fmt.Sprintf("%v", req.input)
}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s)
}

