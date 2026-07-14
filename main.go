package main

import (
	"fmt"
	"os"
)

type Entry struct {
	Service  string
	Username string
	Password string
}

func main() {
	TerminalArgs := os.Args

	switch TerminalArgs[1] {
	case "add":

	case "retrieve":

	case "list":

	}

	// fmt.Println()
	// result, err := SaveData(entries)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(result)
	// }

	fmt.Println(RetrieveData())
}
