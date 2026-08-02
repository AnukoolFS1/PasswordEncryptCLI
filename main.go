package main

import (
	"os"
)

func main() {
	TerminalArgs := os.Args

	if !CheckArgs(TerminalArgs) || !ConfirmMaster() {
		return
	}

	ArgsFunc(TerminalArgs)
}
