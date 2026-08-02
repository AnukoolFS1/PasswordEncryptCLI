package main

import (
	"os"
)

func main() {
	TerminalArgs := os.Args

	if !CheckArgs(TerminalArgs) {
		return
	}

	ArgsFunc(TerminalArgs)
}
