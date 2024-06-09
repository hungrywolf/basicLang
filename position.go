package main

type position struct {
	idx   int
	col   int
	ln    int
	fn    int
	ftext string
}

func NewPosition() *position {
	return &position{idx: -1, col: -1, ln: 0}
}

func (p *position) nextLine(char rune) *position {
	p.idx += 1
	p.col += 1

	if char == '\n' {
		p.ln += 1
		p.col = 0
	}

	return p
}

func (p position) copy() *position {
	return &position{idx: p.idx, col: p.col, ln: p.ln, fn: p.fn, ftext: p.ftext}
}
