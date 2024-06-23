package main

import "fmt"

type IllegalCharError struct {
	char rune
	pos  *position
	fn   string
}

func NewIlliegalCharError(c rune, p *position, fn string) *IllegalCharError {
	return &IllegalCharError{char: c, pos: p, fn: fn}
}

func (i IllegalCharError) Error() string {
	return fmt.Sprintf("Illegal Charachter in file '%v' , '%v' found at line %v and col %v.", i.fn, string(i.char), i.pos.ln, i.pos.col)
}

type InvalidSyntaxError struct {
	char rune
	pos  *position
	fn   string
}

func NewInvalidSyntaxError() *InvalidSyntaxError {
	return &InvalidSyntaxError{}
}

func (s InvalidSyntaxError) Error() string {
	return fmt.Sprintf("InvalidSyntaxError in file '%v' , '%v' found at line %v and col %v.", s.fn, string(s.char), s.pos.ln, s.pos.col)
}
