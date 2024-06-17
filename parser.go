package main

type parser struct {
	tokens     []token
	tokIdx     int
	currentTok token
}

func NewParser(tokens []token) *parser {
	p := parser{tokens: tokens, tokIdx: -1}
	p.next()

	return &p
}

func (p *parser) parse() string {
	return p.expr()
}

func (p *parser) next() token {
	p.tokIdx += 1
	if p.tokIdx < len(p.tokens) {
		p.currentTok = p.tokens[p.tokIdx]
	}

	return p.currentTok
}

func (p *parser) factor() string {
	tok := p.currentTok
	if tok._type == TT_INT || tok._type == TT_FLOAT {
		p.next()
		return NewNumberNode(tok).toString()
	}

	return ""
}

func (p *parser) term() string {
	return opGeneric(p, p.factor, termOPTok)
}

func (p *parser) expr() string {
	return opGeneric(p, p.term, exprOPTok)
}

func opGeneric(p *parser, f func() string, keys map[string]bool) string {
	left := f()
	for {

		if !keys[p.currentTok._type] {
			break
		}

		opTok := p.currentTok
		p.next()
		right := f()
		left = NewBinOpNode(left, opTok, right).toString()
	}

	return left
}
