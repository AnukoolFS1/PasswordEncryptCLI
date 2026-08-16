package main

import (
	"os"
)

func main() {
	TerminalArgs := os.Args
	if !CheckArgs(TerminalArgs) {
		return
	}

	Mpassword, ok := ConfirmMaster()
	if !ok {
		return
	}

	ArgsFunc(TerminalArgs, Mpassword)
}
