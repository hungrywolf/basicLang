package main

import "fmt"

const (
	TT_INT    = "INT"
	TT_FLOAT  = "FLOAT"
	TT_PLUS   = "PLUS"
	TT_MINUS  = "MINUS"
	TT_MUL    = "MUL"
	TT_DIV    = "DIV"
	TT_LPAREN = "LPAREN"
	TT_RPAREN = "RPAEN"
)

var termOPTok = map[string]bool{TT_MUL: true, TT_DIV: true}
var exprOPTok = map[string]bool{TT_PLUS: true, TT_MINUS: true}

type token struct {
	_type string
	value []rune
}

func NewToken(typ string) token {
	return token{_type: typ}
}

func NewTokenNumber(typ string, value []rune) token {
	return token{_type: typ, value: value}
}

func (t token) toString() string {
	if len(t.value) > 0 {
		return fmt.Sprintf("%v:%v", t._type, string(t.value))
	} else {
		return fmt.Sprintf("%v", t._type)
	}
}
