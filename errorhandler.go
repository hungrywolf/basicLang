package main

import "fmt"

type IllegalCharError struct {
	c rune
}

func NewIlliegalCharError(c rune) *IllegalCharError {
	return &IllegalCharError{c}
}

func (i IllegalCharError) Error() string {
	return fmt.Sprintf("Illegal Charachter '%v' found.", string(i.c))
}
