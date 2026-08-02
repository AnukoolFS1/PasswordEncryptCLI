package main

import (
	"fmt"
)

func main() {
	var result string
	var err error

	if len(TerminalArgs) < 2 {
		fmt.Println("Usage:")
		fmt.Println("add [service] [username] [password]")
		fmt.Println("get [username]")
		fmt.Println("list")
		return
	}

	switch TerminalArgs[1] {
	case "add":
		HandleAdd()

	case "get":
		HandleGet()

	case "list":
		HandleList()

	case "delete":
		if len(TerminalArgs) == 2 || len(TerminalArgs) < 3 {
			fmt.Println("Please provide a username to delete entry.")
		} else {
			result = DeleteUser(TerminalArgs[2])
		}

	case "update":
		if len(TerminalArgs) == 3 || len(TerminalArgs) < 4 {
			fmt.Println("Please provide a username to delete entry.")
		} else {
			result = UpdatePassword(TerminalArgs[2], TerminalArgs[3])
		}

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
