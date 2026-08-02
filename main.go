package main

import (
	"fmt"
)

func main() {


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
		HandleDelete()

	case "update":
		HandleUpdate()

	default:
		fmt.Println("Command not found")
	}
}
