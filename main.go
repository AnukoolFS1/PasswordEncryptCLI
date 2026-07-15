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
	var result string
	var err error

	switch TerminalArgs[1] {
	case "add":
		if len(TerminalArgs) < 5 {
			fmt.Println("Please complete provide arguments: add [service] [username] [password]")
		}
		fmt.Println("Adding data, Please Wait...")
		result, err = add(Entry{TerminalArgs[2], TerminalArgs[3], TerminalArgs[4]})

	case "retrieve":
		fmt.Println("Retrieve data, Please Wait...")
		Retrieve(TerminalArgs[2])

	case "list":
		fmt.Println("You want to list")
		List()

	default:
		fmt.Println("Command not found")

	}

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	// result, err := SaveData(entries)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(result)
	// }

	// fmt.Println(RetrieveData())
}
