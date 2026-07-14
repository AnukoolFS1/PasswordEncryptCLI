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
		if len(TerminalArgs) < 5 {
			fmt.Println("Please complete provide arguments: add [service] [username] [password]")
		}
		fmt.Println("Adding data, Please Wait...")
		add(Entry{TerminalArgs[2], TerminalArgs[3], TerminalArgs[4]})
		fmt.Println("Data has been added.")

	case "retrieve":
		fmt.Println("Retrieve data, Please Wait...")
		Retrieve()

	case "list":
		fmt.Println("You want to list")
		List()

	}

	// fmt.Println()
	// result, err := SaveData(entries)

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(result)
	// }

	// fmt.Println(RetrieveData())
}
