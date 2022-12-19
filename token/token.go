// Package token contains the types of tokens that can be present in Monkey source code.
package token

// TokenType is the type of a token.
type TokenType string

// Token represents a token. It stores the type of the token as well as its literal value.
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers and literals
	IDENT TokenType = "IDENT" // add, foobar, x, y
	INT   TokenType = "INT"   // 1, 2, 234234

	// Operators
	ASSIGN TokenType = "="
	PLUS   TokenType = "+"

	// Delimiters
	COMMA     TokenType = ","
	SEMICOLON TokenType = ";"

	LPAREN TokenType = "("
	RPAREN TokenType = ")"
	LBRACE TokenType = "{"
	RBRACE TokenType = "}"

	// Keywords
	FUNCTION TokenType = "FUNCTION"
	LET      TokenType = "LET"
)
