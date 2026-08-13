package main

// import (
// 	"fmt"
// )

// var MasterPassword = "IamOwner"

// func CheckArgs(args []string) bool {
// 	if len(args) < 2 {
// 		fmt.Println("Usage:")
// 		fmt.Println("add [service] [username] [password]")
// 		fmt.Println("update [username] [password]")
// 		fmt.Println("delete [username]")
// 		fmt.Println("get [username]")
// 		fmt.Println("list")
// 		return false
// 	}

// 	return true
// }

// func ArgsFunc(args []string) {
// 	switch args[1] {
// 	case "add":
// 		HandleAdd(args)

// 	case "get":
// 		HandleGet(args)

// 	case "list":
// 		HandleList(args)

// 	case "delete":
// 		HandleDelete(args)

// 	case "update":
// 		HandleUpdate(args)

// 	case "create-db":
// 		CreateDb()
// 		fmt.Println("Database has been created.")

// 	default:
// 		fmt.Println("Command not found")
// 	}
// }

// func ConfirmMaster() bool {
// 	var password string
// 	fmt.Println("Enter the master password")
// 	fmt.Scanln(&password)

// 	if password != MasterPassword {
// 		fmt.Println("You are not authorised")
// 		return false
// 	}
// 	return true
// }

