package lexer

import "ostrich-interpreter/token"

type Lexer struct {
	input string
	position int // current char position
	readPosition int // the next char after the current one
	ch byte // current char
}

/*
 * Initialize a new lexer
 */
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

/*
 * Read the next character and advance the position in the input string
 */
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

/*
 * Return a token depending on the current character
 */
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
		default:
			if isLetter(l.ch) {
				tok.Literal = l.readIdentifier()
				return tok 
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

/*
 * Read and return the current identifier and advance lexer's position until
 * it encounters a non-letter-character
 */
func (l *Lexer) readIdentifier() string {
	start_position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[start_position, l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 
		'A' <= ch && ch <= 'Z' ||
		ch == '_'
}
