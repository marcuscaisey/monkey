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
	Illegal TokenType = "ILLEGAL"
	EOF     TokenType = "EOF"

	// Identifiers and literals
	Ident TokenType = "IDENT" // add, foobar, x, y
	Int   TokenType = "INT"   // 1, 2, 234234

	// Operators
	Assign TokenType = "ASSIGN"
	Plus   TokenType = "PLUS"

	// Delimiters
	Comma     TokenType = "COMMA"
	Semicolon TokenType = "SEMICOLON"

	LParen TokenType = "LPAREN"
	RParen TokenType = "RPAREN"
	LBrace TokenType = "LBRACE"
	RBrace TokenType = "RBRACE"

	// Keywords
	Function TokenType = "FUNCTION"
	Let      TokenType = "LET"
)

var keywordTokenTypesByIdent = map[string]TokenType{
	"fn":  Function,
	"let": Let,
}

// IdentTokenType returns the token type of the given identifier. Identifiers can either be regular identifiers or they
// can be keywords.
func IdentTokenType(ident string) TokenType {
	if tokenType, ok := keywordTokenTypesByIdent[ident]; ok {
		return tokenType
	}
	return Ident
}
