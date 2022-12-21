// Package ast contains the types used to build the AST (abstract syntax tree) for some Monkey source code.
package ast

import "github.com/marcuscaisey/monkey/token"

// Node is the interface that all ast nodes implement.
type Node interface {
	// TokenLiteral returns the literal value of the token associated with the node.
	TokenLiteral() string
}

// Statement is the interface that all statement nodes implement.
type Statement interface {
	Node
	// Dummy method to prevent a Node or Expression being used in place of a Statement.
	statementNode()
}

// Expression is the interface that all expression nodes implement.
type Expression interface {
	Node
	// Dummy method to prevent a Node or Statement being used in place of a Expression.
	expressionNode()
}

// Program is the root node of every AST.
type Program struct {
	Statements []Statement
}

func (p Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is the node for a statement of the form
//   let <identifier> = <expression>
type LetStatement struct {
	Token token.Token
	Name  Identifier
	Value Expression
}

func (ls LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls LetStatement) statementNode() {}

// Identifier is the node for an identifier.
type Identifier struct {
	Token token.Token
	Value string
}

func (i Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i Identifier) expressionNode() {}
