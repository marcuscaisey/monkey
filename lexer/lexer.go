// Package lexer contains the lexer for the Monkey language. It provides a Lexer struct can be used to turn source code
// into lexical tokens from the [token] package.
package lexer

import (
	"strings"

	"github.com/marcuscaisey/monkey/token"
)

// Lexer parses Monkey source code containing only ASCII characters.
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
// Calling repeatedly will return all of the tokens, ending with a token of type [token.EOF]. Calls after this will
// always return a [token.EOF].
func (l *Lexer) NextToken() token.Token {
	l.consumeWhitespace()
	switch char := l.readChar(); char {
	case 0:
		return newToken(token.EOF, "")
	case '=':
		if l.peekChar() == '=' {
			l.readChar()
			return newToken(token.Equal, "==")
		}
		return newToken(token.Assign, string(char))
	case '+':
		return newToken(token.Plus, string(char))
	case '-':
		return newToken(token.Minus, string(char))
	case '/':
		return newToken(token.Slash, string(char))
	case '*':
		return newToken(token.Asterisk, string(char))
	case '!':
		if l.peekChar() == '=' {
			l.readChar()
			return newToken(token.NotEqual, "!=")
		}
		return newToken(token.Bang, string(char))
	case '<':
		return newToken(token.Less, string(char))
	case '>':
		return newToken(token.Greater, string(char))
	case ',':
		return newToken(token.Comma, string(char))
	case ';':
		return newToken(token.Semicolon, string(char))
	case '(':
		return newToken(token.LParen, string(char))
	case ')':
		return newToken(token.RParen, string(char))
	case '{':
		return newToken(token.LBrace, string(char))
	case '}':
		return newToken(token.RBrace, string(char))
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return newToken(token.Int, l.readCharsWhile(char, isNumber))
	default:
		if isValidFirstIdentChar(char) {
			ident := l.readCharsWhile(char, isValidIdentChar)
			tokenType := token.IdentTokenType(ident)
			return newToken(tokenType, ident)
		}
		return newToken(token.Illegal, string(char))
	}
}

// readChar consumes the character at the current position and returns it. If the end of the source has been reached, a
// null character is returned.
func (l *Lexer) readChar() byte {
	if l.pos == len(l.src) {
		return 0
	}
	char := l.src[l.pos]
	l.pos++
	return char
}

// peekChar returns the character at the current position but doesn't consume it.
func (l *Lexer) peekChar() byte {
	if l.pos == len(l.src) {
		return 0
	}
	return l.src[l.pos]
}

// consumeWhitespace consumes the whitespace at the current position in the source.
func (l *Lexer) consumeWhitespace() {
	for isWhitespace(l.peekChar()) {
		l.readChar()
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

// readCharsWhile builds a string starting with the given character and then added to by consuming characters until
// isValid returns false.
func (l *Lexer) readCharsWhile(firstChar byte, isValid func(char byte) bool) string {
	b := strings.Builder{}
	b.WriteByte(firstChar)
	for isValid(l.peekChar()) {
		b.WriteByte(l.readChar())
	}
	return b.String()
}

func isNumber(char byte) bool {
	return '0' <= char && char <= '9'
}

func isValidFirstIdentChar(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || c == '_'
}

func isValidIdentChar(c byte) bool {
	return isValidFirstIdentChar(c) || ('0' <= c && c <= '9')
}

func newToken(tokenType token.TokenType, literal string) token.Token {
	return token.Token{Type: tokenType, Literal: literal}
}
