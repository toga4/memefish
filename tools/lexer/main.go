package main

import (
	"github.com/k0kubun/pp"
	"github.com/MakeNowJust/memefish/pkg/parser"
)

func main() {
	l := &parser.Lexer{
		Buffer:   []byte(`"\077\xFF"`),
		FilePath: "<input>",
	}

	for {
		l.NextToken()
		if l.Token.Kind == parser.TokenEOF {
			return
		}
		pp.Println(l.Token)
	}
}
