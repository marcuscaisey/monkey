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
		if err := printAllTokens(line, out); err != nil {
			fmt.Fprintln(out, err)
			continue
		}
	}
}

func printAllTokens(src string, out io.Writer) error {
	lexer := lexer.New(src)
	for {
		nextToken, err := lexer.NextToken()
		if err != nil {
			return err
		}
		fmt.Fprintf(out, "%+v\n", nextToken)
		if nextToken.Type == token.EOF {
			break
		}
	}
	return nil
}
