package parser

import (
	"ostrich-interpreter/token"
	"ostrich-interpreter/lexer"
	"ostrich-interpreter/ast"
)

type Parser struct {
	l *lexer.Lexer

	curToken token.Token  // current token
	peekToken token.Token  // next token
}

// New function creates a new parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Set curToken and peekToken
	p.nextToken()
	p.nextToken()
	return p
}

// nextToken advances to the next token.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	// TODO
}
