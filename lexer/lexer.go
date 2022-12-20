// Package lexer contains the lexer for the Monkey language. It provides a Lexer struct can be used to turn source code
// into lexical tokens from the [token] package.
package lexer

import (
	"fmt"
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
	if l.pos == len(l.src) {
		l.eofReturned = true
		return token.Token{Type: token.EOF}, nil
	}
	char := l.src[l.pos]
	if char > unicode.MaxASCII {
		return token.Token{}, &InvalidASCIIError{Byte: l.src[l.pos], Position: l.pos}
	}
	l.pos++
	var tokenType token.TokenType
	switch char {
	case '=':
		tokenType = token.Assign
	case '+':
		tokenType = token.Plus
	case '(':
		tokenType = token.LParen
	case ')':
		tokenType = token.RParen
	case '{':
		tokenType = token.LBrace
	case '}':
		tokenType = token.RBrace
	case ',':
		tokenType = token.Comma
	case ';':
		tokenType = token.Semicolon
	default:
		tokenType = token.Illegal
	}
	return token.Token{Type: tokenType, Literal: string(char)}, nil
}
