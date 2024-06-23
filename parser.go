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
	return expr(p)
}

func (p *parser) next() token {
	p.tokIdx += 1
	if p.tokIdx < len(p.tokens) {
		p.currentTok = p.tokens[p.tokIdx]
	}

	return p.currentTok
}

func factor(p *parser) string {
	tok := p.currentTok
	if tok._type == TT_INT || tok._type == TT_FLOAT {
		p.next()
		return NewNumberNode(tok).toString()
	}

	return ""
}

func term(p *parser) string {
	return opGeneric(p, factor, termOPTok)
}

func expr(p *parser) string {
	return opGeneric(p, term, exprOPTok)
}

func opGeneric(p *parser, f func(p *parser) string, keys map[string]bool) string {
	left := f(p)
	for {

		if !keys[p.currentTok._type] {
			break
		}

		opTok := p.currentTok
		p.next()
		right := f(p)
		left = NewBinOpNode(left, opTok, right).toString()
	}

	return left
}
