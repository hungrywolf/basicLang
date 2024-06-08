package main

import (
	"fmt"
	"unicode"
)

type lexer struct {
	text []rune
	pos  int
	char rune
}

func newLexer(text string) *lexer {
	chars := []rune(text)

	lexer := &lexer{text: chars,
		pos: -1,
	}

	lexer.next()

	return lexer
}

func (l *lexer) next() {
	l.pos += 1

	if l.pos < len(l.text) {
		l.char = l.text[l.pos]
	} else {
		l.char = 0
	}
}

func (l *lexer) make_token() ([]token, error) {
	var tokens []token
	for {
		if l.char == rune(0) {
			break
		}

		if unicode.IsDigit(l.char) {
			tokens = append(tokens, l.make_number())
		}

		switch l.char {
		case rune(13):
			l.next()
		case ' ':
			l.next()
		case '\n':
			l.next()
			return tokens, nil
		case '\t':
			l.next()

		case '+':
			tokens = append(tokens, NewToken(TT_PLUS))
			l.next()

		case '-':
			tokens = append(tokens, NewToken(TT_MINUS))
			l.next()

		case '*':
			tokens = append(tokens, NewToken(TT_MUL))
			l.next()

		case '/':
			tokens = append(tokens, NewToken(TT_DIV))
			l.next()

		case '(':
			tokens = append(tokens, NewToken(TT_LPAREN))
			l.next()

		case ')':
			tokens = append(tokens, NewToken(TT_RPAREN))
			l.next()

		default:
			fmt.Println(l.char)
			return tokens, NewIlliegalCharError(l.char)
		}
	}

	return tokens, nil
}

func (l *lexer) make_number() token {
	nStr := ""
	dotCount := 0
	for {

		if l.char == rune(0) || (!unicode.IsDigit(l.char) && l.char != '.') {
			break
		}

		if l.char == '.' {
			if dotCount == 1 {
				break
			}
			dotCount += 1
			nStr += "."
		} else {
			nStr += string(l.char)
		}

		l.next()
	}

	if dotCount == 0 {
		return NewTokenNumber(TT_INT, []rune(nStr))
	} else {
		return NewTokenNumber(TT_FLOAT, []rune(nStr))
	}
}
