package repl

import (
	"bufio"
	"fmt"
	"io"
	"ostrich-interpreter/token"
	"ostrich-interpreter/lexer"
)

const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		tok := l.NextToken()
		for ; tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}