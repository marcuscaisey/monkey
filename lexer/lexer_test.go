package lexer_test

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/google/go-cmp/cmp"

	"github.com/marcuscaisey/monkey/lexer"
	"github.com/marcuscaisey/monkey/token"
)

func TestNextToken(t *testing.T) {
	testCases := []struct {
		name string
		src  string
		want []token.Token
	}{
		{
			name: "ParsesSourceWithAllTokens",
			src: `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);

!-/*5;

5 < 10 > 5;

if (5 < 10) {
    return true;
} else {
    return false;
}

10 == 10;
10 != 9;
`,
			want: []token.Token{
				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "five"},
				{Type: token.Assign, Literal: "="},
				{Type: token.Int, Literal: "5"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "ten"},
				{Type: token.Assign, Literal: "="},
				{Type: token.Int, Literal: "10"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "add"},
				{Type: token.Assign, Literal: "="},
				{Type: token.Function, Literal: "fn"},
				{Type: token.LParen, Literal: "("},
				{Type: token.Ident, Literal: "x"},
				{Type: token.Comma, Literal: ","},
				{Type: token.Ident, Literal: "y"},
				{Type: token.RParen, Literal: ")"},
				{Type: token.LBrace, Literal: "{"},
				{Type: token.Ident, Literal: "x"},
				{Type: token.Plus, Literal: "+"},
				{Type: token.Ident, Literal: "y"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.RBrace, Literal: "}"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.Let, Literal: "let"},
				{Type: token.Ident, Literal: "result"},
				{Type: token.Assign, Literal: "="},
				{Type: token.Ident, Literal: "add"},
				{Type: token.LParen, Literal: "("},
				{Type: token.Ident, Literal: "five"},
				{Type: token.Comma, Literal: ","},
				{Type: token.Ident, Literal: "ten"},
				{Type: token.RParen, Literal: ")"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.Bang, Literal: "!"},
				{Type: token.Minus, Literal: "-"},
				{Type: token.Slash, Literal: "/"},
				{Type: token.Asterisk, Literal: "*"},
				{Type: token.Int, Literal: "5"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.Int, Literal: "5"},
				{Type: token.Less, Literal: "<"},
				{Type: token.Int, Literal: "10"},
				{Type: token.Greater, Literal: ">"},
				{Type: token.Int, Literal: "5"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.If, Literal: "if"},
				{Type: token.LParen, Literal: "("},
				{Type: token.Int, Literal: "5"},
				{Type: token.Less, Literal: "<"},
				{Type: token.Int, Literal: "10"},
				{Type: token.RParen, Literal: ")"},
				{Type: token.LBrace, Literal: "{"},
				{Type: token.Return, Literal: "return"},
				{Type: token.True, Literal: "true"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.RBrace, Literal: "}"},
				{Type: token.Else, Literal: "else"},
				{Type: token.LBrace, Literal: "{"},
				{Type: token.Return, Literal: "return"},
				{Type: token.False, Literal: "false"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.RBrace, Literal: "}"},
				{Type: token.Int, Literal: "10"},
				{Type: token.Equal, Literal: "=="},
				{Type: token.Int, Literal: "10"},
				{Type: token.Semicolon, Literal: ";"},
				{Type: token.Int, Literal: "10"},
				{Type: token.NotEqual, Literal: "!="},
				{Type: token.Int, Literal: "9"},
				{Type: token.Semicolon, Literal: ";"},
			},
		},
		{
			name: "ParsesAllIntegers",
			src:  "0 1 2 3 4 5 6 7 8 9",
			want: []token.Token{
				{Type: token.Int, Literal: "0"},
				{Type: token.Int, Literal: "1"},
				{Type: token.Int, Literal: "2"},
				{Type: token.Int, Literal: "3"},
				{Type: token.Int, Literal: "4"},
				{Type: token.Int, Literal: "5"},
				{Type: token.Int, Literal: "6"},
				{Type: token.Int, Literal: "7"},
				{Type: token.Int, Literal: "8"},
				{Type: token.Int, Literal: "9"},
			},
		},
		{
			name: "ReturnsIllegalTokenTypeForUnknownCharaceters",
			src:  "\\+\\+",
			want: []token.Token{
				{Type: token.Illegal, Literal: `\`},
				{Type: token.Plus, Literal: "+"},
				{Type: token.Illegal, Literal: `\`},
				{Type: token.Plus, Literal: "+"},
			},
		},
		{
			name: "ReturnsNoTokensIfSourceCodeIsEmpty",
			src:  "",
			want: []token.Token{},
		},
		{
			name: "IgnoresASCIIWhitespaceCharacters",
			src:  "\t\n\v\f\r ",
			want: []token.Token{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			lexer := lexer.New(tc.src)
			got := []token.Token{}
			for nextToken := lexer.NextToken(); nextToken.Type != token.EOF; nextToken = lexer.NextToken() {
				got = append(got, nextToken)
			}
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Fatalf("NextToken() returned incorrect tokens from source %q\ndiff:\n--- want\n+++ got\n%s", tc.src, diff)
			}
		})
	}
}

// TODO: add test with characters after starting position
func TestNextTokenOnlyAllowsIdentsToStartWithLettersAndUnderscores(t *testing.T) {
	for i := 0; i <= unicode.MaxASCII; i++ {
		// leading whitespace will be ignored
		switch i {
		case '\t', '\n', '\v', '\f', '\r', ' ':
			continue
		}

		isValidFirstChar := false
		testName := fmt.Sprintf("%vIsNotValidFirstChar", string(rune(i)))
		if ('A' <= i && i <= 'Z') || ('a' <= i && i <= 'z') || i == '_' {
			isValidFirstChar = true
			testName = fmt.Sprintf("%vIsValidFirstChar", string(rune(i)))
		}
		t.Run(testName, func(t *testing.T) {
			src := fmt.Sprint(string(rune(i)), "a")
			lexer := lexer.New(src)
			firstToken := lexer.NextToken()
			if isValidFirstChar {
				if firstToken.Type != token.Ident {
					t.Fatalf("NextToken() = %+v for source %q, want type IDENT", firstToken, src)
				}
			} else {
				if firstToken.Type == token.Ident {
					t.Fatalf("NextToken() = %+v for source %q, should not have type IDENT", firstToken, src)
				}
			}
		})
	}
}

func TestNextTokenReturnsEOFIfCalledAfterEOFReturned(t *testing.T) {
	lexer := lexer.New("")
	want := []token.Token{{Type: token.EOF}, {Type: token.EOF}}
	got := []token.Token{lexer.NextToken(), lexer.NextToken()}
	if diff := cmp.Diff(want, got); diff != "" {
		t.Fatalf("NextToken() returned incorrect tokens when called twice on empty source\ndiff:\n--- want\n+++ got\n%s", diff)
	}
}
