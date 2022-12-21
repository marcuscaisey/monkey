package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/marcuscaisey/monkey/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %v. Welcome to the Monkey REPL!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
