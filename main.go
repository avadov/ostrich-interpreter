package main

import (
	"fmt"
	"os"
	"os/user"
	"ostrich-interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Ostrich Interpreter (alpha) [%s]\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}