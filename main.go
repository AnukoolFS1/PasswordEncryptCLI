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
		fmt.Println("Retrieving data, Please Wait...")
		userdata := Retrieve(TerminalArgs[2])
		fmt.Println(userdata)

	case "list":
		fmt.Println("You want to list")
		entries := List()
		fmt.Println(entries)

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

func (entry Entry) String() string {

	return fmt.Sprintf("Username: %v\nService: %v\nPassword: %v\n", entry.Username, entry.Service, entry.Password)
}

func (entry DuplicateTypeEntry) String() string {

	return fmt.Sprintf(
		"\nUsername: %s | Service: %s | Password: %s",
		entry.Username,
		entry.Service,
		entry.Password)
}
