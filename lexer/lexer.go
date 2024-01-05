package lexer

import "ostrich-interpreter/token"

type Lexer struct {
	input string
	position int // current char position
	readPosition int // the next char after the current one
	ch byte // current char
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	// TODO: Unicode support here
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
		case '=':
			tok = newToken(token.ASSIGN , l.ch)
		case ';':
			tok = newToken(token.SEMICOLON , l.ch)
		case '(':
			tok = newToken(token.LPAREN , l.ch)
		case ')':
			tok = newToken(token.RPAREN , l.ch)
		case ',':
			tok = newToken(token.COMMA , l.ch)
		case '+':
			tok = newToken(token.PLUS , l.ch)
		case '{':
			tok = newToken(token.LBRACE , l.ch)
		case '}':
			tok = newToken(token.RBRACE , l.ch)
		case 0:
			tok.Literal = ""
			tok.Type = token.EOF
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.tokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}
