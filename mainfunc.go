package main

import (
	"fmt"
)

func CheckArgs(args []string) {
	if len(args) < 2 {
		fmt.Println("Usage:")
		fmt.Println("add [service] [username] [password]")
		fmt.Println("get [username]")
		fmt.Println("list")
		return
	}
}

func ArgsFunc(args []string) {
	switch args[1] {
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

	case "create Database":
		CreateDb()

	default:
		fmt.Println("Command not found")
	}
}
