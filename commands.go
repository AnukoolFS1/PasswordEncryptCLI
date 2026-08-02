package main

import (
	"errors"
	"fmt"
)

func HandleAdd(args []string) {
	if len(args) < 5 {
		SuccessOrError("", errors.New("Please provide:\nadd [service] [username] [password]"))
		return
	}
	fmt.Println("Adding data, Please Wait...")
	result, err := Add(Entry{args[2], args[3], args[4]})
	SuccessOrError(result, err)
}

func HandleGet(args []string) {
	fmt.Println("Retrieving data, Please Wait...")
	userdata, err := Retrieve(args[2])
	SuccessOrError(userdata.String(), err)
}

func HandleUpdate(args []string) {
	if len(args) < 4 {
		fmt.Println("Please provide a username to delete entry.")
	} else {
		result, err := UpdatePassword(args[2], args[3])
		SuccessOrError(result, err)
	}
}

func HandleDelete(args []string) {
	if len(args) < 3 {
		fmt.Println("Please provide a username to delete entry.")
	} else {
		result, err := DeleteUser(args[2])
		SuccessOrError(result, err)
	}
}

func HandleList(args []string) {
	fmt.Println("Listing Data... Please Wait.")
	entries, err := List()
	if err != nil {
		fmt.Println(err)
	}
	for _, entry := range entries {
		fmt.Println(entry.SafeString())
	}
}

func SuccessOrError(success string, err error) {
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(success)
}
