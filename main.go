package main

import (
	"os"
)

func main() {
	TerminalArgs := os.Args
	CheckArgs(TerminalArgs)

	ArgsFunc(TerminalArgs)
}
