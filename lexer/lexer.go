// Package lexer contains the lexer for the Monkey language. It provides a Lexer struct can be used to turn source code
// into lexical tokens from the [token] package.
package lexer

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/marcuscaisey/monkey/token"
)

// InvalidASCIIError is returned when the lexer encounters a byte in source code which is not valid ASCII.
type InvalidASCIIError struct {
	Byte     byte
	Position int
}

func (e *InvalidASCIIError) Error() string {
	return fmt.Sprintf("lexer: invalid ASCII character %q at byte %d", e.Byte, e.Position)
}

// Lexer parses Monkey source code which must be valid ASCII.
// TODO: implement similar interface to bufio.Scanner so that we don't have to check for errors on each NextToken call.
type Lexer struct {
	src         string
	eofReturned bool
	pos         int
}

// New initialises a new Lexer with the given source code.
func New(src string) *Lexer {
	return &Lexer{
		src: src,
	}
}

// NextToken returns the next token from the source code.
// Calling repeatedly will return all of the tokens, ending with a token of type [token.EOF]. Calling after this will
// result in a panic.
// An error will be returned if the source code is not valid ASCII.
func (l *Lexer) NextToken() (token.Token, error) {
	if l.eofReturned {
		panic("lexer: NextToken called after EOF returned")
	}

	l.consumeWhitespace()
	if l.pos == len(l.src) {
		l.eofReturned = true
		return token.Token{Type: token.EOF}, nil
	}

	char := l.src[l.pos]
	if char > unicode.MaxASCII {
		return token.Token{}, &InvalidASCIIError{Byte: l.src[l.pos], Position: l.pos}
	}

	switch char {
	case '=':
		if strings.HasPrefix(l.src[l.pos:], "==") {
			l.pos += 2
			return newToken(token.Equal, "=="), nil
		}
		l.pos++
		return newToken(token.Assign, string(char)), nil
	case '+':
		l.pos++
		return newToken(token.Plus, string(char)), nil
	case '-':
		l.pos++
		return newToken(token.Minus, string(char)), nil
	case '/':
		l.pos++
		return newToken(token.Slash, string(char)), nil
	case '*':
		l.pos++
		return newToken(token.Asterisk, string(char)), nil
	case '!':
		if strings.HasPrefix(l.src[l.pos:], "!=") {
			l.pos += 2
			return newToken(token.NotEqual, "!="), nil
		}
		l.pos++
		return newToken(token.Bang, string(char)), nil
	case '<':
		l.pos++
		return newToken(token.Less, string(char)), nil
	case '>':
		l.pos++
		return newToken(token.Greater, string(char)), nil
	case ',':
		l.pos++
		return newToken(token.Comma, string(char)), nil
	case ';':
		l.pos++
		return newToken(token.Semicolon, string(char)), nil
	case '(':
		l.pos++
		return newToken(token.LParen, string(char)), nil
	case ')':
		l.pos++
		return newToken(token.RParen, string(char)), nil
	case '{':
		l.pos++
		return newToken(token.LBrace, string(char)), nil
	case '}':
		l.pos++
		return newToken(token.RBrace, string(char)), nil
	}

	if int := l.readInt(); int != "" {
		return newToken(token.Int, int), nil
	}

	if ident := l.readIdent(); ident != "" {
		tokenType := token.IdentTokenType(ident)
		return newToken(tokenType, ident), nil
	}

	l.pos++
	return newToken(token.Illegal, string(char)), nil
}

// consumeWhitespace consumes the whitespace at the current position in the source.
func (l *Lexer) consumeWhitespace() {
	for l.pos < len(l.src) && isWhitespace(l.src[l.pos]) {
		l.pos++
	}
}

func isWhitespace(char byte) bool {
	switch char {
	case '\t', '\n', '\v', '\f', '\r', ' ':
		return true
	default:
		return false
	}
}

// readInt reads the integer at the current position in the source and returns "" if there isn't one.
func (l *Lexer) readInt() string {
	int := strings.Builder{}
	for ; l.pos < len(l.src) && isNumber(l.src[l.pos]); l.pos++ {
		int.WriteString(string(l.src[l.pos]))
	}
	return int.String()
}

func isNumber(char byte) bool {
	return '0' <= char && char <= '9'
}

// readIdent reads the identifier at the current position in the source and returns "" if there isn't one.
func (l *Lexer) readIdent() string {
	ident := strings.Builder{}
	for ; l.pos < len(l.src) && isValidIdentChar(l.src[l.pos]); l.pos++ {
		ident.WriteString(string(l.src[l.pos]))
	}
	return ident.String()
}

func isValidIdentChar(char byte) bool {
	return ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') || ('0' <= char && char <= '9') || char == '_'
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
