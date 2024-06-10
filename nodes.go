package main

type NumberNode struct {
	tok token
}

func NewNumberNode(tok token) *NumberNode {
	return &NumberNode{tok: tok}
}

type BinOpNode struct {
	leftNode  token
	optok     token
	rightNode token
}

func NewBinOpNode(leftNode token, optok token, rightNode token) *BinOpNode {
	return &BinOpNode{leftNode: leftNode, optok: optok, rightNode: rightNode}
}
