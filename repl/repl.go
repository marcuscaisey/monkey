// Package repl exports a Start function which starts a Read-Eval-Print-Loop (REPL) for the Monkey language.
package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/marcuscaisey/monkey/lexer"
	"github.com/marcuscaisey/monkey/token"
)

// Start starts the REPL, reading input from the given [io.Reader] and writing output to the given [io.Writer].
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, "> ")
		if !scanner.Scan() {
			return
		}
		line := scanner.Text()
		lexer := lexer.New(line)
		for tok := lexer.NextToken(); tok.Type != token.EOF; tok = lexer.NextToken() {
			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}
