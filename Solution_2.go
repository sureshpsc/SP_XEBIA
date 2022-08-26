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

//request struct would consist of input parameters that are required to process the requested input.
type request struct {
	position int
	input    string
	length   int
}

//GetValueAsRuneSlice would convert the input string to rune slice.
func (req *request) GetValueAsRuneSlice() []rune {
	return []rune(req.input)
}

//TransformRune method is used to check for each and every character in string. It checks if string is having valid character with the help of unicode library.
//If there is a valid letter or number then it looks for every alternative third character in the string.
//Once the third character position is found, it will make that character to capital case.
func (req *request) TransformRune(position int) {
	if unicode.IsLetter(rune(req.input[position])) || unicode.IsNumber(rune(req.input[position])) {
		req.length += 1
		if req.length%req.position == 0 {
			req.input = req.input[:position] + strings.ToUpper(string(req.input[position])) + req.input[position+1:]
		}
	}
}

//The MapString consists of interface as a parameter and calls the both implemented methods of an interface
func MapString(i Interface) {
	for pos := range i.GetValueAsRuneSlice() {
		i.TransformRune(pos)
	}
}

//NewSkipString would take input and position as a parameters and returns the request object.
func NewSkipString(position int, input string) request {
	return request{position: position, input: strings.ToLower(input), length: 0}
}

//The String method is used to modify the response by providing only the required output string.
func (req request) String() string {
	return fmt.Sprintf("%v", req.input)
}

func main() {
	s := NewSkipString(3, "Aspiration.com")
	MapString(&s)
	fmt.Println(s)
}

