package main

import (
	"fmt"
)

var Commands = map[string]struct{}{
	"add":       {},
	"get":       {},
	"list":      {},
	"delete":    {},
	"update":    {},
	"create-db": {}}

var MasterPassword = "IamOwner"

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

	_, CommandExists := Commands[args[1]]

	fmt.Println(CommandExists)

	if !CommandExists {
		fmt.Println("Command not found")
		return false
	}

	return true
}

func ArgsFunc(args []string, Mpassword string) {
	switch args[1] {
	case "add":
		HandleAdd(args, Mpassword)

	case "get":
		HandleGet(args, Mpassword)

	case "list":
		HandleList(args, Mpassword)

	case "delete":
		HandleDelete(args, Mpassword)

	case "update":
		HandleUpdate(args, Mpassword)

	case "create-db":
		CreateDb(Mpassword)
		fmt.Println("Database has been created.")

	default:
		return
	}
}

func ConfirmMaster() (string, bool) {
	var password string
	fmt.Println("Enter the master password")
	fmt.Scanln(&password)

	if password != MasterPassword {
		fmt.Println("You are not authorised")
		return "", false
	}
	return password, true
}
