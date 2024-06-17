package main

import "fmt"

type NumberNode struct {
	tok token
}

func NewNumberNode(tok token) *NumberNode {
	return &NumberNode{tok: tok}
}

func (n NumberNode) toString() string {
	return fmt.Sprintf(n.tok.toString())
}

type BinOpNode struct {
	leftNode  string
	optok     token
	rightNode string
}

func NewBinOpNode(leftNode string, optok token, rightNode string) *BinOpNode {
	return &BinOpNode{leftNode: leftNode, optok: optok, rightNode: rightNode}
}

func (b BinOpNode) toString() string {
	return fmt.Sprintf("(%v , %v, %v)", b.leftNode, b.optok.toString(), b.rightNode)
}
