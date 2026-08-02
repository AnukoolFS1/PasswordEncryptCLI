package main

import (
	"errors"
	"fmt"
	"os"
)

var TerminalArgs = os.Args

func HandleAdd() {
	if len(TerminalArgs) < 5 {
		SuccessOrError("", errors.New("Please provide:\nadd [service] [username] [password]"))
		return
	}
	fmt.Println("Adding data, Please Wait...")
	result, err := Add(Entry{TerminalArgs[2], TerminalArgs[3], TerminalArgs[4]})
	SuccessOrError(result, err)
}

func HandleGet() {
	fmt.Println("Retrieving data, Please Wait...")
	userdata, err := Retrieve(TerminalArgs[2])
	SuccessOrError(userdata.String(), err)
}

func HandleUpdate() {
	if len(TerminalArgs) == 3 || len(TerminalArgs) < 4 {
		fmt.Println("Please provide a username to delete entry.")
	} else {
		result, err := UpdatePassword(TerminalArgs[2], TerminalArgs[3])
		SuccessOrError(result, err)
	}
}

func HandleDelete() {
	if len(TerminalArgs) == 2 || len(TerminalArgs) < 3 {
		fmt.Println("Please provide a username to delete entry.")
	} else {
		result, err := DeleteUser(TerminalArgs[2])
		SuccessOrError(result, err)
	}
}

func HandleList() {
	fmt.Println("Listing Data... Please Wait.")
	entries, err := List()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(entries)
}

func SuccessOrError(success string, err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(success)
}
