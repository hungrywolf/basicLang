package main

import (
	"bufio"
	"fmt"
	"os"
)

type shell struct{}

func newShell() *shell {
	return &shell{}
}

func (s *shell) startShell() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Basic> ")
		cmd, err := reader.ReadBytes('\n')
		if err != nil {
			fmt.Println(err)
		} else {
			l := newLexer(string(cmd))
			tokens, err := l.make_token()
			if err != nil {
				panic(err)
			}

			for _, token := range tokens {
				fmt.Println(token.toString())
			}
		}
	}
}
