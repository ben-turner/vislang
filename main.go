package main

import (
	"bufio"
	"fmt"
	"github.com/ben-turner/vislang/interpreter"
	"github.com/ben-turner/vislang/tokenizer"
	"os"
)

func run() int {
	r := bufio.NewReader(os.Stdin)
	tokens, line, err := tokenizer.Parse(r)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error on line %d: %v", line+1, err)
		return 1
	}
	if err := interpreter.Run(tokens); err != nil {
		fmt.Fprintf(os.Stderr, "Runtime error: %v", err)
	}

	return 0
}

func main() {
	os.Exit(run())
}
