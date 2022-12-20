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
	Assign   TokenType = "ASSIGN"
	Plus     TokenType = "PLUS"
	Minus    TokenType = "MINUS"
	Slash    TokenType = "SLASH"
	Asterisk TokenType = "ASTERISK"
	Bang     TokenType = "BANG"
	Less     TokenType = "LESS"
	Greater  TokenType = "GREATER"
	Equal    TokenType = "EQUAL"
	NotEqual TokenType = "NOT_EQUAL"

	// Delimiters
	Comma     TokenType = "COMMA"
	Semicolon TokenType = "SEMICOLON"

	LParen TokenType = "L_PAREN"
	RParen TokenType = "R_PAREN"
	LBrace TokenType = "L_BRACE"
	RBrace TokenType = "R_BRACE"

	// Keywords
	Function TokenType = "FUNCTION"
	Return   TokenType = "RETURN"
	Let      TokenType = "LET"
	If       TokenType = "IF"
	Else     TokenType = "ELSE"
	True     TokenType = "TRUE"
	False    TokenType = "FALSE"
)

var keywordTokenTypesByIdent = map[string]TokenType{
	"fn":     Function,
	"return": Return,
	"let":    Let,
	"if":     If,
	"else":   Else,
	"true":   True,
	"false":  False,
}

// IdentTokenType returns the token type of the given identifier. Identifiers can either be regular identifiers or they
// can be keywords.
func IdentTokenType(ident string) TokenType {
	if tokenType, ok := keywordTokenTypesByIdent[ident]; ok {
		return tokenType
	}
	return Ident
}
