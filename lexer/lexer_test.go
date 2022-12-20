package lexer_test

import (
	"errors"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/marcuscaisey/monkey/lexer"
	"github.com/marcuscaisey/monkey/token"
)

func TestNextTokenWithSimpleInput(t *testing.T) {
	src := "=+(){},;"
	want := []token.Token{
		{Type: token.Assign, Literal: "="},
		{Type: token.Plus, Literal: "+"},
		{Type: token.LParen, Literal: "("},
		{Type: token.RParen, Literal: ")"},
		{Type: token.LBrace, Literal: "{"},
		{Type: token.RBrace, Literal: "}"},
		{Type: token.Comma, Literal: ","},
		{Type: token.Semicolon, Literal: ";"},
		{Type: token.EOF},
	}
	got := mustReadAllTokens(t, src)
	equalTokens(t, got, want, src)
}

func TestNextTokenReturnsEOFIfSourceCodeIsEmpty(t *testing.T) {
	want := []token.Token{{Type: token.EOF}}
	got := mustReadAllTokens(t, "")
	equalTokens(t, got, want, "")
}

func TestNextTokenPanicsIfCalledAfterEOFReturned(t *testing.T) {
	testCases := []struct {
		name string
		src  string
	}{
		{"ForEmptySourceCode", ""},
		{"ForNonEmptySourceCode", "+"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lexer := lexer.New(tc.src)
			mustReadAllTokensFromLexer(t, lexer)
			defer func() {
				want := "lexer: NextToken called after EOF returned"
				if got := recover(); got == nil {
					t.Fatalf("NextToken() should have panicked")
				} else if got != want {
					t.Fatalf("NextToken() panicked with %q, want %q", got, want)
				}
			}()
			lexer.NextToken()
		})
	}
}

func TestNextTokenReturnsInvalidUTF8ErrorIfSourceCodeIsNotValidUTF8(t *testing.T) {
	src := "=\xFF"
	wantErr := &lexer.InvalidUTF8Error{
		Byte:     0xFF,
		Position: 1,
	}

	l := lexer.New(src)
	l.NextToken()

	nextToken, gotErr := l.NextToken()
	if gotErr == nil {
		t.Fatalf("NextToken() should have returned an error, got (%+v, %v)", nextToken, gotErr)
	}
	var invalidUTF8Err *lexer.InvalidUTF8Error
	if !errors.As(gotErr, &invalidUTF8Err) {
		t.Fatalf("NextToken() returned error %q of type %T, should have been type %T", gotErr, gotErr, wantErr)
	}
	if diff := cmp.Diff(wantErr, gotErr); diff != "" {
		t.Fatalf("NextToken() returned incorrect error from source %q\ndiff:\n--- want\n+++ got\n%s", src, diff)
	}
}

func mustReadAllTokens(t *testing.T, src string) []token.Token {
	lexer := lexer.New(src)
	return mustReadAllTokensFromLexer(t, lexer)
}

func mustReadAllTokensFromLexer(t *testing.T, lexer *lexer.Lexer) []token.Token {
	var tokens []token.Token
	for {
		nextToken, err := lexer.NextToken()
		if err != nil {
			t.Fatal(err)
		}
		tokens = append(tokens, nextToken)
		if nextToken == (token.Token{Type: token.EOF}) {
			break
		}
	}
	return tokens
}

func equalTokens(t *testing.T, got, want []token.Token, src string) {
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("NextToken() returned incorrect tokens from source %q\ndiff:\n--- want\n+++ got\n%s", src, diff)
	}
}
