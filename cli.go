package main

import (
	"fmt"
)

func CheckArgs(args []string) bool {
	if len(args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("add [service] [username] [password]")
		fmt.Println("update [username] [password]")
		fmt.Println("delete [username]")
		fmt.Println("get [username]")
		fmt.Println("list")
		return false
	}

	return true
}

func ArgsFunc(args []string) {
	switch args[1] {
	case "add":
		HandleAdd(args)

	case "get":
		HandleGet(args)

	case "list":
		HandleList(args)

	case "delete":
		HandleDelete(args)

	case "update":
		HandleUpdate(args)

	case "create-db":
		CreateDb()
		fmt.Println("Database has been created.")
		
	default:
		fmt.Println("Command not found")
	}
}
