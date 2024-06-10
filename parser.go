package main

type parser struct {
	tokens []token
	tokIdx int
}

func NewParser(tokens []token) *parser {
	return &parser{tokens: tokens, tokIdx: -1}
}

func (p parser) next() *token {
	var currentTok token
	p.tokIdx += 1
	if p.tokIdx < len(p.tokens) {
		currentTok = p.tokens[p.tokIdx]
	}

	return &currentTok
}

func (p parser) factor() {}
func (p parser) term()   {}
func (p parser) expr()   {}
